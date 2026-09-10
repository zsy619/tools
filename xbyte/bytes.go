package xbyte

import (
	"encoding/hex"
	"errors"
)

// Byte32ToHexString 把 [32]byte 转成 64 个字符的十六进制字符串。
//
// 常用于以太坊 / 区块链场景下的 32 字节哈希地址、签名 R/S 拼接
// 等。返回字符串始终为 64 个字符（小写 a-f）。
func Byte32ToHexString(input [32]byte) string {
	return hex.EncodeToString(input[:])
}

// HexStringToByte32 把 64 个字符的十六进制字符串解码为 [32]byte。
//
// 解码失败时返回 ([32]byte{}, err)。当 input 长度超过 64 字符时，
// 仅复制前 32 字节；如果长度不足，剩余字节为零值。
func HexStringToByte32(input string) ([32]byte, error) {
	bys, err := hex.DecodeString(input)
	fixed := [32]byte{}
	if err != nil {
		return fixed, err
	}
	copy(fixed[:], bys)
	return fixed, nil
}

// BlockCopy 在 src 与 dst 之间复制 count 个字节，等价于带边界检查的 copy。
//
// 边界条件：
//   - src 必须满足 srcOffset + count <= len(src)；
//   - dst 必须满足 dstOffset + count <= len(dst)。
//
// 任一边界越界即返回 (false, error)，不修改 dst。
func BlockCopy(src []byte, srcOffset int, dst []byte, dstOffset, count int) (bool, error) {
	srcLen := len(src)
	if srcOffset > srcLen || count > srcLen || srcOffset+count > srcLen {
		return false, errors.New("源缓冲区 索引超出范围")
	}
	dstLen := len(dst)
	if dstOffset > dstLen || count > dstLen || dstOffset+count > dstLen {
		return false, errors.New("目标缓冲区 索引超出范围")
	}
	index := 0
	for i := srcOffset; i < srcOffset+count; i++ {
		dst[dstOffset+index] = src[srcOffset+index]
		index++
	}
	return true, nil
}
