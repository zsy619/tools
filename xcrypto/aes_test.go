package xcrypto

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"testing"
)

func Test_Aes(t *testing.T) {
	origData := []byte(`{"year":2021,"ksh":"ksh","yxdm":"yxdm","yxmc":"ceshi","sfzh":"412924197602102558","xm":"朱书彦","xbdm":"1","mzdm":"01","zzmmdm":"02","xldm":"03","xxxs":"999","zydm":"90","pyfsdm":"1","syszddm":"8889","xz":"2.5","rxsj":"20200809","bysj":"20200809","sfslbdm":"09","knslbdm":"04","mobilePhone":"13633861512"}`) // 待加密的数据，注意加密数据里面的yxdm与参数的yxdm保持一致
	// origData := []byte(`{"year":2021,"ksh":"ksh","byqxdm":"02"}`) // 待加密的数据	毕业生去向
	// origData := []byte(`{"year":2019,"ksh":"1071016141305506"}`) // 待加密的数据
	key := []byte("911dd5a382534bccd38a80e51489fd93") // AES32位加密的密钥
	log.Println("原文：", string(origData))

	log.Println("------------------ CBC模式 --------------------")
	encrypted := AesEncryptCBC(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted, _ := AesDecryptCBC(encrypted, key)
	log.Println("解密结果：", string(decrypted))

	log.Println("------------------ ECB模式 --------------------")
	encrypted = AesEncryptECB(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AesDecryptECB(encrypted, key)
	log.Println("解密结果：", string(decrypted))

	log.Println("------------------ CFB模式 --------------------")
	encrypted = AesEncryptCFB(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AesDecryptCFB(encrypted, key)
	log.Println("解密结果：", string(decrypted))
}
