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

// 通过公钥验证签名通过与否
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

// 公钥加密
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

// 私钥加密
func (r *ChsiRsa) PrivateEncrypt(data string) (string, error) {
	signData, err := rsa.SignPKCS1v15(nil, r.privateKey, crypto.Hash(0), []byte(data))
	if err != nil {
		return "", err
	}
	rt := base64.StdEncoding.EncodeToString(signData)
	return rt, nil
}

// 公钥解密
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
