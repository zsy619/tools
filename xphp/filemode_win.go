package xphp

import (
	"os"
	"path/filepath"
	"strings"
)

// IsExecutable tells whether the filename is executable
func IsExecutable(filename string) bool {
	if !IsReadable(filename) {
		return false
	}

	return strings.ToUpper(filepath.Ext(filename)) == ".EXE"
}

// IsWritable tells whether the filename is writable
func IsWritable(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		return false
	}

	return fi.Mode() == 0666
}
