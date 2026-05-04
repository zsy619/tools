package xcrypto

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func FormatPrivateKey(key string) string {
	rsa := ""
	temp := strings.ReplaceAll(key, "\n", "")
	allLen := len(temp)
	for i := 0; i <= allLen; i = i + 64 {
		if allLen-i < 64 {
			rsa += "\n" + temp[i:]
		} else {
			rsa += "\n" + temp[i:i+64]
		}
	}

	rsa = "-----BEGIN RSA PRIVATE KEY-----" + rsa + "\n-----END RSA PRIVATE KEY-----"
	// fmt.Println(strings.TrimRight(rsa, ""))
	return rsa
}

func getPublicKey(key string) string {
	rsa := ""
	temp := strings.ReplaceAll(key, "\n", "")
	allLen := len(temp)
	for i := 0; i <= allLen; i = i + 64 {
		if allLen-i < 64 {
			rsa += "\n" + temp[i:]
		} else {
			rsa += "\n" + temp[i:i+64]
		}
	}

	rsa = "-----BEGIN PUBLIC KEY-----" + rsa + "\n-----END PUBLIC KEY-----"
	// fmt.Println(strings.TrimRight(rsa, ""))
	return rsa
}

func Test_RsaXXXX(t *testing.T) {
	data := `{"appName":"就业数据上报系统","custNum":"xgckg6ps2eoshvin7xcb5xr710dkw29v","mode":0,"timeStamp":"20210426184141582"}`
	// 消息的签名信息： fOZTWBlpbZ98nX29zndmO8aBRfzILTZpw_qVO8_z-zWtsvWPuUihi6sYx0eYyqPSS6H3QEZpOl03lRIG4wnpBnrlNb3lP7Rg-FymAELwnxbMMWHeFEJLMf8rrLTV1Bpc8yCdoIeh-xAmAcoFMjhoDiF6wApPh1iyBGlYpuOfRHA
	prvKey := FormatPrivateKey("MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAI/lL2TG10qtjxI4ruBVNyxsk+YsXQu1LYKjJzJZtERMzHJhVWU9s/8o/NtmPjjA61jH7f4HCj5QL14ync79RNBq3rod+CVFfx9BSZmh3ryz2JBSW5Sejfz8dMq6t96Ku1upIDBSPIbtD+VwdN4AO1U4po3t77/khABrRyin39d1AgMBAAECgYAL8/nP4USZC3nLBbJhKDMBGbPfdufzxQUWyP7Ei/cRhV+mULeLRWjiVUFL6F5a0Iu8QD9gzqznKDoHFSVOwHMqzPWIyq9rQulGeSNM+hxCqZTq2qo3PZ5+P8twWQ9yrX+fl/2aDQP4doRXYjF/Vuf0c8phOmECY/0MMtalc2FEgQJBANDktXp6HM4fDiVjI3Q7Gv1fsMiNN4yNHNy2FFJnajupMN+Aaf6fG8VtZu3M2jNDUAVCkwWJ/Uq0VsXoV1/rnd0CQQCwWCwg8W7enYJrAXK55r74I0TBDmo+sGBMbsO9J9ZzdB1DREN55EL5IxAqleKtNq/REhwqUHgRR3i/AYH/OIJ5AkEAv1Fo4OasMSACPb3Bv/dOLdcRO20S/jhTwdVFYX9zrXa0205qRZiFv9kGFy+yfJbe2CJ0MvOBt4TZoGK+e4x5RQJAasGGPY9L2lMqkBM5XBe4Bsp7JhDO+xKVyc/IievjJNPnn0BlRRaOAPtcHxvMNaaEu6ImEOvUNEm7bI7CHzsbqQJBALeShyUZRezxacl8Ez3+OU0z0L+Dfas6fGxSu2NbH6NtiimUvv7BYx/Wp+1Et5BnDbIzHIxcgDlL0/dF7LNgdww=")
	// prvKey := pt()
	pubKey := getPublicKey("MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCByCBkUwMIOfmPgTo/5Ygv9UjtwR1oUFHobwrVvhJmXgnxXqF1UKPEZ54dFFgFRL1icn789Kk3GJ0kPTbG95+ZJlPLDl98LIh+GVeSb2yqYzZFcI28LRF/UK726Wcv5cTgdTYr9fU4ymsL+xCr3K8MOxCvY62v0wk4lDnSgHXehQIDAQAB")
	xrsa, err := NewXRsa([]byte(pubKey), []byte(prvKey))
	if err != nil {
		return
	}
	rt, err := xrsa.Sign(data)
	if err != nil {
		return
	}
	fmt.Println(rt)
	err = xrsa.Verify(data, rt)
	if err != nil {
		fmt.Println(err.Error())
		t.Error(err)
	}
}

// 对应java的TestXrsa示例
func Test_RsaXXXXDDD(t *testing.T) {
	data := `{"appName":"就业数据上报系统","custNum":"xgckg6ps2eoshvin7xcb5xr710dkw29v","mode":0,"timeStamp":"20210426184141582"}`
	prvKey := getPrvKey()
	pubKey := getPubKey()
	xrsa, err := NewXRsa([]byte(pubKey), []byte(prvKey))
	if err != nil {
		return
	}
	rt, err := xrsa.Sign(data)
	if err != nil {
		return
	}
	fmt.Println(rt)
	err = xrsa.Verify(data, rt)
	if err != nil {
		fmt.Println(err.Error())
		t.Error(err)
	}
}

func Test_Rsa(t *testing.T) {
	// rsa 密钥文件产生
	fmt.Println("-------------------------------获取RSA公私钥-----------------------------------------")
	prvKey, pubKey, _ := GenRsaKey()
	fmt.Println(string(prvKey))
	fmt.Println(string(pubKey))

	xrsa, err := NewXRsa(pubKey, prvKey)
	if err != nil {
		return
	}

	fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
	data := "卧了个槽，这么神奇的吗？？！！！  ԅ(¯﹃¯ԅ) ！！！！！！）"
	fmt.Println("对消息进行签名操作...")
	signData, err := xrsa.Sign(data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("消息的签名信息： ", hex.EncodeToString([]byte(signData)))
	fmt.Println("\n对签名信息进行验证...")
	if err := xrsa.Verify(data, signData); err == nil {
		fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
	}

	fmt.Println("-------------------------------进行加密解密操作-----------------------------------------")
	ciphertext, _ := xrsa.EncryptWithPublicKey(data)
	fmt.Println("公钥加密后的数据：", hex.EncodeToString([]byte(ciphertext)))
	sourceData, _ := xrsa.DecryptWithPrivateKey(ciphertext)
	fmt.Println("私钥解密后的数据：", sourceData)
}

func Test_Rsax(t *testing.T) {
	origData := []byte(`{"year":2019,"ksh":"1071016141305507"}`) // 待加密的数据
	key := []byte("d235788d00bfef3a373bdb8c708f8f92")            // AES32位加密的密钥
	prvKey := getPrvKey()
	pubKey := getPubKey()
	xrsa, err := NewXRsa([]byte(pubKey), []byte(prvKey))
	if err != nil {
		return
	}

	fmt.Println("-------------------------------Key加密-----------------------------------------")
	keyEn, _ := xrsa.EncryptWithPublicKey(string(key))
	fmt.Println("Key加密：", base64.StdEncoding.EncodeToString([]byte(keyEn)))

	fmt.Println("------------------ CBC模式 --------------------")
	encrypted := AesEncryptCBC(origData, key)
	fmt.Println("密文(hex)：", hex.EncodeToString(encrypted))
	fmt.Println("Data密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted, _ := AesDecryptCBC(encrypted, key)
	fmt.Println("解密结果：", string(decrypted))

	fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
	data := base64.StdEncoding.EncodeToString(encrypted)
	fmt.Println("对消息进行签名操作...")
	signData, _ := xrsa.Sign(data)
	fmt.Println("消息的签名信息：", base64.StdEncoding.EncodeToString([]byte(signData)))

	fmt.Println("\n对签名信息进行验证...")
	if err := xrsa.Verify(data, signData); err == nil {
		fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
	}
}

func pt() string {
	return `-----BEGIN RSA PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAI/lL2TG10qtjxI4
ruBVNyxsk+YsXQu1LYKjJzJZtERMzHJhVWU9s/8o/NtmPjjA61jH7f4HCj5QL14y
nc79RNBq3rod+CVFfx9BSZmh3ryz2JBSW5Sejfz8dMq6t96Ku1upIDBSPIbtD+Vw
dN4AO1U4po3t77/khABrRyin39d1AgMBAAECgYAL8/nP4USZC3nLBbJhKDMBGbPf
dufzxQUWyP7Ei/cRhV+mULeLRWjiVUFL6F5a0Iu8QD9gzqznKDoHFSVOwHMqzPWI
yq9rQulGeSNM+hxCqZTq2qo3PZ5+P8twWQ9yrX+fl/2aDQP4doRXYjF/Vuf0c8ph
OmECY/0MMtalc2FEgQJBANDktXp6HM4fDiVjI3Q7Gv1fsMiNN4yNHNy2FFJnajup
MN+Aaf6fG8VtZu3M2jNDUAVCkwWJ/Uq0VsXoV1/rnd0CQQCwWCwg8W7enYJrAXK5
5r74I0TBDmo+sGBMbsO9J9ZzdB1DREN55EL5IxAqleKtNq/REhwqUHgRR3i/AYH/
OIJ5AkEAv1Fo4OasMSACPb3Bv/dOLdcRO20S/jhTwdVFYX9zrXa0205qRZiFv9kG
Fy+yfJbe2CJ0MvOBt4TZoGK+e4x5RQJAasGGPY9L2lMqkBM5XBe4Bsp7JhDO+xKV
yc/IievjJNPnn0BlRRaOAPtcHxvMNaaEu6ImEOvUNEm7bI7CHzsbqQJBALeShyUZ
Rezxacl8Ez3+OU0z0L+Dfas6fGxSu2NbH6NtiimUvv7BYx/Wp+1Et5BnDbIzHIxc
gDlL0/dF7LNgdww=
-----END RSA PRIVATE KEY-----`
}

func getPrvKey() string {
	return `-----BEGIN PRIVATE KEY-----
MIIEugIBADALBgkqhkiG9w0BAQEEggSmMIIEogIBAAKCAQEAl9WjS3e+IaoA8XrV
fQyznczQv0bpKoN84W7NUnBZl30VG8WFix79Ir+msazoh18uivzIUHPcY0+jHMQc
/Qjo+S/5FOC9LBURCtla3QU3ty3zn0i+cKW6lk6c4SbOAP3zHVsf3l0rGswBxND9
x4QQICiBx1IxS6NlSeSDlwBerkJCAblEF6pTGBwZDYL/eJ0eTPH9U4ismH5OYs4w
snRlBGJUKj7YZR4CvkTErygSNxataaLuP1Qqp8Spr2cksnOfC1/tYy/N0lydpQN9
UfL6NqpKiGg6Rsp4k1skiKR535Eue8YJb0RqQ7kjr9H4F18cIwfI5iz9vwSQfmZx
mqgENwIDAQABAoIBAG8khm0G0RnJXPlnBgGMm6qGM8Pgf2uMZoyKVCflb9+RQzNa
CiBFZdza14W14Vy+ks5Qrb0eopPbxrWW5PVgYVGPCVB8Fl2/agM8CeRCHn+rVmsh
j63b0tKV5wZ1JlTZj+3MN27JWnU6Io1UwoAarscrf5xNESKiD9HgQWb2cVgyrXXJ
MvqaXATutIwZdwfkLADuFg06DBTC6pFb8rB3TWSzY8x9VDztU66ecL/B8Mcn3NM9
pzJjKXdvQRgQdud9GvQrY3EdzzmhFb74fgtUwZ20dkON1Zfj4m7y2ApB45slkdKi
4LclRqtaFobCq+hywKIScLOI/S1f+FhOCf9kC2kCgYEAwSCKaTp+HSBCvX5f3Log
QkZp5k735RwMMIe2xtaRUGQqYs1sknO7d/CjK+htayk9bP1SZHnpowXHFGSW/nuB
p/tBDv/24OPQNdgo089+T2ne7MCl3SC/9wa2neeb/JT49jexZjq6mbxucPeujzE1
2kF3jxhtYgXEUFq1JSuNvnUCgYEAyUO4Y7na4xfjzxjfkloE/A8Z7TztaTniu54l
06IwzLlb8cTPUtxeD8/Ks1sC8kTr6bUFBXeZizBdj0bH+sF59DImCDoi9vTattS+
zRJmKyx5duScve4J+zq5Hm5kR8tgHR0S17bynwUbnGdOy8JmG22sflq5tTVBhaeG
UW/rOnsCgYASSD5SD9N4dmFbBueUQZpkK75Cqx8UdT9CKNbIo+9FqPXKPKAWjRYm
GIWZ1nrlNhY2hxSRpmjToexipdMVbCOt/z79aIW6bFZ9gmT7CB1w7xjHWMVa1YrW
m7AV6qL9miynQkZs4wpfG1NpJklEDOiILMJgrXNNYDZhVPTo++KDMQKBgDw4J/6m
yGh0aHQ5tANdLeqhNhe2yC5Y5I9QhW7qM4G94FXZllLrnrVKbhL2I06L8q5tvD/j
hiyQXx4Uhpdvtmarbpe9lWKg5qQXybMgUzONzhYV1xQ5GgFyk5sYWqbkojBz14R1
t+h+pcFJY9kxpE2Gpjr0OGaQtbcg5d6OByrrAoGAZAkRVxGdy13rPx4bV55ZaR46
2Si7wb8SIFbngO8Chnt/sPhkExBMel81yIHTuMhO3sOBr/X60iqOK+zUiuo4XShA
admg6OChpyC0h+uUCDOjStYfx7eWMpCqTazUwYrwBEB3zy6LTc288kXg3jT1HJpc
RK1faU5CS8ah4qZ33Qk=
-----END PRIVATE KEY-----`
}

func getPubKey() string {
	return `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAl9WjS3e+IaoA8XrVfQyz
nczQv0bpKoN84W7NUnBZl30VG8WFix79Ir+msazoh18uivzIUHPcY0+jHMQc/Qjo
+S/5FOC9LBURCtla3QU3ty3zn0i+cKW6lk6c4SbOAP3zHVsf3l0rGswBxND9x4QQ
ICiBx1IxS6NlSeSDlwBerkJCAblEF6pTGBwZDYL/eJ0eTPH9U4ismH5OYs4wsnRl
BGJUKj7YZR4CvkTErygSNxataaLuP1Qqp8Spr2cksnOfC1/tYy/N0lydpQN9UfL6
NqpKiGg6Rsp4k1skiKR535Eue8YJb0RqQ7kjr9H4F18cIwfI5iz9vwSQfmZxmqgE
NwIDAQAB
-----END PUBLIC KEY-----`
}
