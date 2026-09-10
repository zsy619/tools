package xcrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

// ChsiAes 封装了使用 AES 对称加解密所需的密钥。
type ChsiAes struct {
	key []byte
}

// NewChsiAes 函数用于创建一个新的ChsiAes实例
//
// 参数：
//
//	key []byte - 用于AES加密的密钥
//
// 返回值：
//
//	*ChsiAes - 指向新创建的ChsiAes实例的指针
//	error - 如果密钥为nil，则返回非零错误码，否则返回nil
func NewChsiAes(key []byte) (*ChsiAes, error) {
	if key == nil {
		return nil, errors.New("key错误")
	}
	return &ChsiAes{
		key: key,
	}, nil
}

// Encrypt 使用AES加密算法对字符串src进行加密
// 参数：
//
//	a: *ChsiAes ChsiAes类型的指针，包含加密所需的key
//	src: string 待加密的字符串
//
// 返回值：
//
//	string 加密后的字符串（base64编码）
//	error 加密过程中出现的错误，如果没有错误则为nil
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
	// Encrypt 使用AES加密算法对内容进行加密
	// 加密后的结果使用Base64编码
	// ivKey 是初始化向量，此处使用了16个0
	// aesBlockEncrypter 是AES加密块
	// content 是待加密的内容
	// encrypted 是加密后的结果
	// err 是加密过程中可能发生的错误
	ivKey := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	aesEncrypter := cipher.NewCBCEncrypter(aesBlockEncrypter, ivKey)
	// aesEncrypter := cipher.NewCBCEncrypter(aesBlockEncrypter, a.key[:aesBlockEncrypter.BlockSize()])
	aesEncrypter.CryptBlocks(encrypted, content)

	return base64.StdEncoding.EncodeToString(encrypted), err
}

// Decrypt 使用AES算法解密给定的加密字符串
//
// 参数：
//
//	a *ChsiAes - ChsiAes结构体指针，包含解密所需的密钥
//	crypt string - 待解密的加密字符串
//
// 返回值：
//
//	data string - 解密后的字符串
//	err error - 解密过程中出现的错误，如果解密成功则为nil
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

// PKCS5UnPadding 从给定的 origData 字节切片中去除 PKCS#5 填充
// 参数 origData: 需要去除填充的字节切片
// 返回值: 去除填充后的字节切片
func (a *ChsiAes) PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// PKCS5Padding 使用PKCS#5填充方式对密文进行填充
//
// 参数：
//
//	a *ChsiAes：ChsiAes类型的指针，用于调用该方法的实例
//	cipherText []byte：待填充的密文
//	blockSize int：块大小
//
// 返回值：
//
//	[]byte：填充后的密文
func (a *ChsiAes) PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// PKCS5Trimming 从PKCS5填充的加密字节切片中移除填充部分
// 参数:
//
//	encrypt []byte - 带有PKCS5填充的加密字节切片
//
// 返回值:
//
//	[]byte - 移除PKCS5填充后的加密字节切片
func (a *ChsiAes) PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
