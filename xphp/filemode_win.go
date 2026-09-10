package xphp

import (
	"os"
	"path/filepath"
	"strings"
)

// IsExecutable 判断 filename 是否为 Windows 下的可执行文件。
// 先检查文件是否可读，再判断后缀（不区分大小写）是否为 .EXE。
// 文件不存在或 stat 出错时返回 false。
func IsExecutable(filename string) bool {
	if !IsReadable(filename) {
		return false
	}

	return strings.ToUpper(filepath.Ext(filename)) == ".EXE"
}

// IsWritable 判断 filename 是否可写。
// 当前仅当文件的 Mode 严格等于 0666 时返回 true。
// stat 出错时返回 false。
func IsWritable(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		return false
	}

	return fi.Mode() == 0666
}
