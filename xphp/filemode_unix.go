package xphp

// IsExecutableSyscall 判断 filename 是否可执行。
// 当前始终返回 false，仅作为占位实现以保持跨平台编译。
func IsExecutableSyscall(filename string) bool {
	// return syscall.Access(filename, 0x1) == nil
	return false
}

// IsReadableSyscall 判断 filename 是否存在且可读。
// 当前始终返回 false，仅作为占位实现以保持跨平台编译。
func IsReadableSyscall(filename string) bool {
	// return syscall.Access(filename, 0x4) == nil
	return false
}

// IsWritableSyscall 判断 filename 是否可写。
// 当前始终返回 false，仅作为占位实现以保持跨平台编译。
func IsWritableSyscall(filename string) bool {
	// return syscall.Access(filename, 0x2) == nil
	return false
}
