package xcrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"testing"
)

func Test_Aes_002(t *testing.T) {
	aesKey := "3PywuK7B4stmJT1Q5APOlg=="
	content := "{\"random\":\"\",\"year\":2018,\"type\":\"add\",\"ksh\":\"14410102151881\",\"sfzh\":\"411325199701010749\",\"xm\":\"卫苹到\",\"xbdm\":\"2\",\"mzdm\":\"1\",\"zzmmdm\":\"1\",\"xldm\":\"10\",\"zydm\":\"070503\",\"zyfx\":\"人文地理\",\"pyfsdm\":\"02\",\"dxhwpdw\":\"石河子大学\",\"syszddm\":\"17\",\"cxsy\":\"城市\",\"xz\":\"4\",\"rxsj\":\"20140901\",\"bysj\":\"20180612\",\"sfslbdm\":\"06\",\"knslbdm\":\"07\"}"

	txt := En(content, aesKey)
	log.Println(txt)
	// txt = "D+L8LDiZ05nOD1i7AmVl8ytEbQ++urjehGP5JIiqNmkzYV46m3CyQZ9VvHO8xxvhb28xnYktIboOEUbvj6B1OYjuMW1McZ3DkhbKiObxJDlTo4i58wwvLSQkc71+OJN2YjCvl7LUzFvyXP+SOIc+B6R8xsZ2PkqvXKWfDMW+psf/K+tm37nE0URLGeL14tlNroJTLYFrmNc4F6p2n7P5NlgvLwMjQkpqc2fbMJYxkHNYBSRrPvaELtUlTLA41D27hOVHQmWShfPwyX2Jhqvlgamteu4Z5JGsP8ZOtaeGO/mZwdeoIXTmMKgLzizhSWw/mIy1gvnxxOkLwbdty5OxZ1/Z8l5aZ5zPC7xskWrjlwRwkgla2xfMNnf9x83cFRIRQkU60MQiO2FNiEu+lQspwQXHxR0gxP8ohlef5Jls2JUBxMW/OQjKCpqvaHdr+EDbY7Nou03uFPzcfREb6581JA=="
	// rt, _ := base64.RawStdEncoding.DecodeString(txt)
	var source string = UnEn(txt, aesKey)
	log.Println(source)
}

// 填充
func paddingkey(key string) string {
	var buffer bytes.Buffer
	buffer.WriteString(key)

	for i := len(key); i < 16; i++ {
		buffer.WriteString("0")
	}

	return buffer.String()
}

// 加密
func En(src string, srckey string) string {
	key := []byte(paddingkey(srckey))
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	result, err := AesEncryptx([]byte(src), key, iv)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(result)
}

// 解密
func UnEn(src string, srckey string) string {
	key := []byte(paddingkey(srckey))

	var result []byte
	var err error

	result, err = base64.StdEncoding.DecodeString(src)
	if err != nil {
		panic(err)
	}
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	origData, err := AesDecryptx(result, key, iv)
	if err != nil {
		panic(err)
	}
	return string(origData)
}

func AesEncryptx(origData, key []byte, IV []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, IV[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecryptx(crypted, key []byte, IV []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, IV[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}
