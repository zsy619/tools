package xcrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

type ChsiAes struct {
	key []byte
}

func NewChsiAes(key []byte) (*ChsiAes, error) {
	if key == nil {
		return nil, errors.New("key错误")
	}
	return &ChsiAes{
		key: key,
	}, nil
}

//加密数据
func (a *ChsiAes) Encrypt(src string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(string(a.key))
	if err != nil {
		fmt.Println("key error1", err)
		return "", errors.New("key error")
	}
	// fmt.Println("Encrypt：", string(a.key), string(key))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
		return "", errors.New("key empty")
	}
	if src == "" {
		fmt.Println("plain content empty")
		return "", errors.New("plain content empty")
	}
	content := a.PKCS5Padding([]byte(src), aesBlockEncrypter.BlockSize())
	encrypted := make([]byte, len(content))
	ivKey := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	aesEncrypter := cipher.NewCBCEncrypter(aesBlockEncrypter, ivKey)
	// aesEncrypter := cipher.NewCBCEncrypter(aesBlockEncrypter, a.key[:aesBlockEncrypter.BlockSize()])
	aesEncrypter.CryptBlocks(encrypted, content)

	return base64.StdEncoding.EncodeToString(encrypted), err
}

//解密数据
func (a *ChsiAes) Decrypt(crypt string) (data string, err error) {
	dest, err := base64.StdEncoding.DecodeString(crypt)
	if err != nil {
		return "", err
	}
	key, err := base64.StdEncoding.DecodeString(string(a.key))
	if err != nil {
		fmt.Println("key error1", err)
		return "", errors.New("key error")
	}
	// fmt.Println("Decrypt：", string(a.key), string(key))
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
		return "", errors.New("key empty")
	}
	if len(crypt) == 0 {
		fmt.Println("plain content empty")
		return "", errors.New("plain content empty")
	}
	ivKey := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ecb := cipher.NewCBCDecrypter(block, ivKey)
	// ecb := cipher.NewCBCDecrypter(block, a.key[:block.BlockSize()])
	decrypted := make([]byte, len(dest))
	ecb.CryptBlocks(decrypted, dest)
	rt := a.PKCS5Trimming(decrypted)
	return string(rt), nil
}

//PKCS5UnPadding
func (a *ChsiAes) PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

/**
PKCS5包装
*/
func (a *ChsiAes) PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

/*
解包装
*/
func (a *ChsiAes) PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
