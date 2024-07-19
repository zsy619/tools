package xcrypto

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"io"
	"math/big"
)

var (
	ErrDataToLarge     = errors.New("message too long for RSA public key size")
	ErrDataLen         = errors.New("data length error")
	ErrDataBroken      = errors.New("data broken, first byte is not zero")
	ErrKeyPairDismatch = errors.New("data is not encrypted by the private key")
	ErrDecryption      = errors.New("decryption error")
	ErrPublicKey       = errors.New("get public key error")
	ErrPrivateKey      = errors.New("get private key error")
)

// pubKeyByte 根据RSA公钥对数据进行加密或解密操作
//
// pub: RSA公钥指针
// in: 待加密或解密的数据
// isEncrytp: 是否为加密操作，true表示加密，false表示解密
//
// 返回值：
// []byte: 加密或解密后的数据
// error: 如果发生错误，则返回非零的错误码
func pubKeyByte(pub *rsa.PublicKey, in []byte, isEncrytp bool) ([]byte, error) {
	k := (pub.N.BitLen() + 7) / 8
	if isEncrytp {
		k = k - 11
	}
	if len(in) <= k {
		if isEncrytp {
			return rsa.EncryptPKCS1v15(rand.Reader, pub, in)
		} else {
			return pubKeyDecrypt(pub, in)
		}
	} else {
		iv := make([]byte, k)
		out := bytes.NewBuffer(iv)
		if err := pubKeyIO(pub, bytes.NewReader(in), out, isEncrytp); err != nil {
			return nil, err
		}
		return io.ReadAll(out)
	}
}

// priKeyByte 是一个用于处理RSA私钥加密或解密数据的函数
//
// pri 是指向RSA私钥的指针
// in 是待加密或解密的数据切片
// isEncrytp 是一个布尔值，表示是否进行加密操作，如果为true则进行加密，否则进行解密
//
// 返回值：
// - []byte：加密或解密后的数据切片
// - error：如果在处理过程中发生错误，则返回非零的错误码
func priKeyByte(pri *rsa.PrivateKey, in []byte, isEncrytp bool) ([]byte, error) {
	k := (pri.N.BitLen() + 7) / 8
	if isEncrytp {
		k = k - 11
	}
	if len(in) <= k {
		if isEncrytp {
			return priKeyEncrypt(rand.Reader, pri, in)
		} else {
			return rsa.DecryptPKCS1v15(rand.Reader, pri, in)
		}
	} else {
		iv := make([]byte, k)
		out := bytes.NewBuffer(iv)
		if err := priKeyIO(pri, bytes.NewReader(in), out, isEncrytp); err != nil {
			return nil, err
		}
		return io.ReadAll(out)
	}
}

// pubKeyIO 函数用于对输入流进行公钥加密或解密，并将结果写入输出流。
//
// pub 是待使用的RSA公钥。
// in 是输入流，需要从中读取待加密或解密的数据。
// out 是输出流，用于写入加密或解密后的数据。
// isEncrytp 是一个布尔值，指示是否进行加密操作（true为加密，false为解密）。
//
// 如果加密，则使用RSA PKCS1v15填充方式进行加密；如果解密，则调用pubKeyDecrypt函数进行解密。
//
// 函数返回可能发生的错误。
func pubKeyIO(pub *rsa.PublicKey, in io.Reader, out io.Writer, isEncrytp bool) (err error) {
	k := (pub.N.BitLen() + 7) / 8
	if isEncrytp {
		k = k - 11
	}
	buf := make([]byte, k)
	var b []byte
	size := 0
	for {
		size, err = in.Read(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if size < k {
			b = buf[:size]
		} else {
			b = buf
		}
		if isEncrytp {
			b, err = rsa.EncryptPKCS1v15(rand.Reader, pub, b)
		} else {
			b, err = pubKeyDecrypt(pub, b)
		}
		if err != nil {
			return err
		}
		if _, err = out.Write(b); err != nil {
			return err
		}
	}
}

// priKeyIO 是一个从io.Reader读取数据，通过RSA私钥进行加密或解密，然后将结果写入到io.Writer的函数
//
// 参数：
//
//	pri: *rsa.PrivateKey - RSA私钥
//	r: io.Reader - 数据读取的源
//	w: io.Writer - 数据写入的目标
//	isEncrytp: bool - 是否进行加密操作，true表示加密，false表示解密
//
// 返回值：
//
//	err: error - 如果操作过程中出现错误，则返回非nil的错误信息，否则返回nil
func priKeyIO(pri *rsa.PrivateKey, r io.Reader, w io.Writer, isEncrytp bool) (err error) {
	k := (pri.N.BitLen() + 7) / 8
	if isEncrytp {
		k = k - 11
	}
	buf := make([]byte, k)
	var b []byte
	size := 0
	for {
		size, err = r.Read(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if size < k {
			b = buf[:size]
		} else {
			b = buf
		}
		if isEncrytp {
			b, err = priKeyEncrypt(rand.Reader, pri, b)
		} else {
			b, err = rsa.DecryptPKCS1v15(rand.Reader, pri, b)
		}
		if err != nil {
			return err
		}
		if _, err = w.Write(b); err != nil {
			return err
		}
	}
}

// pubKeyDecrypt 使用RSA公钥对密文进行解密
//
// 参数：
//
//	pub *rsa.PublicKey：RSA公钥
//	data []byte：待解密的密文数据
//
// 返回值：
//
//	[]byte：解密后的明文数据
//	error：解密过程中出现的错误，如果解密成功则为nil
func pubKeyDecrypt(pub *rsa.PublicKey, data []byte) ([]byte, error) {
	k := (pub.N.BitLen() + 7) / 8
	if k != len(data) {
		return nil, ErrDataLen
	}
	m := new(big.Int).SetBytes(data)
	if m.Cmp(pub.N) > 0 {
		return nil, ErrDataToLarge
	}
	m.Exp(m, big.NewInt(int64(pub.E)), pub.N)
	d := leftPad(m.Bytes(), k)
	if d[0] != 0 {
		return nil, ErrDataBroken
	}
	if d[1] != 0 && d[1] != 1 {
		return nil, ErrKeyPairDismatch
	}
	i := 2
	for ; i < len(d); i++ {
		if d[i] == 0 {
			break
		}
	}
	i++
	if i == len(d) {
		return nil, nil
	}
	return d[i:], nil
}

// 私钥加密
func priKeyEncrypt(rand io.Reader, priv *rsa.PrivateKey, hashed []byte) ([]byte, error) {
	tLen := len(hashed)
	k := (priv.N.BitLen() + 7) / 8
	if k < tLen+11 {
		return nil, ErrDataLen
	}
	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < k-tLen-1; i++ {
		em[i] = 0xff
	}
	copy(em[k-tLen:k], hashed)
	m := new(big.Int).SetBytes(em)
	c, err := rsa_decrypt(rand, priv, m)
	if err != nil {
		return nil, err
	}
	copyWithLeftPad(em, c.Bytes())
	return em, nil
}

// 从crypto/rsa复制
var (
	bigZero = big.NewInt(0)
	bigOne  = big.NewInt(1)
)

// rsa_encrypt 使用公钥pub对消息m进行RSA加密，并将结果存储在c中
// 参数：
//
//	c *big.Int: 存储加密结果的变量，不应为nil
//	pub *rsa.PublicKey: 用于加密的公钥
//	m *big.Int: 待加密的消息
//
// 返回值：
//
//	*big.Int: 加密后的结果，存储在c中并返回
func rsa_encrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) *big.Int {
	e := big.NewInt(int64(pub.E))
	c.Exp(m, e, pub.N)
	return c
}

// rsa_decrypt 使用RSA私钥解密密文
//
// 参数：
//
//	random: io.Reader 用于生成随机数的读取器，可以为nil
//	priv: *rsa.PrivateKey RSA私钥
//	c: *big.Int 待解密的密文
//
// 返回值：
//
//	m: *big.Int 解密后的明文
//	err: error 错误信息，如果解密成功则为nil
func rsa_decrypt(random io.Reader, priv *rsa.PrivateKey, c *big.Int) (m *big.Int, err error) {
	if c.Cmp(priv.N) > 0 {
		err = ErrDecryption
		return
	}
	var ir *big.Int
	if random != nil {
		var r *big.Int

		for {
			r, err = rand.Int(random, priv.N)
			if err != nil {
				return
			}
			if r.Cmp(bigZero) == 0 {
				r = bigOne
			}
			var ok bool
			ir, ok = modInverse(r, priv.N)
			if ok {
				break
			}
		}
		bigE := big.NewInt(int64(priv.E))
		rpowe := new(big.Int).Exp(r, bigE, priv.N)
		cCopy := new(big.Int).Set(c)
		cCopy.Mul(cCopy, rpowe)
		cCopy.Mod(cCopy, priv.N)
		c = cCopy
	}
	if priv.Precomputed.Dp == nil {
		m = new(big.Int).Exp(c, priv.D, priv.N)
	} else {
		m = new(big.Int).Exp(c, priv.Precomputed.Dp, priv.Primes[0])
		m2 := new(big.Int).Exp(c, priv.Precomputed.Dq, priv.Primes[1])
		m.Sub(m, m2)
		if m.Sign() < 0 {
			m.Add(m, priv.Primes[0])
		}
		m.Mul(m, priv.Precomputed.Qinv)
		m.Mod(m, priv.Primes[0])
		m.Mul(m, priv.Primes[1])
		m.Add(m, m2)

		for i, values := range priv.Precomputed.CRTValues {
			prime := priv.Primes[2+i]
			m2.Exp(c, values.Exp, prime)
			m2.Sub(m2, m)
			m2.Mul(m2, values.Coeff)
			m2.Mod(m2, prime)
			if m2.Sign() < 0 {
				m2.Add(m2, prime)
			}
			m2.Mul(m2, values.R)
			m.Add(m, m2)
		}
	}
	if ir != nil {
		m.Mul(m, ir)
		m.Mod(m, priv.N)
	}

	return
}

// copyWithLeftPad 将源字节切片src的内容复制到目标字节切片dest中，
// 如果dest的长度大于src的长度，则在dest的左侧填充0字节。
// dest和src都不应被nil，且dest的长度应大于等于src的长度。
func copyWithLeftPad(dest, src []byte) {
	numPaddingBytes := len(dest) - len(src)
	for i := 0; i < numPaddingBytes; i++ {
		dest[i] = 0
	}
	copy(dest[numPaddingBytes:], src)
}

// nonZeroRandomBytes 函数使用传入的io.Reader生成指定长度的非零随机字节序列，并填充到传入的切片s中。
// 如果生成随机字节或进行异或操作时发生错误，则返回该错误。
// 参数：
// s - 用于填充随机字节的切片
// rand - 实现了io.Reader接口的对象，用于生成随机字节
// 返回值：
// err - 如果发生错误，返回该错误；否则为nil
func nonZeroRandomBytes(s []byte, rand io.Reader) (err error) {
	_, err = io.ReadFull(rand, s)
	if err != nil {
		return
	}
	for i := 0; i < len(s); i++ {
		for s[i] == 0 {
			_, err = io.ReadFull(rand, s[i:i+1])
			if err != nil {
				return
			}
			s[i] ^= 0x42
		}
	}
	return
}

// leftPad 函数用于在字节切片 input 的左侧填充 0，直到切片长度达到 size
// 参数：
// input: 需要进行填充的字节切片
// size: 填充后的目标长度
// 返回值：
// out: 填充后的字节切片
func leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}

// modInverse 函数计算a关于n的模逆元
// 如果存在模逆元，则返回结果和true，否则返回nil和false
// 参数:
// a: 第一个整数，类型为*big.Int
// n: 第二个整数，类型为*big.Int
// 返回值:
// ia: a关于n的模逆元，类型为*big.Int，如果不存在则返回nil
// ok: 如果找到模逆元则返回true，否则返回false
func modInverse(a, n *big.Int) (ia *big.Int, ok bool) {
	g := new(big.Int)
	x := new(big.Int)
	y := new(big.Int)
	g.GCD(x, y, a, n)
	if g.Cmp(bigOne) != 0 {
		return
	}
	if x.Cmp(bigOne) < 0 {
		x.Add(x, n)
	}
	return x, true
}
