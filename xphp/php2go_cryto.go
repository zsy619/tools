package xphp

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"hash/crc32"
	"io"
	"os"
)

// Md5 计算字符串 s 的 MD5 摘要并以十六进制字符串返回。
func Md5(s string) (string, error) {
	h := md5.New()
	if _, err := h.Write([]byte(s)); err != nil {
		return "", err
	}
	result := h.Sum(nil)
	return hex.EncodeToString(result), nil
}

// Md5_file 计算 filepath 指向文件的 MD5 摘要并以十六进制字符串返回。
// 文件不存在或读取失败时返回错误。
func Md5_file(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	contentbyte, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	h := md5.New()
	h.Write(contentbyte)
	return hex.EncodeToString(h.Sum(nil)), nil
}

// Base64_decode 使用标准 Base64 解码 str；输入非法时返回错误。
func Base64_decode(str string) (string, error) {
	bt, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(bt), nil
}

// Base64_encode 使用标准 Base64 编码 str。
func Base64_encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Md5File 实现 PHP md5_file()：读取 path 文件并返回其 MD5 摘要的十六进制字符串。
func Md5File(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Sha1 实现 PHP sha1()：计算 str 的 SHA-1 摘要并以十六进制字符串返回。
func Sha1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha1File 实现 PHP sha1_file()：计算 path 文件的 SHA-1 摘要的十六进制字符串。
func Sha1File(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := sha1.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Crc32 实现 PHP crc32()：使用 IEEE 多项式计算 str 的 CRC32 校验值。
func Crc32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}
