package xbyte

import (
	"bytes"
	"compress/gzip"
	"io"
)

// DecodeGzipBytes 把 gzip 压缩的字节流解压缩。
//
// 当 meta 不是合法的 gzip 数据时返回 (nil, error)；io.ReadAll 阶段
// 的错误也会原样返回。
//
// 注意：本函数忽略了 gzip.NewReader 的错误，调用方应检查返回的
// error 字段来判断输入是否合法。
func DecodeGzipBytes(meta []byte) ([]byte, error) {
	b := bytes.Buffer{}
	b.Write(meta)
	r, _ := gzip.NewReader(&b)
	defer r.Close()
	datas, readErr := io.ReadAll(r)

	if readErr != nil {
		return nil, readErr
	}

	return datas, nil
}

// EncodeGzipBytes 用 gzip 默认级别压缩字节流并返回结果。
//
// 始终返回新分配的字节切片，调用方持有所有权；写入 / Flush / Close
// 错误被显式忽略，因为对内存 buffer 来说几乎不可能失败。
func EncodeGzipBytes(meta []byte) []byte {
	b := bytes.Buffer{}
	w := gzip.NewWriter(&b)
	defer w.Close()

	_, _ = w.Write(meta)
	_ = w.Flush()

	return b.Bytes()
}
