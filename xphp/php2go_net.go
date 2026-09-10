package xphp

import (
	"encoding/binary"
	"net"
	"os"
	"strings"
)

// Gethostname 获取本地主机名（对应 PHP gethostname()）。
func Gethostname() (string, error) {
	return os.Hostname()
}

// Gethostbyname 获取指定 Internet 主机名对应的 IPv4 地址（对应 PHP gethostbyname()）。
func Gethostbyname(hostname string) (string, error) {
	ips, err := net.LookupIP(hostname)
	if ips != nil {
		for _, v := range ips {
			if v.To4() != nil {
				return v.String(), nil
			}
		}
		return "", nil
	}
	return "", err
}

// Gethostbynamel 获取指定 Internet 主机名对应的全部 IPv4 地址列表（对应 PHP gethostbynamel()）。
func Gethostbynamel(hostname string) ([]string, error) {
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

// Gethostbyaddr 获取指定 IP 地址对应的 Internet 主机名（对应 PHP gethostbyaddr()）。
func Gethostbyaddr(ipAddress string) (string, error) {
	names, err := net.LookupAddr(ipAddress)
	if names != nil {
		return strings.TrimRight(names[0], "."), nil
	}
	return "", err
}

// IP2long 将 IPv4 地址字符串转换为 32 位无符号整数（对应 PHP ip2long()）。
// 仅支持 IPv4。
func IP2long(ipAddress string) uint32 {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ip.To4())
}

// Long2ip 将 32 位无符号整数转换为 IPv4 地址字符串（对应 PHP long2ip()）。
// 仅支持 IPv4。
func Long2ip(properAddress uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, properAddress)
	ip := net.IP(ipByte)
	return ip.String()
}
