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
// Sm2GenerateKey 生成一对新的 SM2 公私钥。
//
// 返回值：
//   - *sm2.PrivateKey: 生成的 SM2 私钥指针。
//   - *sm2.PublicKey: 从私钥派生的公钥指针。
//   - error: 生成失败时返回错误（如底层熵源不足）。
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
// Sm2PriKeyToPubKey 由 SM2 私钥派生出对应的公钥（基点乘法）。
//
// 参数：
//   - priKey: SM2 私钥指针。
//
// 返回值：派生的 SM2 公钥指针。
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
// Sm2Encrypt 使用 SM2 公钥加密 data，并返回 Base64 编码的密文。
//
// 参数：
//   - pubKey: SM2 公钥。
//   - data: 待加密的明文字符串。
//   - mode: 密文顺序模式（C1C2C3 或 C1C3C2）。
//
// 返回值：
//   - string: Base64 编码的密文。
//   - error: 加密失败时返回错误。
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
// Sm2Decrypt 使用 SM2 私钥解密 Base64 编码的密文 data。
//
// 参数：
//   - priKey: SM2 私钥。
//   - data: Base64 编码的密文。
//   - mode: 密文顺序模式，需要与加密时一致。
//
// 返回值：
//   - string: 解密后的明文字符串。
//   - error: Base64 解码失败或解密失败时返回错误。
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
// Sm2Sign 使用 SM2 私钥对 data 进行签名。
//
// 参数：
//   - priKey: SM2 私钥。
//   - data: 待签名的原始字符串。
//
// 返回值：
//   - string: 签名结果（字节字符串形式）。
//   - *sm2.PublicKey: 用于验签的公钥。
//   - error: 签名失败时返回错误。
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
// Sm2Verify 使用 SM2 公钥校验 data 与签名 sign 是否一致。
//
// 参数：
//   - pubKey: SM2 公钥。
//   - data: 原始数据。
//   - sign: 签名结果。
//
// 返回值：验签通过返回 true，否则返回 false。
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
// Sm2PubKeyToBase64 将 SM2 公钥转换为带 X.509 算法 OID 头部的标准 Base64 字符串。
//
// 参数：
//   - pubKey: SM2 公钥指针。
//
// 返回值：编码后的 Base64 字符串。
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
// Sm2Base64ToPubKey 将 Sm2PubKeyToBase64 编码的 Base64 字符串还原为 SM2 公钥指针。
//
// 参数：
//   - pubStr: 公钥的 Base64 字符串。
//
// 返回值：对应的 SM2 公钥指针；解析失败时其内部字段为零值。
func Sm2Base64ToPubKey(pubStr string) *sm2.PublicKey {
	decode, _ := base64.StdEncoding.DecodeString(pubStr)
	pubHex := hex.EncodeToString(decode)
	pubHex = strings.ReplaceAll(pubHex, "3059301306072a8648ce3d020106082a811ccf5501822d03420004", "")
	pubByte, _ := hex.DecodeString(pubHex)
	pub, _ := RawBytesToPubKey(pubByte)
	return pub
}

// Sm2PriKeyToBase64 返回 SM2 私钥的十六进制字符串（等价于 Sm2PriKeyToHex）。
func Sm2PriKeyToBase64(priKey *sm2.PrivateKey) string {
	hex := hex.EncodeToString(GetPriKeyRawBytes(priKey))
	return hex
}

/**
 * @description: 私钥转hex
 * @param {*sm2.PrivateKey} priKey
 * @return {*}
 */
// Sm2PriKeyToHex 将 SM2 私钥转换为十六进制字符串。
//
// 参数：
//   - priKey: SM2 私钥指针。
//
// 返回值：固定长度（KeyBytes*2）字符的十六进制字符串。
func Sm2PriKeyToHex(priKey *sm2.PrivateKey) string {
	hex := hex.EncodeToString(GetPriKeyRawBytes(priKey))
	return hex
}

/**
 * @description: Hex私钥转私钥对象
 * @param {string} priStr
 * @return {*}
 */
// Sm2HexToPriKey 将 Sm2PriKeyToHex 输出的十六进制字符串解析为 SM2 私钥指针。
//
// 参数：
//   - priStr: SM2 私钥的十六进制字符串。
//
// 返回值：对应的 SM2 私钥指针。
//
// panic 条件：当私钥长度不合法或解析失败时 panic。
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
