package xcrypto

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/tjfoc/gmsm/sm2"
)

func TestSm2(t *testing.T) {
	// 生成私钥
	privateKey, e := sm2.GenerateKey(rand.Reader)
	if e != nil {
		fmt.Println("sm2 encrypt faild！")
	}
	// 从私钥中获取公钥
	pubkey := &privateKey.PublicKey

	msg := []byte("i am   wek &&  i am The_Reader too 。")
	// 用公钥加密msg
	bytes, i := pubkey.EncryptAsn1(msg, rand.Reader)

	if i != nil {
		fmt.Println("使用私钥加密失败！")
	}

	fmt.Println("the encrypt msg  =  ", hex.EncodeToString(bytes))
	// 用私钥解密msg
	decrypt, i2 := privateKey.DecryptAsn1(bytes)

	if i2 != nil {
		fmt.Println("使用私钥解密失败！")
	}

	fmt.Println("the msg  = ", string(decrypt))
}

func TestSm2GenerateKey(t *testing.T) {
	priKey, pubKey, err := Sm2GenerateKey()
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(Sm2PriKeyToHex(priKey))
		t.Log(Sm2PubKeyToBase64(pubKey))

		pubTest := Sm2PriKeyToPubKey(priKey)
		t.Log(Sm2PubKeyToBase64(pubTest))
	}
}
