package xphp

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"hash/crc32"
	"io/ioutil"
	"os"
)

func Md5(s string) (string, error) {
	h := md5.New()
	if _, err := h.Write([]byte(s)); err != nil {
		return "", err
	}
	result := h.Sum(nil)
	return hex.EncodeToString(result), nil
}

func Md5_file(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	contentbyte, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	h := md5.New()
	h.Write(contentbyte)
	return hex.EncodeToString(h.Sum(nil)), nil

}

func Base64_decode(str string) (string, error) {
	bt, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(bt), nil
}

func Base64_encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Md5File md5_file()
func Md5File(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Sha1 sha1()
func Sha1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha1File sha1_file()
func Sha1File(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := sha1.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Crc32 crc32()
func Crc32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}
