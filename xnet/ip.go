// Package xnet 提供 IP / 地理位置相关的网络工具。
//
// 包含：
//   - 内网 IP 判定（HasLocalIP、HasLocalIPAddr 等）；
//   - HTTP 客户端 IP 解析（ClientIP、ClientPublicIP、RemoteIP），
//     支持 X-Forwarded-For / X-Real-IP 反向代理头；
//   - IPv4 与 uint 的双向转换（IPString2Long、Long2IPString、IP2Long、
//     Long2IP）。
package xnet

import (
	"errors"
	"math"
	"net"
	"net/http"
	"strings"
)

// HasLocalIPddr 是 HasLocalIPAddr 的旧拼写兼容版本。
//
// 由于原名 HasLocalIPddr 拼写错误，仅为兼容历史调用保留，**已弃用**，
// 新代码请直接使用 HasLocalIPAddr。
//
// Deprecated: 此为一个错误名称错误拼写的函数，计划在将来移除，请使用 HasLocalIPAddr 函数
func HasLocalIPddr(ip string) bool {
	return HasLocalIPAddr(ip)
}

// HasLocalIPAddr 判断字符串形式的 IP 是否为内网 / 回环地址。
//
// 不可解析的 ip 返回 false（与 HasLocalIP 行为一致）。如需支持
// IPv6，请直接使用 net.IP.IsPrivate / IsLoopback 等。
func HasLocalIPAddr(ip string) bool {
	return HasLocalIP(net.ParseIP(ip))
}

// HasLocalIP 判断 net.IP 是否为内网 / 回环地址。
//
// 覆盖段：
//   - IPv4 loopback (127.0.0.0/8)；
//   - 10.0.0.0/8、172.16.0.0/12、192.168.0.0/16（RFC1918 私网）；
//   - 169.254.0.0/16（链路本地）。
//
// 实现采用直接字节比较，效率高于 net.IPNet.Contains，详见：
// https://github.com/thinkeridea/go-extend/issues/2
//
// 只支持 IPv4，IPv6 直接返回 false（除 IsLoopback 命中外）。
func HasLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}

// ClientIP 在反向代理场景下尽力获取真实客户端 IP。
//
// 优先级：X-Forwarded-For（取首项）→ X-Real-IP → RemoteAddr。
//
// 安全注意：以上头部由客户端 / 上游代理控制，直接信任存在伪造风险；
// 建议仅在内网或已校验代理链的环境使用，公开接入请改用受信任的
// ProxyHeaders 中间件或调用 ClientPublicIP 排除内网伪造。
func ClientIP(r *http.Request) string {
	ip := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

// ClientPublicIP 尽力获取客户端的公网 IP（排除内网 / 回环）。
//
// 与 ClientIP 的区别：会跳过任何内网 / 回环 IP；X-Forwarded-For 列表
// 按从左到右遍历，第一个非内网项即返回；最后才回退到 RemoteAddr。
//
// 全部为内网或解析失败时返回空串。
func ClientPublicIP(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		if ip = strings.TrimSpace(ip); ip != "" && !HasLocalIPAddr(ip) {
			return ip
		}
	}

	if ip = strings.TrimSpace(r.Header.Get("X-Real-Ip")); ip != "" && !HasLocalIPAddr(ip) {
		return ip
	}

	if ip = RemoteIP(r); !HasLocalIPAddr(ip) {
		return ip
	}

	return ""
}

// RemoteIP 快速从 r.RemoteAddr 解析出 IP 部分（去掉端口）。
//
// 当 RemoteAddr 不是合法的 host:port 形式时返回空串。
func RemoteIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

// IPString2Long 把点分十进制 IPv4 字符串转换为 uint。
//
// 仅支持 IPv4；传入 IPv6 / 无法解析字符串时返回 (0, error)。
// 字节序为网络字节序（大端），等价于 inet_aton。
func IPString2Long(ip string) (uint, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, errors.New("invalid ipv4 format")
	}

	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}

// Long2IPString 把 uint 转回点分十进制 IPv4 字符串，是 IPString2Long 的逆操作。
//
// 当 i > math.MaxUint32 时返回错误，避免 64 位平台上的静默截断。
func Long2IPString(i uint) (string, error) {
	if i > math.MaxUint32 {
		return "", errors.New("beyond the scope of ipv4")
	}

	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)

	return ip.String(), nil
}

// IP2Long 把 net.IP 转成网络字节序的 uint，是 IPString2Long 在
// net.IP 输入下的版本。
//
// 仅接受 4 字节 IPv4；IPv6 或 v4-in-v6 之外的 IP 返回错误。
func IP2Long(ip net.IP) (uint, error) {
	b := ip.To4()
	if b == nil {
		return 0, errors.New("invalid ipv4 format")
	}

	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}

// Long2IP 把 uint 转回 net.IP，是 IP2Long 的逆操作。
//
// 当 i > math.MaxUint32 时返回错误，避免 64 位平台上的静默截断。
// 与 Long2IPString 的区别：返回 net.IP 而非字符串，适合继续参与
// net.IP 相关运算。
func Long2IP(i uint) (net.IP, error) {
	if i > math.MaxUint32 {
		return nil, errors.New("beyond the scope of ipv4")
	}

	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)

	return ip, nil
}
