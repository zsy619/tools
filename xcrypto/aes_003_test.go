package xcrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"testing"

	"haedu.gov.cn/tools/xuuid"
)

// 填充
func pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil
}

func encryptx(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	msg := pad([]byte(text))
	ciphertext := make([]byte, aes.BlockSize+len(msg))

	// 随机生成向量
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], msg)

	finalMsg := (base64.StdEncoding.EncodeToString(ciphertext))

	return finalMsg, nil
}

func decryptx(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	decodedMsg, err := base64.StdEncoding.DecodeString((text))
	if err != nil {
		return "", err
	}

	if (len(decodedMsg) % aes.BlockSize) != 0 {
		return "", errors.New("blocksize must be multipe of decoded message length")
	}

	iv := decodedMsg[:aes.BlockSize]
	msg := decodedMsg[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(msg, msg)

	unpadMsg, err := unpad(msg)
	if err != nil {
		return "", err
	}

	return string(unpadMsg), nil
}

func Test_Aes_003(t *testing.T) {
	for i := 0; i < 10; i++ {
		key, err := xuuid.GenerateRandomBytes(18)
		if err != nil {
			return
		}
		ky := base64.StdEncoding.EncodeToString(key)
		fmt.Println(i, ky, len(ky))
	}

	aesKey := "3PywuK7B4stmJT1Q5APOlg=="
	content := "{\"random\":\"\",\"year\":2018,\"type\":\"add\",\"ksh\":\"14410102151881\",\"sfzh\":\"411325199701010749\",\"xm\":\"卫苹到\",\"xbdm\":\"2\",\"mzdm\":\"1\",\"zzmmdm\":\"1\",\"xldm\":\"10\",\"zydm\":\"070503\",\"zyfx\":\"人文地理\",\"pyfsdm\":\"02\",\"dxhwpdw\":\"石河子大学\",\"syszddm\":\"17\",\"cxsy\":\"城市\",\"xz\":\"4\",\"rxsj\":\"20140901\",\"bysj\":\"20180612\",\"sfslbdm\":\"06\",\"knslbdm\":\"07\"}"

	encryptText, _ := encryptx([]byte(aesKey), content)

	fmt.Println("\r\nencrypt text is \n", encryptText)

	rawText, err := decryptx([]byte(aesKey), encryptText)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("raw text is  \n", rawText)
}
