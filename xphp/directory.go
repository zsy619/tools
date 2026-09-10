package xphp

import (
	"os"
)

// Chdir 将当前工作目录切换到 dir；底层包装 os.Chdir。
func Chdir(dir string) error {
	return os.Chdir(dir)
}

// Scandir 列出 dir 下的所有文件和目录名（不递归）。
// 读取或打开失败时返回错误；成功时返回的切片不保证顺序。
func Scandir(dir string) ([]string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}

	names, err := f.Readdirnames(-1)
	f.Close()

	return names, err
}
