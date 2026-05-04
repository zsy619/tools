package xcrypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
)

type ChsiRsa struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// NewChsiRSA 根据提供的公钥和私钥字节数组创建一个ChsiRsa结构体实例
// publicKey 表示公钥的字节数组
// privateKey 表示私钥的字节数组
// 返回值：
// *ChsiRsa 返回一个ChsiRsa结构体指针，包含解析后的公钥和私钥
// error 如果解析过程中发生错误，则返回非空错误
func NewChsiRSA(publicKey []byte, privateKey []byte) (*ChsiRsa, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)

	block, _ = pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pri, ok := priv.(*rsa.PrivateKey)
	if ok {
		return &ChsiRsa{
			publicKey:  pub,
			privateKey: pri,
		}, nil
	} else {
		return nil, errors.New("private key not supported")
	}
}

// SetPrvKey 方法为 ChsiRsa 结构体设置私钥
//
// 参数：
// pkey: []byte - 私钥的字节切片
//
// 返回值：
// 无
func (chsirsa *ChsiRsa) SetPrvKey(pkey []byte) {
	block, _ := pem.Decode(pkey)
	if block == nil {
		fmt.Println("pem.Decode err")
		return
	}

	private, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		return
	}

	chsirsa.privateKey = private.(*rsa.PrivateKey)
}

// SetPubKey 方法用于为ChsiRsa结构体设置公钥
// 参数pkey是公钥的字节切片
func (chsirsa *ChsiRsa) SetPubKey(pkey []byte) {
	block, _ := pem.Decode(pkey)
	if block == nil {
		return
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}

	chsirsa.publicKey = pubInterface.(*rsa.PublicKey)
}

// Sign 使用给定的私钥对字符串数据进行签名
//
// 参数：
//
//	chsirsa *ChsiRsa: 指向ChsiRsa类型的指针，包含私钥信息
//	data string: 待签名的字符串数据
//
// 返回值：
//
//	string: 签名后的字符串，使用base64编码
//	error: 如果签名过程中出现错误，则返回非零的错误码；否则返回nil
func (chsirsa *ChsiRsa) Sign(data string) (string, error) {
	h := sha1.New()
	h.Write([]byte(data))
	digest := h.Sum(nil)

	s, err := rsa.SignPKCS1v15(nil, chsirsa.privateKey, crypto.SHA1, digest)
	if err != nil {
		return "", err
	}
	signRet := base64.StdEncoding.EncodeToString(s)
	return signRet, nil
}

// Verify 使用ChsiRsa公钥验证签名的合法性
//
// 参数：
//
//	chsirsa *ChsiRsa：ChsiRsa公钥对象
//	content string：待验证的原文内容
//	sign string：签名值（base64编码）
//
// 返回值：
//
//	bool：验证结果，true表示验证通过，false表示验证不通过
//	error：如果验证过程中发生错误，则返回非零的错误码；否则返回nil
func (chsirsa *ChsiRsa) Verify(content, sign string) (bool, error) {
	h := RSA_ALGORITHM_SIGN.New()
	h.Write([]byte(content))
	hashed := h.Sum(nil)

	decodedSign, err := base64.RawURLEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}

	return false, rsa.VerifyPKCS1v15(chsirsa.publicKey, RSA_ALGORITHM_SIGN, hashed, decodedSign)
}

// PublicEncrypt 是ChsiRsa结构体上的方法，用于对传入的明文字符串进行RSA公钥加密
// 参数:
//
//	r: ChsiRsa结构体指针，包含RSA公钥信息
//	data: 待加密的明文字符串
//
// 返回值:
//
//	string: 加密后的密文字符串（Base64编码）
//	error: 如果加密过程中发生错误，则返回非零的错误码；否则返回nil
func (r *ChsiRsa) PublicEncrypt(data string) (string, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, []byte(data))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	rt := base64.StdEncoding.EncodeToString(ciphertext)
	return rt, nil
}

// 私钥解密
func (r *ChsiRsa) PrivateDecrypt(data string) (string, error) {
	edata, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	ciphertext, err := rsa.DecryptPKCS1v15(rand.Reader, r.privateKey, edata)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(ciphertext), nil
}

// PrivateEncrypt 使用 RSA 私钥对数据进行加密（实际为签名）
//
// 参数：
//
//	data string - 待加密（签名）的字符串
//
// 返回值：
//
//	string - 加密（签名）后的字符串，以 Base64 编码表示
//	error - 如果加密（签名）过程中发生错误，则返回非零的错误码
//
// 注意：
//
//	此函数实际上执行的是 RSA 签名操作，而不是传统意义上的加密，因为 RSA 私钥通常用于签名，公钥用于验证签名。
//	在这里，加密一词可能被误用，更准确的说法应该是签名。
func (r *ChsiRsa) PrivateEncrypt(data string) (string, error) {
	signData, err := rsa.SignPKCS1v15(nil, r.privateKey, crypto.Hash(0), []byte(data))
	if err != nil {
		return "", err
	}
	rt := base64.StdEncoding.EncodeToString(signData)
	return rt, nil
}

// PublicDecrypt 使用公钥解密传入的base64编码后的密文数据
// 参数：
//
//	data：待解密的base64编码后的密文数据
//
// 返回值：
//
//	string：解密后的明文数据
//	error：解密过程中遇到的错误信息，如果解密成功则为nil
func (r *ChsiRsa) PublicDecrypt(data string) (string, error) {
	edata, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	decData, err := r.publicDecrypt(r.publicKey, crypto.Hash(0), nil, edata)
	if err != nil {
		return "", err
	}
	return string(decData), nil
}

// copy&modified from crypt/rsa/pkcs1v5.go
// publicDecrypt 使用给定的公钥对签名进行公钥解密操作
//
// 参数：
// pub: *rsa.PublicKey - 公钥对象
// hash: crypto.Hash - 散列函数类型
// hashed: []byte - 待验证的原始数据散列值
// sig: []byte - 待解密的签名
//
// 返回值：
// out: []byte - 解密后的数据
// err: error - 错误信息，如果操作成功则为nil
func (r *ChsiRsa) publicDecrypt(pub *rsa.PublicKey, hash crypto.Hash, hashed []byte, sig []byte) (out []byte, err error) {
	hashLen, prefix, err := r.pkcs1v15HashInfo(hash, len(hashed))
	if err != nil {
		return nil, err
	}

	tLen := len(prefix) + hashLen
	k := (pub.N.BitLen() + 7) / 8
	if k < tLen+11 {
		return nil, fmt.Errorf("length illegal")
	}

	c := new(big.Int).SetBytes(sig)
	m := r.encrypt(new(big.Int), pub, c)
	em := r.leftPad(m.Bytes(), k)
	out = r.unLeftPad(em)

	err = nil
	return
}

// copy from crypt/rsa/pkcs1v5.go
var hashPrefixes = map[crypto.Hash][]byte{
	crypto.MD5:       {0x30, 0x20, 0x30, 0x0c, 0x06, 0x08, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x02, 0x05, 0x05, 0x00, 0x04, 0x10},
	crypto.SHA1:      {0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14},
	crypto.SHA224:    {0x30, 0x2d, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x04, 0x05, 0x00, 0x04, 0x1c},
	crypto.SHA256:    {0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20},
	crypto.SHA384:    {0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30},
	crypto.SHA512:    {0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40},
	crypto.MD5SHA1:   {}, // A special TLS case which doesn't use an ASN1 prefix.
	crypto.RIPEMD160: {0x30, 0x20, 0x30, 0x08, 0x06, 0x06, 0x28, 0xcf, 0x06, 0x03, 0x00, 0x31, 0x04, 0x14},
}

// copy from crypt/rsa/pkcs1v5.go
func (r *ChsiRsa) encrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) *big.Int {
	e := big.NewInt(int64(pub.E))
	c.Exp(m, e, pub.N)
	return c
}

// copy from crypt/rsa/pkcs1v5.go
func (r *ChsiRsa) pkcs1v15HashInfo(hash crypto.Hash, inLen int) (hashLen int, prefix []byte, err error) {
	// Special case: crypto.Hash(0) is used to indicate that the data is
	// signed directly.
	if hash == 0 {
		return inLen, nil, nil
	}

	hashLen = hash.Size()
	if inLen != hashLen {
		return 0, nil, errors.New("crypto/rsa: input must be hashed message")
	}
	prefix, ok := hashPrefixes[hash]
	if !ok {
		return 0, nil, errors.New("crypto/rsa: unsupported hash function")
	}
	return
}

// copy from crypt/rsa/pkcs1v5.go
func (r *ChsiRsa) leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}

func (r *ChsiRsa) unLeftPad(input []byte) (out []byte) {
	n := len(input)
	t := 2
	for i := 2; i < n; i++ {
		if input[i] == 0xff {
			t = t + 1
		} else {
			if input[i] == input[0] {
				t = t + int(input[1])
			}
			break
		}
	}
	out = make([]byte, n-t)
	copy(out, input[t:])
	return
}
