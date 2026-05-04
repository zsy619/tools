package xphp

import (
	"os"
	"path/filepath"
	"syscall"
)

// Dirname returns a parent directory's path
func Dirname(path string) string {
	return filepath.Dir(path)
}

// DirnameWithLevels returns a parent directory's path which is
// levels up from the current directory.
func DirnameWithLevels(path string, levels int) string {
	for i := 0; i < levels; i++ {
		path = Dirname(path)
	}
	return path
}

// Rmdir removes empty directory
func Rmdir(dirname string) error {
	return syscall.Rmdir(dirname)
}

// Symlink creates a symbolic link
func Symlink(target, link string) error {
	return os.Symlink(target, link)
}

// Link create a hard link
func Link(target, link string) error {
	return os.Link(target, link)
}

// IsLink tells whether the filename is a symbolic link
func IsLink(filename string) bool {
	fi, err := os.Lstat(filename)
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeSymlink == os.ModeSymlink
}
