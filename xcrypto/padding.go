package xcrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// decrypt 函数用于解密给定的密文数据
//
// 参数：
//
//	secretData []byte - 需要解密的密文数据
//	key []byte - 解密密钥
//
// 返回值：
//
//	originByte []byte - 解密后的明文数据
//	err error - 错误信息（如果解密过程中发生错误）
//
// 功能：
//  1. 使用AES算法和提供的密钥key创建一个cipher.Block接口
//  2. 如果创建cipher.Block接口失败，则返回错误
//  3. 获取cipher.Block接口的块大小
//  4. 创建一个cipher.BlockMode接口的CBC解密模式
//  5. 创建一个与密文数据相同长度的切片originByte用于存储解密后的明文数据
//  6. 调用blockMode.CryptBlocks对密文数据进行解密，并将解密结果存储在originByte中
//  7. 如果解密后的明文数据长度为0，则返回错误
//  8. 去除解密后明文数据的PKCS5填充，并返回结果
func decrypt(secretData, key []byte) (originByte []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	originByte = make([]byte, len(secretData))
	blockMode.CryptBlocks(originByte, secretData)
	if len(originByte) == 0 {
		return nil, errors.New("blockMode.CryptBlocks error")
	}
	return PKCS5UnPadding(originByte), nil
}

// 解密填充模式（去除补全码） PKCS5UnPadding
// 解密时，需要在最后面去掉加密时添加的填充byte
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])   // 找到Byte数组最后的填充byte
	return origData[:(length - unpadding)] // 只截取返回有效数字内的byte数组
}

// 加密填充模式（添加补全码） PKCS7Padding
// 加密时，如果加密bytes的length不是blockSize的整数倍，需要在最后面添加填充byte
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	paddingCount := blockSize - len(ciphertext)%blockSize // 需要padding的数目
	paddingBytes := []byte{byte(paddingCount)}
	padtext := bytes.Repeat(paddingBytes, paddingCount) // 生成填充的文本
	return append(ciphertext, padtext...)
}

// 解密填充模式（去除补全码） PKCS7UnPadding
// 解密时，需要在最后面去掉加密时添加的填充byte
func PKCS7UnPadding(origData []byte) (bs []byte) {
	length := len(origData)
	unPaddingNumber := int(origData[length-1]) // 找到Byte数组最后的填充byte 数字
	if unPaddingNumber <= 16 {
		bs = origData[:(length - unPaddingNumber)] // 只截取返回有效数字内的byte数组
	} else {
		bs = origData
	}
	return
}
