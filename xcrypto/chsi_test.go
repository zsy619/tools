package xcrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"testing"

	"haedu.gov.cn/tools/xuuid"
)

func Test_LIBS(t *testing.T) {
	key := xuuid.GetAes()
	rtx := base64.StdEncoding.EncodeToString([]byte(key))
	fmt.Println(rtx)

	aes, err := NewChsiAes([]byte("xgckg6ps2eoshvin7xcb5xr710dkw29v"))
	if err != nil {
		return
	}
	rt, err := aes.Encrypt("aaa")
	if err != nil {
		return
	}
	fmt.Println(string(rt))

	data := `{"appName":"就业数据上报系统","custNum":"xgckg6ps2eoshvin7xcb5xr710dkw29v","mode":0,"timeStamp":"20210427100713137"}`

	fmt.Println(base64.StdEncoding.EncodeToString([]byte(data)))

	rsax := new(ChsiRsa)
	rsax.SetPrvKey([]byte(FormatPrivateKey("MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAI/lL2TG10qtjxI4ruBVNyxsk+YsXQu1LYKjJzJZtERMzHJhVWU9s/8o/NtmPjjA61jH7f4HCj5QL14ync79RNBq3rod+CVFfx9BSZmh3ryz2JBSW5Sejfz8dMq6t96Ku1upIDBSPIbtD+VwdN4AO1U4po3t77/khABrRyin39d1AgMBAAECgYAL8/nP4USZC3nLBbJhKDMBGbPfdufzxQUWyP7Ei/cRhV+mULeLRWjiVUFL6F5a0Iu8QD9gzqznKDoHFSVOwHMqzPWIyq9rQulGeSNM+hxCqZTq2qo3PZ5+P8twWQ9yrX+fl/2aDQP4doRXYjF/Vuf0c8phOmECY/0MMtalc2FEgQJBANDktXp6HM4fDiVjI3Q7Gv1fsMiNN4yNHNy2FFJnajupMN+Aaf6fG8VtZu3M2jNDUAVCkwWJ/Uq0VsXoV1/rnd0CQQCwWCwg8W7enYJrAXK55r74I0TBDmo+sGBMbsO9J9ZzdB1DREN55EL5IxAqleKtNq/REhwqUHgRR3i/AYH/OIJ5AkEAv1Fo4OasMSACPb3Bv/dOLdcRO20S/jhTwdVFYX9zrXa0205qRZiFv9kGFy+yfJbe2CJ0MvOBt4TZoGK+e4x5RQJAasGGPY9L2lMqkBM5XBe4Bsp7JhDO+xKVyc/IievjJNPnn0BlRRaOAPtcHxvMNaaEu6ImEOvUNEm7bI7CHzsbqQJBALeShyUZRezxacl8Ez3+OU0z0L+Dfas6fGxSu2NbH6NtiimUvv7BYx/Wp+1Et5BnDbIzHIxcgDlL0/dF7LNgdww=")))
	signData, err := rsax.Sign(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(signData)
	// O3oBrWwHXPcyzxPJuy2zQ7Tl4cE9jRdCiYnPJ8LbLbIRuK3hVwTxafhygRcL7HRztGvy48Lp2c9z+hemSmLcQPe0qH5nC+IG/DtTMD5hL5ID0kN0AXA7bZEjxZl/ZNWLnKHHN2IiMehOpeznN8HhyfCScgixSswPP3SmBlGTRfw=
	// sdx := base64.StdEncoding.EncodeToString([]byte(signData))
	// fmt.Println(sdx)
}

func Test_AES(t *testing.T) {
	// taking a string
	givenString := "3PywuK7B4stmJT1Q5APOlg=="

	// using the function
	decodedString, err := base64.RawURLEncoding.DecodeString(givenString)
	if err != nil {
		t.Fatal(err)
	}
	buf := new(bytes.Buffer)

	err = binary.Write(buf, binary.BigEndian, decodedString)

	if err != nil {
		fmt.Println("Error Found:", err)
		return
	}

	fmt.Print("Decoded Bytes: ")
	fmt.Println(decodedString)

	fmt.Print("Decoded String: ")
	fmt.Println(string(decodedString))

	// aesxx, err := GenerateRandomKey(16)
	// if err != nil {
	// 	fmt.Println("错误03：", err.Error())
	// 	return
	// }
	// fmt.Println("----->", aesxx, len(aesxx))
	// return

	// 8HAb6jYbxcaSRpcdSVUaOg==
	// [B@282003e1
	tkey := "8HAb6jYbxcaSRpcdSVUaOg=="
	tkeyb, _ := base64.StdEncoding.DecodeString(tkey)
	fmt.Println("-----> ", tkeyb, len(tkeyb), len(fmt.Sprintf("%x", tkeyb)))
	// return

	prvKey := []byte(FormatPrivateKey("MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJvbS7UNce/07ShHiyb+1ooBHAH1HwVq7DnzU2NAfR5BMXNG2jJ2Ag8u79Y4QZFpj/bHr5End/8MbzijN1jif8DhzMyUcf3+ITYZiHDRuhCnD3ef8dVnEQeqyKrlKsffcTPp+vJgZRiOqbGZZzuUfPpWbIVB0ZxnvwYXQ2GLZ161AgMBAAECgYArXi4GxyL5HjIPjzjNNQQFiqF8efST0VjCF08QwxUNoh5ccU6t0+Bm0SyzcxvrlnAUvyO/RDhDo/Ye0GvKM9xQJAXJgO+3RZ//xEwspz56Sjc+ONiQIrurwLTtlEBmTucGLSnT9GVOvXmB5NTFGTwm7qb+R0W1bW86O0KJ2B4egQJBANfJz9pTluDHf/XTlggep5xIFVZpf3v2O2sjRioqyN3d3e9fELniif8cZQxbMH+9cofabm4D5DLNBWg7+PXsJ00CQQC45nCzQEKKWl2jgsDMCKlwgYYcg+LCVy7nTvFDl3XHYfZni4Z7wpVCkCvzGF6qCmEwz5AyluecmA0Am6LVRXEJAkBPqjbtUHzcQWrRU6sJFmAkx0vxWgNxvWcUV7J4sND1cAqWa89eAO+XWmFH3YabMlLNKuwn+5HM23oKkFGKYQPlAkAgMad/3nl3g4J4XOTa4cs21qaWQnRyKCH3jmw9u5p7S9hOcSHKXLgGbfnpCt44tzPy/sD5vgK35lWlPHQetEeZAkAxjMoK0XNplKP+kVxYXSWnH6Jf02Mgkqu36iDSsggs5/lmEvJ+xgSPGMcqs7jyA2XlZvEG6kp7M716oqd08+e1"))
	pubKey := []byte(getPublicKey("MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCb20u1DXHv9O0oR4sm/taKARwB9R8Fauw581NjQH0eQTFzRtoydgIPLu/WOEGRaY/2x6+RJ3f/DG84ozdY4n/A4czMlHH9/iE2GYhw0boQpw93n/HVZxEHqsiq5SrH33Ez6fryYGUYjqmxmWc7lHz6VmyFQdGcZ78GF0Nhi2detQIDAQAB"))

	aesKey := "3PywuK7B4stmJT1Q5APOlg=="
	fmt.Println(aesKey, len(aesKey))
	rsax, err := NewChsiRSA(pubKey, prvKey)
	if err != nil {
		t.Fatal(err)
	}
	rsaEncrypt, err := rsax.PublicEncrypt(aesKey)
	if err != nil {
		return
	}
	fmt.Println("\r\n1 使用公钥对AES进行RSA加密:", string(rsaEncrypt))
	rsaEncrypt = "bz6WgMSChn4evM63EAqzJFQFh5s88ZMHp+ymrEeAglsUb2IF0Kdnz4lqXc8/6n+UpHPvE2+50jv6dUojBlHxysUFDqsi7FRpgORHls39S1+yNGJK6IvpinZUfmUJCmPVrmtIks+1bq+h2dQFLr8On6oE2mFjZBpT52gPM0mI7TU="
	rsaDecrypt, err := rsax.PrivateDecrypt(rsaEncrypt)
	if err != nil {
		return
	}
	fmt.Println("1 使用私钥对AES进行RSA解密:", rsaDecrypt)

	rsaEncryptPrivateKey, err := rsax.PrivateEncrypt(aesKey)
	if err != nil {
		return
	}
	fmt.Println("\r\n2 使用私钥对AES进行RSA加密:", rsaEncryptPrivateKey)
	rsaDecryptPublicKey, err := rsax.PublicDecrypt(rsaEncryptPrivateKey)
	if err != nil {
		return
	}
	fmt.Println("2 使用公钥对AES进行RSA解密:", rsaDecryptPublicKey)

	fmt.Println("\r\n---------------------AES 加/解密测试---------------------")
	content := "{\"random\":\"\",\"year\":2018,\"type\":\"add\",\"ksh\":\"14410102151881\",\"sfzh\":\"411325199701010749\",\"xm\":\"卫苹到\",\"xbdm\":\"2\",\"mzdm\":\"1\",\"zzmmdm\":\"1\",\"xldm\":\"10\",\"zydm\":\"070503\",\"zyfx\":\"人文地理\",\"pyfsdm\":\"02\",\"dxhwpdw\":\"石河子大学\",\"syszddm\":\"17\",\"cxsy\":\"城市\",\"xz\":\"4\",\"rxsj\":\"20140901\",\"bysj\":\"20180612\",\"sfslbdm\":\"06\",\"knslbdm\":\"07\"}"

	// 加密方式 一

	aesx, err := NewChsiAes([]byte(rsaDecrypt))
	if err != nil {
		return
	}
	aesEncrypt, err := aesx.Encrypt(content)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("AES 加密:\r\n", aesEncrypt)
	// aesEncrypt = "D+L8LDiZ05nOD1i7AmVl8ytEbQ++urjehGP5JIiqNmkzYV46m3CyQZ9VvHO8xxvhb28xnYktIboOEUbvj6B1OYjuMW1McZ3DkhbKiObxJDlTo4i58wwvLSQkc71+OJN2YjCvl7LUzFvyXP+SOIc+B6R8xsZ2PkqvXKWfDMW+psf/K+tm37nE0URLGeL14tlNroJTLYFrmNc4F6p2n7P5NlgvLwMjQkpqc2fbMJYxkHNYBSRrPvaELtUlTLA41D27hOVHQmWShfPwyX2Jhqvlgamteu4Z5JGsP8ZOtaeGO/mZwdeoIXTmMKgLzizhSWw/mIy1gvnxxOkLwbdty5OxZ1/Z8l5aZ5zPC7xskWrjlwRwkgla2xfMNnf9x83cFRIRQkU60MQiO2FNiEu+lQspwQXHxR0gxP8ohlef5Jls2JUBxMW/OQjKCpqvaHdr+EDbY7Nou03uFPzcfREb6581JA=="
	aesDecrypt, err := aesx.Decrypt(aesEncrypt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("AES 解密:\r\n", aesDecrypt)

	// 加密方式 二

	// rtCBC := AesEncryptECB([]byte(content), []byte(rsaDecrypt))
	// fmt.Println("======》：", base64.StdEncoding.EncodeToString(rtCBC))
	// rtCBC = AesDecryptECB(rtCBC, []byte(rsaDecrypt))
	// if err != nil {
	// 	return
	// }
	// fmt.Println("======》：", base64.StdEncoding.EncodeToString(rtCBC))

	// 加密方式 三

	// ret, err := Encrypt(rsaDecrypt, content)
	// if err != nil {
	// 	return
	// }
	// fmt.Println("--->:", ret)

	// // ret = "D+L8LDiZ05nOD1i7AmVl8ytEbQ++urjehGP5JIiqNmkzYV46m3CyQZ9VvHO8xxvhb28xnYktIboOEUbvj6B1OYjuMW1McZ3DkhbKiObxJDlTo4i58wwvLSQkc71+OJN2YjCvl7LUzFvyXP+SOIc+B6R8xsZ2PkqvXKWfDMW+psf/K+tm37nE0URLGeL14tlNroJTLYFrmNc4F6p2n7P5NlgvLwMjQkpqc2fbMJYxkHNYBSRrPvaELtUlTLA41D27hOVHQmWShfPwyX2Jhqvlgamteu4Z5JGsP8ZOtaeGO/mZwdeoIXTmMKgLzizhSWw/mIy1gvnxxOkLwbdty5OxZ1/Z8l5aZ5zPC7xskWrjlwRwkgla2xfMNnf9x83cFRIRQkU60MQiO2FNiEu+lQspwQXHxR0gxP8ohlef5Jls2JUBxMW/OQjKCpqvaHdr+EDbY7Nou03uFPzcfREb6581JA=="
	// ret, err = Decrypt(rsaDecrypt, ret)
	// if err != nil {
	// 	return
	// }
	// fmt.Println("--->:", ret)
}

func getKeyBytes(key string) []byte {
	keyBytes := []byte(key)
	switch l := len(keyBytes); {
	case l < 16:
		keyBytes = append(keyBytes, make([]byte, 16-l)...)
	case l > 16:
		keyBytes = keyBytes[:16]
	}
	return keyBytes
}

func encrypt(key string, origData []byte) ([]byte, error) {
	keyBytes := getKeyBytes(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, keyBytes[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func decrpt(key string, crypted []byte) ([]byte, error) {
	keyBytes := getKeyBytes(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyBytes[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// func PKCS5UnPadding(origData []byte) []byte {
// 	length := len(origData)
// 	unpadding := int(origData[length-1])
// 	return origData[:(length - unpadding)]
// }

func Encrypt(key string, val string) (string, error) {
	origData := []byte(val)
	crypted, err := encrypt(key, origData)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func Decrypt(key string, val string) (string, error) {
	crypted, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return "", err
	}
	origData, err := decrpt(key, crypted)
	if err != nil {
		return "", err
	}
	return string(origData), nil
}
