package xphp

// IsExecutableSyscall tells whether the filename is executable
func IsExecutableSyscall(filename string) bool {
	// return syscall.Access(filename, 0x1) == nil
	return false
}

// IsReadableSyscall tells whether a file exists and is readable
func IsReadableSyscall(filename string) bool {
	// return syscall.Access(filename, 0x4) == nil
	return false
}

// IsWritableSyscall tells whether the filename is writable
func IsWritableSyscall(filename string) bool {
	// return syscall.Access(filename, 0x2) == nil
	return false
}
