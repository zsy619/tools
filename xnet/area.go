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

// GetAreaByIp 获取IP地区信息
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

// GetCityByIp 获取IP城市信息
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

// src 字符串
// srcCode 字符串当前编码
// tagCode 要转换的编码
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

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
