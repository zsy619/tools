package xcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"testing"
)

func Test_Aes_001(t *testing.T) {
	aesKey := "3PywuK7B4stmJT1Q5APOlg=="
	content := "{\"random\":\"\",\"year\":2018,\"type\":\"add\",\"ksh\":\"14410102151881\",\"sfzh\":\"411325199701010749\",\"xm\":\"卫苹到\",\"xbdm\":\"2\",\"mzdm\":\"1\",\"zzmmdm\":\"1\",\"xldm\":\"10\",\"zydm\":\"070503\",\"zyfx\":\"人文地理\",\"pyfsdm\":\"02\",\"dxhwpdw\":\"石河子大学\",\"syszddm\":\"17\",\"cxsy\":\"城市\",\"xz\":\"4\",\"rxsj\":\"20140901\",\"bysj\":\"20180612\",\"sfslbdm\":\"06\",\"knslbdm\":\"07\"}"

	txt, err := AESBase64Encrypt(content, aesKey)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(txt)
	// txt = "D+L8LDiZ05nOD1i7AmVl8ytEbQ++urjehGP5JIiqNmkzYV46m3CyQZ9VvHO8xxvhb28xnYktIboOEUbvj6B1OYjuMW1McZ3DkhbKiObxJDlTo4i58wwvLSQkc71+OJN2YjCvl7LUzFvyXP+SOIc+B6R8xsZ2PkqvXKWfDMW+psf/K+tm37nE0URLGeL14tlNroJTLYFrmNc4F6p2n7P5NlgvLwMjQkpqc2fbMJYxkHNYBSRrPvaELtUlTLA41D27hOVHQmWShfPwyX2Jhqvlgamteu4Z5JGsP8ZOtaeGO/mZwdeoIXTmMKgLzizhSWw/mIy1gvnxxOkLwbdty5OxZ1/Z8l5aZ5zPC7xskWrjlwRwkgla2xfMNnf9x83cFRIRQkU60MQiO2FNiEu+lQspwQXHxR0gxP8ohlef5Jls2JUBxMW/OQjKCpqvaHdr+EDbY7Nou03uFPzcfREb6581JA=="
	// rt, _ := base64.RawStdEncoding.DecodeString(txt)
	var source string
	source, err = AESBase64Decrypt(txt, aesKey)
	if err != nil {
		log.Println(err)
	}
	log.Println(source)
}

func AESBase64Encrypt(origin_data string, key string) (base64_result string, err error) {
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var block cipher.Block
	if block, err = aes.NewCipher([]byte(key)); err != nil {
		log.Println(err)
		return
	}
	encrypt := cipher.NewCBCEncrypter(block, iv)
	var source []byte = PKCS5Padding([]byte(origin_data), 16)
	var dst []byte = make([]byte, len(source))
	encrypt.CryptBlocks(dst, source)
	base64_result = base64.StdEncoding.EncodeToString(dst)
	return
}

func AESBase64Decrypt(encrypt_data string, key string) (origin_data string, err error) {
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var block cipher.Block
	if block, err = aes.NewCipher([]byte(key)); err != nil {
		log.Println(err)
		return
	}
	encrypt := cipher.NewCBCDecrypter(block, iv)

	var source []byte
	if source, err = base64.StdEncoding.DecodeString(encrypt_data); err != nil {
		log.Println(err)
		return
	}
	var dst []byte = make([]byte, len(source))
	encrypt.CryptBlocks(dst, source)
	origin_data = string(PKCS5Unpadding(dst))
	return
}

// func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
// 	padding := blockSize - len(ciphertext)%blockSize
// 	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
// 	return append(ciphertext, padtext...)
// }

func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
