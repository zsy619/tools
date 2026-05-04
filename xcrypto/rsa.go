package xcrypto

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io"
)

// https://cloud.tencent.com/developer/article/1475706

const (
	CHAR_SET               = "UTF-8"
	BASE_64_FORMAT         = "UrlSafeNoPadding"
	RSA_ALGORITHM_KEY_TYPE = "PKCS8"
	RSA_ALGORITHM_SIGN     = crypto.SHA256
)

type XRsa struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// GenRsaKey 函数生成RSA密钥对并返回
//
// 返回值：
//   - []byte: 私钥的字节切片
//   - []byte: 公钥的字节切片
//   - error: 若有错误则返回错误信息，否则为nil
func GenRsaKey() ([]byte, []byte, error) {
	publicKey := bytes.NewBufferString("")
	privateKey := bytes.NewBufferString("")

	err := CreateKeys(publicKey, privateKey, 2048)
	if err != nil {
		return nil, nil, err
	}
	return privateKey.Bytes(), publicKey.Bytes(), nil
}

// 生成密钥对
// CreateKeys 函数生成RSA密钥对，并将私钥和公钥分别写入指定的io.Writer中。
// publicKeyWriter是公钥写入的目标io.Writer，privateKeyWriter是私钥写入的目标io.Writer。
// keyLength是密钥的长度（以位为单位）。
// 如果在生成密钥对或写入文件的过程中发生错误，将返回非零错误码。
func CreateKeys(publicKeyWriter, privateKeyWriter io.Writer, keyLength int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		return err
	}
	derStream := MarshalPKCS8PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	err = pem.Encode(privateKeyWriter, block)
	if err != nil {
		return err
	}

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	err = pem.Encode(publicKeyWriter, block)
	if err != nil {
		return err
	}

	return nil
}

// NewXRsa 是一个构造函数，用于根据给定的公钥和私钥字节数组创建XRsa实例
//
// publicKey: 公钥的字节数组
// privateKey: 私钥的字节数组
//
// 返回值：
// *XRsa：包含公钥和私钥的XRsa实例指针
// error：如果发生错误，则返回非零的错误码
func NewXRsa(publicKey []byte, privateKey []byte) (*XRsa, error) {
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
		return &XRsa{
			publicKey:  pub,
			privateKey: pri,
		}, nil
	} else {
		return nil, errors.New("private key not supported")
	}
}

// EncryptWithPublicKey 使用RSA公钥对数据进行加密
//
// 参数：
//
//	r *XRsa - 包含RSA公钥的XRsa对象指针
//	data string - 待加密的字符串数据
//
// 返回值：
//
//	string - 加密后的Base64编码字符串
//	error - 如果加密过程中发生错误，则返回非nil的错误信息
func (r *XRsa) EncryptWithPublicKey(data string) (string, error) {
	partLen := r.publicKey.N.BitLen()/8 - 11
	chunks := split([]byte(data), partLen)

	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		bytes, err := rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(bytes)
	}

	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

// EncryptWithPrivateKey 使用私钥对给定的字符串进行加密（实际上是签名），并返回加密后的Base64字符串和可能发生的错误。
// 参数：
//
//	data string - 待加密的字符串
//
// 返回值：
//
//	string - 加密后的Base64字符串
//	error - 如果加密过程中发生错误，则返回非零的错误码；否则返回nil
func (r *XRsa) EncryptWithPrivateKey(data string) (string, error) {
	output := bytes.NewBuffer(nil)
	err := priKeyIO(r.privateKey, bytes.NewReader([]byte(data)), output, true)
	if err != nil {
		return "", err
	}
	out, err := io.ReadAll(output)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(out), nil
}

// 私钥解密
// DecryptWithPrivateKey 使用私钥对加密后的字符串进行解密
//
// 参数：
//
//	encrypted string - 待解密的加密字符串（应为base64编码）
//
// 返回值：
//
//	string - 解密后的明文字符串
//	error - 解密过程中可能遇到的错误，若解密成功则为nil
func (r *XRsa) DecryptWithPrivateKey(encrypted string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	partLen := r.publicKey.N.BitLen() / 8
	chunks := split([]byte(raw), partLen)

	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, r.privateKey, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(decrypted)
	}

	return buffer.String(), err
}

// DecryptWithPublicKey 使用RSA公钥解密方法解密字符串
// 参数：
//
//	encrypted: 待解密的字符串（Base64编码）
//
// 返回值：
//
//	string: 解密后的字符串
//	error: 如果解密过程中发生错误，则返回非零的错误码
func (r *XRsa) DecryptWithPublicKey(encrypted string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	output := bytes.NewBuffer(nil)
	err = pubKeyIO(r.publicKey, bytes.NewReader(raw), output, false)
	if err != nil {
		return "", err
	}
	outStr, err := io.ReadAll(output)
	if err != nil {
		return "", err
	}
	return string(outStr), nil
}

// 数据加签
func (r *XRsa) Sign(data string) (string, error) {
	h := RSA_ALGORITHM_SIGN.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	sign, err := rsa.SignPKCS1v15(rand.Reader, r.privateKey, RSA_ALGORITHM_SIGN, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), err
}

// 数据验签
func (r *XRsa) Verify(data string, sign string) error {
	h := RSA_ALGORITHM_SIGN.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	decodedSign, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(r.publicKey, RSA_ALGORITHM_SIGN, hashed, decodedSign)
}

func MarshalPKCS8PrivateKey(key *rsa.PrivateKey) []byte {
	info := struct {
		Version             int
		PrivateKeyAlgorithm []asn1.ObjectIdentifier
		PrivateKey          []byte
	}{}
	info.Version = 0
	info.PrivateKeyAlgorithm = make([]asn1.ObjectIdentifier, 1)
	info.PrivateKeyAlgorithm[0] = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
	info.PrivateKey = x509.MarshalPKCS1PrivateKey(key)

	k, _ := asn1.Marshal(info)
	return k
}

func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}
