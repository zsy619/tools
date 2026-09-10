package xbyte

import (
	"fmt"
)

// ByteSize 表示以字节为单位的可读容量，支持 KB 到 YB 的自动格式化。
type ByteSize float64

const (
	_           = iota // 跳过 1 << 0
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// String 把字节数格式化为最合适的单位（保留两位小数）。
//
// 自适应选择 KB / MB / GB / TB / PB / EB / ZB / YB；小于 1 KB 时按
// "B" 输出。实现 fmt.Stringer 接口，可直接用于 fmt.Print 等场景。
func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}
