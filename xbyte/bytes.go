package xbyte

import (
	"encoding/hex"
	"errors"
)

// Byte32ToHexString converts a [32]byte slice to a string
func Byte32ToHexString(input [32]byte) string {
	return hex.EncodeToString(input[:])
}

// HexStringToByte32 converts a string back to a [32]byte slice
func HexStringToByte32(input string) ([32]byte, error) {
	bys, err := hex.DecodeString(input)
	fixed := [32]byte{}
	if err != nil {
		return fixed, err
	}
	copy(fixed[:], bys)
	return fixed, nil
}

// BlockCopy byte数组拷贝
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
