package xphp

import "syscall"

// IsExecutableSyscall tells whether the filename is executable
func IsExecutableSyscall(filename string) bool {
	return syscall.Access(filename, 0x1) == nil
}

// IsReadableSyscall tells whether a file exists and is readable
func IsReadableSyscall(filename string) bool {
	return syscall.Access(filename, 0x4) == nil
}

// IsWritableSyscall tells whether the filename is writable
func IsWritableSyscall(filename string) bool {
	return syscall.Access(filename, 0x2) == nil
}
