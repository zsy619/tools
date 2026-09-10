package xphp

import (
	"encoding/binary"
	"net"
	"os"
	"strings"
)

// GetHostByAddr 通过 IP 地址反查主机名。
// 返回第一个 PTR 记录（去掉尾部点号）；解析失败时返回错误。
func GetHostByAddr(ipAddress string) (string, error) {
	names, err := net.LookupAddr(ipAddress)
	if len(names) > 0 {
		return strings.TrimRight(names[0], "."), nil
	}
	return "", err
}

// GetHostByName 查询 hostname 对应的第一个 IPv4 地址。
// hostname 不可解析或没有 IPv4 记录时返回空串与错误。
func GetHostByName(hostname string) (string, error) {
	ips, err := net.LookupIP(hostname)
	if len(ips) != 0 {
		for _, v := range ips {
			if v.To4() != nil {
				return v.String(), nil
			}
		}
	}
	return "", err
}

// GetHostByNamel 返回 hostname 对应的全部 IPv4 地址列表。
// 查询失败或无 IPv4 记录时返回 nil 与错误。
func GetHostByNamel(hostname string) ([]string, error) {
	ips, err := net.LookupIP(hostname)
	if ips != nil {
		var ipstrs []string
		for _, v := range ips {
			if v.To4() != nil {
				ipstrs = append(ipstrs, v.String())
			}
		}
		return ipstrs, nil
	}
	return nil, err
}

// GetHostName 返回本地主机名（包装 os.Hostname）。
func GetHostName() (string, error) {
	return os.Hostname()
}

// IP2Long 将点分十进制 IPv4 字符串转换为 uint32（大端）。
// 无法解析或非 IPv4 时返回 0。
func IP2Long(ipAddress string) uint32 {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return 0
	}
	ipByte := ip.To4()
	if ipByte == nil {
		return 0
	}

	return binary.BigEndian.Uint32(ipByte)
}

// Long2IP 将 uint32 数值转换回点分十进制 IPv4 字符串（大端）。
func Long2IP(properAddress uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, properAddress)
	ip := net.IP(ipByte)
	return ip.String()
}
