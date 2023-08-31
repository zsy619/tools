package xnet

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/axgle/mahonia"
	"haedu.gov.cn/tools/xjson"
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
	if ip == "::1" || ip == "127.0.0.1" {
		info.Err = "内网IP"
		return info
	}
	url := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		fmt.Println(string(body))
		if err != nil {
			info.Err = err.Error()
			return info
		}
		xjson.Unmarshal(body, &info)
	}
	return info
}

// GetCityByIp 获取IP城市信息
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "::1" || ip == "127.0.0.1" {
		return "内网IP"
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
