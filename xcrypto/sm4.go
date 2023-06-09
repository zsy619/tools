package xcrypto

import (
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/tjfoc/gmsm/sm4"
)

/*
国产密码算法中的 SM4 算法比较适合用于密码加密，它是一种分组加密算法，每组数据为 128 位（16 字节），
支持 128 位、192 位和 256 位三种密钥长度。SM4 算法的加密速度较慢，但是安全性较高，因此在数据的保
密性要求较高的场景下可以考虑使用 SM4 算法进行密码加密。

在使用 SM4 算法进行密码加密时，需要注意以下几个方面：

密钥管理和保密性：SM4 算法使用的是对称密钥加密，因此需要保证密钥的保密性和安全性，以避免出现密
钥泄露等安全问题。同时需要合理设置密钥的生命周期，并定期更换和更新密钥。

密码填充和向量：在加密时，需要对密码进行填充，并使用随机向量对密码加密，以增加密码的安全性。

防止攻击：在使用 SM4 算法进行密码加密时，需要注意防止常见攻击方式，例如选用高强度的密钥、使用密
码学安全的伪随机数生成器等。

在使用 SM4 算法进行加密和解密时，密钥的长度必须是 16、24 或 32 字节，否则会导致加密或解密失败。
另外，在加密时需要使用随机向量进行填充，以增加密码的安全性。
*/

/**
 * @description: SM4 PKCS7 加密
 * @param {string} data 明文
 * @param {string} key 密钥
 * @return {*}
 */
func Sm4Pkcs7Encrypt(data string, key string) (string, error) {
	// 密码填充、向量和密钥
	dataByte := []byte(data)
	keyHash := sha256.Sum256([]byte(key))
	key16 := keyHash[0:16]

	// 加密
	block, err := sm4.NewCipher(key16)
	if err != nil {
		return "", fmt.Errorf("sm4.NewCipher error: %v", err)
	}
	bsize := block.BlockSize()
	plaintextBytes := []byte(dataByte)
	plaintextBytes = PKCS7Padding(plaintextBytes, bsize)

	ciphertext := make([]byte, len(plaintextBytes))
	iv := make([]byte, bsize)
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintextBytes)
	outPwd := hex.EncodeToString(ciphertext)
	fmt.Printf("密文: %s\n", outPwd)
	return outPwd, nil
}

/**
 * @description: SM4 PKCS7 解密
 * @param {string} data 密文
 * @param {string} key 密钥
 * @return {*}
 */
func Sm4Pkcs7Decrypt(data string, key string) (string, error) {
	keyHash := sha256.Sum256([]byte(key))
	key16 := keyHash[0:16]
	block, err := sm4.NewCipher(key16)
	if err != nil {
		return "", err
	}
	ciphertext, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}
	bsize := block.BlockSize()
	iv := make([]byte, bsize)
	decrypted := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, ciphertext)
	decrypted = PKCS7UnPadding(decrypted)
	outStr := string(decrypted)
	fmt.Printf("明文: %s\n", outStr)
	return outStr, nil
}

/**
 * @description: SM4 加密
 * @param {string} data 明文
 * @param {string} key 密钥
 * @return {*}
 */
func Sm4Encrypt(data string, key string) (string, error) {
	// 密码填充、向量和密钥
	keyHash := sha256.Sum256([]byte(key))
	key16 := keyHash[0:16]
	// 加密
	block, err := sm4.NewCipher(key16)
	if err != nil {
		return "", fmt.Errorf("sm4.NewCipher error: %v", err)
	}
	// 定义加密模式为 CBC
	// 加密和解密的数据长度必须是 SM4 块大小的整数倍。
	// 如果数据长度不足 SM4 块大小的整数倍，可以使用 Zero Padding 将数据长度补足。
	mode := cipher.NewCBCEncrypter(block, keyHash[:sm4.BlockSize])
	datax := ZeroPadding([]byte(data), sm4.BlockSize)
	encrypted := make([]byte, len(datax))
	mode.CryptBlocks(encrypted, datax)
	outPwd := hex.EncodeToString(encrypted)
	fmt.Printf("加密后的密文: %s\n", outPwd)
	return outPwd, nil
}

/**
 * @description: SM4 解密
 * @param {string} data 密文
 * @param {string} key 密钥
 * @return {*}
 */
func Sm4Decrypt(data string, key string) (string, error) {
	keyHash := sha256.Sum256([]byte(key))
	key16 := keyHash[0:16]
	block, err := sm4.NewCipher(key16)
	if err != nil {
		return "", err
	}
	ciphertext, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}
	decrypted := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, keyHash[:sm4.BlockSize])
	mode.CryptBlocks(decrypted, ciphertext)
	outStr := string(ZeroUnPadding(decrypted))
	fmt.Printf("解密后的明文: %s\n", outStr)
	return outStr, nil
}

func ZeroPadding(data []byte, blocksize int) []byte {
	paddingLen := blocksize - len(data)%blocksize
	padding := make([]byte, paddingLen)
	return append(data, padding...)
}

func ZeroUnPadding(data []byte) []byte {
	length := len(data)
	i := length - 1
	for ; i >= 0; i-- {
		if data[i] != 0 {
			break
		}
	}
	return data[0 : i+1]
}
