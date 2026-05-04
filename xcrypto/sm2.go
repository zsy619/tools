package xcrypto

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/tjfoc/gmsm/sm2"
)

// https://blog.csdn.net/weixin_42117918/article/details/130558002?spm=1001.2101.3001.6650.2&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-2-130558002-blog-83998138.235%5Ev38%5Epc_relevant_anti_vip&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-2-130558002-blog-83998138.235%5Ev38%5Epc_relevant_anti_vip&utm_relevant_index=2

/**
 * @description: 生成秘钥对
 * @return {*}
 */
func Sm2GenerateKey() (*sm2.PrivateKey, *sm2.PublicKey, error) {
	priKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("秘钥产生失败：", err)
		return nil, nil, err
	}
	pubKey := &priKey.PublicKey
	return priKey, pubKey, nil
}

/**
 * @description: 私钥生成公钥
 * @param {*sm2.PrivateKey} priKey
 * @return {*}
 */
func Sm2PriKeyToPubKey(priKey *sm2.PrivateKey) *sm2.PublicKey {
	pub := new(sm2.PublicKey)
	pub.Curve = priKey.Curve
	pub.X, pub.Y = priKey.Curve.ScalarBaseMult(priKey.D.Bytes())
	return pub
}

/**
 * @description: 通过公钥加密
 * @param {*sm2.PublicKey} pubKey
 * @param {[]byte} data
 * @return {*}
 */
func Sm2Encrypt(pubKey *sm2.PublicKey, data string, mode int) (string, error) {
	ciphertxt, err := sm2.Encrypt(pubKey, []byte(data), rand.Reader, mode)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertxt), nil
}

/**
 * @description: 通过私钥解密
 * @param {*sm2.PrivateKey} priKey
 * @param {string} data
 * @return {*}
 */
func Sm2Decrypt(priKey *sm2.PrivateKey, data string, mode int) (string, error) {
	ciphertxt, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	plaintxt, err := sm2.Decrypt(priKey, []byte(ciphertxt), mode)
	if err != nil {
		return "", err
	}
	return string(plaintxt), nil
}

/**
 * @description: 通过私钥签名
 * @param {*sm2.PrivateKey} priKey
 * @param {string} data
 * @return {*}
 */
func Sm2Sign(priKey *sm2.PrivateKey, data string) (string, *sm2.PublicKey, error) {
	curve := sm2.P256Sm2() // 椭圆曲线
	priv := new(sm2.PrivateKey)
	priv.PublicKey.Curve = curve
	priv.D = priKey.D
	priv.PublicKey.X = priKey.X
	priv.PublicKey.Y = priKey.Y

	sign, err := priv.Sign(rand.Reader, []byte(data), nil) // sm2签名
	if err != nil {
		return "", nil, err
	}
	return string(sign), &priv.PublicKey, nil
}

/**
 * @description: 通过公钥验签
 * @param {*sm2.PublicKey} pubKey
 * @param {*} data
 * @param {string} sign
 * @return {*}
 */
func Sm2Verify(pubKey *sm2.PublicKey, data, sign string) bool {
	curve := sm2.P256Sm2() // 椭圆曲线
	pub := new(sm2.PublicKey)
	pub.Curve = curve
	pub.X = pubKey.X
	pub.Y = pubKey.Y

	return pub.Verify([]byte(data), []byte(sign))
}

/**
 * @description: 公钥转base64
 * @param {*sm2.PublicKey} pubKey
 * @return {*}
 */
func Sm2PubKeyToBase64(pubKey *sm2.PublicKey) string {
	// pubKey.GetRawBytes()
	bytes := pubKey.X.Bytes()
	bytes = append(bytes, pubKey.Y.Bytes()...)
	pubHex := "3059301306072a8648ce3d020106082a811ccf5501822d03420004" + hex.EncodeToString(bytes)
	decode, _ := hex.DecodeString(pubHex)
	base64Pub := base64.StdEncoding.EncodeToString(decode)
	return base64Pub
}

/**
 * @description: base64公钥转公钥对象
 * @param {string} pubStr
 * @return {*}
 */
func Sm2Base64ToPubKey(pubStr string) *sm2.PublicKey {
	decode, _ := base64.StdEncoding.DecodeString(pubStr)
	pubHex := hex.EncodeToString(decode)
	pubHex = strings.ReplaceAll(pubHex, "3059301306072a8648ce3d020106082a811ccf5501822d03420004", "")
	pubByte, _ := hex.DecodeString(pubHex)
	pub, _ := RawBytesToPubKey(pubByte)
	return pub
}

func Sm2PriKeyToBase64(priKey *sm2.PrivateKey) string {
	hex := hex.EncodeToString(GetPriKeyRawBytes(priKey))
	return hex
}

/**
 * @description: 私钥转hex
 * @param {*sm2.PrivateKey} priKey
 * @return {*}
 */
func Sm2PriKeyToHex(priKey *sm2.PrivateKey) string {
	hex := hex.EncodeToString(GetPriKeyRawBytes(priKey))
	return hex
}

/**
 * @description: Hex私钥转私钥对象
 * @param {string} priStr
 * @return {*}
 */
func Sm2HexToPriKey(priStr string) *sm2.PrivateKey {
	// 解码hex私钥
	privateKeyByte, _ := hex.DecodeString(priStr)
	// 转成go版的私钥
	pri, err := RawBytesToPriKey(privateKeyByte)
	if err != nil {
		panic("私钥加载异常")
	}
	return pri
}
