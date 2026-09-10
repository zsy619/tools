package xphp

import (
	"os"
	"path/filepath"
	"syscall"
)

// Dirname 返回 path 的父目录，等价于 filepath.Dir。
func Dirname(path string) string {
	return filepath.Dir(path)
}

// DirnameWithLevels 沿父目录方向回溯 levels 次，返回最终的目录路径。
// levels <= 0 时返回 path 本身的父目录。
func DirnameWithLevels(path string, levels int) string {
	for i := 0; i < levels; i++ {
		path = Dirname(path)
	}
	return path
}

// Rmdir 删除 dirname 指定的空目录；底层包装 syscall.Rmdir。
// 目录非空或不存在时返回错误。
func Rmdir(dirname string) error {
	return syscall.Rmdir(dirname)
}

// Symlink 创建一个指向 target 的符号链接 link，等价于 os.Symlink。
func Symlink(target, link string) error {
	return os.Symlink(target, link)
}

// Link 创建一个从 link 指向 target 的硬链接，等价于 os.Link。
func Link(target, link string) error {
	return os.Link(target, link)
}

// IsLink 判断 filename 是否为符号链接。
// 使用 Lstat 而非 Stat，避免跟随符号链接；stat 出错时返回 false。
func IsLink(filename string) bool {
	fi, err := os.Lstat(filename)
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeSymlink == os.ModeSymlink
}
