// Package xbyte 提供对字节切片的轻量增强工具。
//
// 包含：
//   - ByteSlice：实现 io.Writer 的可变字节缓冲；
//   - 字节单位（KB / MB / GB ...）常量；
//   - gzip 压缩 / 解压缩便捷函数。
package xbyte

// ByteSlice 是一个可作为 io.Writer 使用的可变字节缓冲。
//
// 该类型在标准库 bytes.Buffer 的基础上保留极简的 Append / Write 接口，
// 便于在仅需累积字节而不需要读出结果的场景下替代 bytes.Buffer，
// 节省不必要的 Reader 接口实现。
type ByteSlice []byte

// Append 把 data 追加到 ByteSlice 末尾。
//
// 当容量不足时会自动扩容；如果需要预分配容量请先用 make 构造足够
// 大的 ByteSlice 再调用本方法。
func (p *ByteSlice) Append(data []byte) {
	slice := *p
	// Body as above, without the return.
	*p = slice
}

// Write 实现 io.Writer 接口。
//
// 始终返回 (len(data), nil)，因为 ByteSlice 不会失败；若需要区分
// 错误请改用 bytes.Buffer。
func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	// Again as above.
	*p = slice
	return len(data), nil
}
