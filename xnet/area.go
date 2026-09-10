// Package xnet 提供 IP / 地理位置相关的网络工具。
//
// 包含：
//   - 通过公网接口查询 IP 归属地（GetAreaByIp、GetCityByIp）；
//   - 编码转换辅助（ConvertToString）；
//   - 内网 IP 段判断（isPrivateSubnet）。
package xnet

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/axgle/mahonia"
	"golang.org/x/net/html/charset"

	"github.com/zsy619/tools/xjson"
)

// AreaInfo 描述一个 IP 的归属地信息，对应太平洋网络 IP 库返回字段。
//
// 字段含义：
//   - IP：查询的原始 IP；
//   - Pro / ProCode：省份与省份编码；
//   - City / CityCode：城市与城市编码；
//   - Region / RegionCode：区/县与区/县编码；
//   - Addr：拼接后的详细地址；
//   - RegionNames：省市县三级组合名；
//   - Err：当 IP 为内网、解析失败或接口不可用时填写的错误描述。
type AreaInfo struct {
	IP          string `json:"ip"`          // 本机IP
	Pro         string `json:"pro"`         // 省份
	ProCode     string `json:"proCode"`     // 省份编码
	City        string `json:"city"`        // 城市
	CityCode    string `json:"cityCode"`    // 城市编码
	Region      string `json:"region"`      // 地区
	RegionCode  string `json:"regionCode"`  // 地区编码
	Addr        string `json:"addr"`        // 详细地址
	RegionNames string `json:"regionNames"` // 地区名称
	Err         string `json:"err"`         // 错误信息
}

// GetAreaByIp 通过太平洋网络 IP 库查询公网 IP 的归属地信息。
//
// 行为：
//   - ip 为空字符串或 IPv6 回环 ::1：返回 Err="内网"；
//   - ip 属于 RFC1918 内网段：返回 Err="内网"；
//   - HTTP 调用失败：在 Err 中填入错误描述；
//   - 接口返回的 GBK 数据会自动用 charset.NewReader 转为 UTF-8。
//
// 注意：太平洋网络为外部免费接口，可能随时失效或限流；生产环境
// 建议替换为更稳定的服务或本地 IP 库。
func GetAreaByIp(ip string) AreaInfo {
	info := AreaInfo{}
	if ip == "" {
		return info
	}
	if ip == "::1" {
		info.Err = "内网"
		return info
	}
	netIp := net.ParseIP(ip)
	if isPrivateSubnet(netIp) {
		info.Err = "内网"
		return info
	}
	url := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		info.Err = err.Error()
		return info
	}
	defer resp.Body.Close()
	// 创建一个带自动转换编码的Reader
	utf8Reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		fmt.Println("Error creating charset reader:", err)
		return info
	}

	body, err := io.ReadAll(utf8Reader)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return info
	}
	xjson.Unmarshal(body, &info)

	return info
}

// GetCityByIp 仅返回 ip 所在的城市名；非内网时调用太平洋网络 IP 库，
// 失败或解析错误统一返回空串。
//
// 与 GetAreaByIp 的区别：
//   - 仅返回 city 字段；
//   - 不向 AreaInfo 写回详细错误信息，错误一律归一为空串；
//   - 不打印中间错误到 stdout，便于在静默环境使用。
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "::1" {
		return "内网"
	}
	netIp := net.ParseIP(ip)
	if isPrivateSubnet(netIp) {
		return "内网"
	}
	url := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := io.ReadAll(response.Body)
		bodystr := string(body)
		tmp := ConvertToString(bodystr, "gbk", "utf-8")
		p := make(map[string]interface{}, 0)
		if err := json.Unmarshal([]byte(tmp), &p); err == nil {
			return p["city"].(string)
		}
	}
	return ""
}

// ConvertToString 把 src 从 srcCode 编码转为 tagCode 编码的字符串。
//
// 参数：
//   - src：源字符串；
//   - srcCode：源字符串的字符编码，如 "gbk"、"big5"；
//   - tagCode：目标字符编码，通常为 "utf-8"。
//
// 实现基于 axgle/mahonia 的码表查找，零依赖 cgo；如果编码名无效
// 会原样返回 src 或部分转换结果。
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// isPrivateSubnet 判断 ip 是否属于 IPv4 内网 / 回环段。
//
// 覆盖：127.0.0.0/8、10.0.0.0/8、172.16.0.0/12、192.168.0.0/16。
// IPv6 地址直接返回 false（仅 ::1 在调用方单独判断）。
//
// 注意：每次调用都会重新构造 []*net.IPNet。如需高频调用，建议把
// CIDR 列表提升为包级变量以减少重复解析。
func isPrivateSubnet(ip net.IP) bool {
	if ip.To4() == nil {
		return false
	}

	privateIPBlocks := []*net.IPNet{}
	for _, cidr := range []string{
		"127.0.0.0/8",    // IPv4 loopback
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
	} {
		if _, block, err := net.ParseCIDR(cidr); err == nil {
			privateIPBlocks = append(privateIPBlocks, block)
		}
	}

	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}
