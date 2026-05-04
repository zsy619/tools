package xcrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// AesEncryptCBC 使用AES算法在CBC模式下对原始数据进行加密
//
// origData: 待加密的原始数据（[]byte类型）
// key:      用于加密的密钥（[]byte类型），长度必须为16, 24或32字节
//
// 返回值:
// encrypted: 加密后的数据（[]byte类型）
func AesEncryptCBC(origData []byte, key []byte) (encrypted []byte) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()               // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize) // 补全码
	ivKey := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	blockMode := cipher.NewCBCEncrypter(block, ivKey) // 加密模式
	encrypted = make([]byte, len(origData))           // 创建数组
	blockMode.CryptBlocks(encrypted, origData)        // 加密
	return encrypted
}

// AesDecryptCBC 函数用于使用AES算法和CBC模式解密字节数组
//
// encrypted: 待解密的密文字节数组
// key: 解密使用的密钥字节数组
//
// 返回解密后的明文字节数组和可能发生的错误
func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte, err error) {
	block, err := aes.NewCipher(key) // 分组秘钥
	// blockSize := block.BlockSize()   // 获取秘钥块的长度
	ivKey := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	blockMode := cipher.NewCBCDecrypter(block, ivKey) // 加密模式
	decrypted = make([]byte, len(encrypted))          // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)       // 解密
	decrypted = pkcs5UnPadding(decrypted)             // 去除补全码
	return decrypted, err
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// pkcs5UnPadding 移除PKCS#5填充的数据
//
// 参数：
//
//	origData []byte - 原始数据，需要进行去填充
//
// 返回值：
//
//	[]byte - 移除填充后的数据
//
// 注意：
//
//	如果原始数据的长度小于填充长度，则返回空字节数组
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length < unpadding {
		return []byte{}
	}
	return origData[:(length - unpadding)]
}

// AesEncryptECB 使用ECB模式对给定的原始数据（origData）进行AES加密
// origData: 待加密的原始数据
// key: 加密密钥
// 返回值：
// encrypted: 加密后的数据
func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

// AesDecryptECB 使用ECB模式解密AES加密的数据
//
// 参数：
//
//	encrypted []byte：需要解密的密文
//	key []byte：解密密钥
//
// 返回值：
//
//	decrypted []byte：解密后的明文
//
// 注意：
//
//	该函数使用AES算法在ECB模式下解密数据，并假设密文的最后一个字节是填充字节，用于表示填充的长度
func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

// generateKey 函数接受一个字节切片 key 作为输入，并返回一个长度固定为 16 的字节切片 genKey
//
// key 参数：要生成密钥的原始字节切片
//
// 返回值：
// genKey：生成的长度为 16 的字节切片，它是基于输入 key 的字节切片生成的
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// AesEncryptCFB 函数用于对原始数据进行AES加密算法加密（CFB模式）
//
// origData: 待加密的原始数据，类型为[]byte
// key: 加密密钥，类型为[]byte
//
// 返回值：
// encrypted: 加密后的数据，类型为[]byte
//
// 注意：
// 1. 如果密钥长度不符合AES算法要求，将触发panic
// 2. 加密过程中使用随机生成的初始化向量（IV）
func AesEncryptCFB(origData []byte, key []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}

// AesDecryptCFB 函数使用AES算法和CFB模式解密给定的加密数据
//
// 参数：
// encrypted []byte - 待解密的加密数据
// key []byte - AES加密密钥
//
// 返回值：
// decrypted []byte - 解密后的数据
//
// 注意：
// 如果encrypted的长度小于AES块大小（aes.BlockSize），则会发生panic
func AesDecryptCFB(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}
