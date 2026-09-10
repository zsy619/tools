package xcrypto

import (
	"bufio"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// Base64StdEncode 使用标准 Base64 编码对字符串进行编码。
//
// 参数：
//   - s: 待编码的原始字符串。
//
// 返回值：Base64 编码结果字符串。
func Base64StdEncode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64StdDecode 对标准 Base64 编码的字符串进行解码；解码失败时返回空字符串。
//
// 参数：
//   - s: 已编码的 Base64 字符串。
//
// 返回值：解码后的字符串；解码失败时返回 ""。
func Base64StdDecode(s string) string {
	b, _ := base64.StdEncoding.DecodeString(s)
	return string(b)
}

// Md5String 计算字符串的 MD5 值（十六进制字符串形式）。
//
// 参数：
//   - s: 待计算哈希的字符串。
//
// 返回值：MD5 哈希值的十六进制字符串（32 个字符）。
func Md5String(s string) string {
	h := md5.New()
	_, _ = h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5File 计算文件的 MD5 值（按 64KB 分块读取以避免大文件占用过多内存）。
//
// 参数：
//   - filename: 文件路径。
//
// 返回值：
//   - string: 文件的 MD5 哈希值的十六进制字符串。
//   - error: 文件不存在、不可读或读取过程中发生错误时返回错误；路径为目录时返回 ("", nil)。
func Md5File(filename string) (string, error) {
	if fileInfo, err := os.Stat(filename); err != nil {
		return "", err
	} else if fileInfo.IsDir() {
		return "", nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()

	chunkSize := 65536
	for buf, reader := make([]byte, chunkSize), bufio.NewReader(file); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		hash.Write(buf[:n])
	}

	checksum := fmt.Sprintf("%x", hash.Sum(nil))
	return checksum, nil
}

// HmacMd5 使用 MD5 算法计算给定数据与密钥的 HMAC 值（十六进制字符串）。
//
// 参数：
//   - data: 待签名的数据。
//   - key: 密钥。
//
// 返回值：HMAC-MD5 值的十六进制字符串。
func HmacMd5(data, key string) string {
	h := hmac.New(md5.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// HmacSha1 使用 SHA1 算法计算给定数据与密钥的 HMAC 值（十六进制字符串）。
//
// 参数：
//   - data: 待签名的数据。
//   - key: 密钥。
//
// 返回值：HMAC-SHA1 值的十六进制字符串。
func HmacSha1(data, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// HmacSha256 使用 SHA256 算法计算给定数据与密钥的 HMAC 值（十六进制字符串）。
//
// 参数：
//   - data: 待签名的数据。
//   - key: 密钥。
//
// 返回值：HMAC-SHA256 值的十六进制字符串。
func HmacSha256(data, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// HmacSha512 使用 SHA512 算法计算给定数据与密钥的 HMAC 值（十六进制字符串）。
//
// 参数：
//   - data: 待签名的数据。
//   - key: 密钥。
//
// 返回值：HMAC-SHA512 值的十六进制字符串。
func HmacSha512(data, key string) string {
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// Sha1 计算字符串的 SHA1 哈希值（十六进制字符串形式）。
//
// 参数：
//   - data: 待计算哈希的字符串。
//
// 返回值：SHA1 哈希值的十六进制字符串（40 个字符）。
func Sha1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

// Sha256 计算字符串的 SHA256 哈希值（十六进制字符串形式）。
//
// 参数：
//   - data: 待计算哈希的字符串。
//
// 返回值：SHA256 哈希值的十六进制字符串（64 个字符）。
func Sha256(data string) string {
	sha256 := sha256.New()
	sha256.Write([]byte(data))
	return hex.EncodeToString(sha256.Sum([]byte("")))
}

// Sha512 计算字符串的 SHA512 哈希值（十六进制字符串形式）。
//
// 参数：
//   - data: 待计算哈希的字符串。
//
// 返回值：SHA512 哈希值的十六进制字符串（128 个字符）。
func Sha512(data string) string {
	sha512 := sha512.New()
	sha512.Write([]byte(data))
	return hex.EncodeToString(sha512.Sum([]byte("")))
}
