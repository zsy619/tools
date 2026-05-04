package xcrypto

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"hash"
	"io"
	"math/big"

	"github.com/tjfoc/gmsm/sm2"
)

// https://blog.csdn.net/weixin_42704356/article/details/129669531?spm=1001.2101.3001.6650.5&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromBaidu%7ERate-5-129669531-blog-130558002.235%5Ev38%5Epc_relevant_anti_vip&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromBaidu%7ERate-5-129669531-blog-130558002.235%5Ev38%5Epc_relevant_anti_vip&utm_relevant_index=6

const (
	BitSize    = 256
	KeyBytes   = (BitSize + 7) / 8
	UnCompress = 0x04
)

type Sm2CipherTextType int32

const (
	// 旧标准的密文顺序
	C1C2C3 Sm2CipherTextType = 1
	// [GM/T 0009-2012]标准规定的顺序
	C1C3C2 Sm2CipherTextType = 2
)

var (
	sm2H                 = new(big.Int).SetInt64(1)
	sm2SignDefaultUserId = []byte{
		0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
		0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
	}
)
var sm2P256V1 P256V1Curve

type P256V1Curve struct {
	*elliptic.CurveParams
	A *big.Int
}

func init() {
	initSm2P256V1()
}

// initSm2P256V1 初始化SM2-P-256-V1椭圆曲线参数
func initSm2P256V1() {
	sm2P, _ := new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF00000000FFFFFFFFFFFFFFFF", 16)
	sm2A, _ := new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF00000000FFFFFFFFFFFFFFFC", 16)
	sm2B, _ := new(big.Int).SetString("28E9FA9E9D9F5E344D5A9E4BCF6509A7F39789F515AB8F92DDBCBD414D940E93", 16)
	sm2N, _ := new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFF7203DF6B21C6052B53BBF40939D54123", 16)
	sm2Gx, _ := new(big.Int).SetString("32C4AE2C1F1981195F9904466A39C9948FE30BBFF2660BE1715A4589334C74C7", 16)
	sm2Gy, _ := new(big.Int).SetString("BC3736A2F4F6779C59BDCEE36B692153D0A9877CC62A474002DF32E52139F0A0", 16)
	sm2P256V1.CurveParams = &elliptic.CurveParams{Name: "SM2-P-256-V1"}
	sm2P256V1.P = sm2P
	sm2P256V1.A = sm2A
	sm2P256V1.B = sm2B
	sm2P256V1.N = sm2N
	sm2P256V1.Gx = sm2Gx
	sm2P256V1.Gy = sm2Gy
	sm2P256V1.BitSize = BitSize
}

// GetSm2P256V1 函数返回一个P256V1Curve类型的实例，该实例对应于SM2 P-256 V1椭圆曲线
func GetSm2P256V1() P256V1Curve {
	return sm2P256V1
}

// RawBytesToPubKey 将SM2公钥的原始字节转换为sm2.PublicKey结构体指针
//
// 参数：
//
//	bytes []byte - 公钥的原始字节，长度必须为KeyBytes*2
//
// 返回值：
//
//	*sm2.PublicKey - 转换后的公钥结构体指针
//	error - 如果原始字节长度不符合要求，则返回错误；否则返回nil
func RawBytesToPubKey(bytes []byte) (*sm2.PublicKey, error) {
	if len(bytes) != KeyBytes*2 {
		return nil, fmt.Errorf("public key raw bytes length must be %d", KeyBytes*2)
	}
	publicKey := new(sm2.PublicKey)
	publicKey.Curve = sm2P256V1
	publicKey.X = new(big.Int).SetBytes(bytes[:KeyBytes])
	publicKey.Y = new(big.Int).SetBytes(bytes[KeyBytes:])
	return publicKey, nil
}

// RawBytesToPriKey 函数将原始的字节数组转换为sm2.PrivateKey指针。
// 参数：
//
//	bytes []byte - 要转换的原始字节数组
//
// 返回值：
//
//	*sm2.PrivateKey - 转换得到的私钥对象指针
//	error - 如果转换过程中发生错误，则返回非零的错误码
func RawBytesToPriKey(bytes []byte) (*sm2.PrivateKey, error) {
	if len(bytes) != KeyBytes {
		return nil, fmt.Errorf("private key raw bytes length must be %d", KeyBytes)
	}
	privateKey := new(sm2.PrivateKey)
	privateKey.Curve = sm2P256V1
	privateKey.D = new(big.Int).SetBytes(bytes)
	return privateKey, nil
}

// GetPriKeyRawBytes 将sm2.PrivateKey的D字段转换为长度为KeyBytes的字节切片并返回
// 如果D字段的长度大于KeyBytes，则返回D字段的最后KeyBytes个字节
// 如果D字段的长度小于KeyBytes，则在返回的字节切片的前面用0填充至KeyBytes长度
// 参数：
//
//	priKey: sm2.PrivateKey类型的指针，表示要获取私钥字节切片的私钥对象
//
// 返回值：
//
//	[]byte：表示私钥的字节切片
func GetPriKeyRawBytes(priKey *sm2.PrivateKey) []byte {
	dBytes := bigIntTo32Bytes(priKey.D)
	dl := len(dBytes)
	if dl > KeyBytes {
		raw := make([]byte, KeyBytes)
		copy(raw, dBytes[dl-KeyBytes:])
		return raw
	} else if dl < KeyBytes {
		raw := make([]byte, KeyBytes)
		copy(raw[KeyBytes-dl:], dBytes)
		return raw
	} else {
		return dBytes
	}
}

func GetPubKeyRawBytes(pubKey *sm2.PublicKey) []byte {
	raw := GetUnCompressBytes(pubKey)
	return raw[1:]
}

// GetUnCompressBytes 将SM2公钥转换为未压缩格式的字节切片
// pubKey：指向sm2.PublicKey类型的指针，表示要转换的公钥
// 返回值：未压缩格式的公钥字节切片
func GetUnCompressBytes(pubKey *sm2.PublicKey) []byte {
	xBytes := bigIntTo32Bytes(pubKey.X)
	yBytes := bigIntTo32Bytes(pubKey.Y)
	xl := len(xBytes)
	yl := len(yBytes)

	raw := make([]byte, 1+KeyBytes*2)
	raw[0] = UnCompress
	if xl > KeyBytes {
		copy(raw[1:1+KeyBytes], xBytes[xl-KeyBytes:])
	} else if xl < KeyBytes {
		copy(raw[1+(KeyBytes-xl):1+KeyBytes], xBytes)
	} else {
		copy(raw[1:1+KeyBytes], xBytes)
	}

	if yl > KeyBytes {
		copy(raw[1+KeyBytes:], yBytes[yl-KeyBytes:])
	} else if yl < KeyBytes {
		copy(raw[1+KeyBytes+(KeyBytes-yl):], yBytes)
	} else {
		copy(raw[1+KeyBytes:], yBytes)
	}
	return raw
}

// nextK 从给定的 io.Reader 中生成一个大于等于 1 且小于 max 的随机大整数
//
// rnd: 用于生成随机数的 io.Reader
// max: 生成随机数的上界（不包含该值）
//
// 返回值：
// - *big.Int: 生成的随机大整数
// - error: 如果发生错误，则返回非零的错误码
func nextK(rnd io.Reader, max *big.Int) (*big.Int, error) {
	intOne := new(big.Int).SetInt64(1)
	var k *big.Int
	var err error
	for {
		k, err = rand.Int(rnd, max)
		if err != nil {
			return nil, err
		}
		if k.Cmp(intOne) >= 0 {
			return k, err
		}
	}
}

// xor 函数接受三个参数：
// data: 需要进行异或操作的字节切片
// kdfOut: 与data进行异或操作的另一个字节切片
// dRemaining: 需要进行异或操作的字节数量
//
// 函数将data中前dRemaining个字节与kdfOut中对应位置的字节进行异或操作
func xor(data []byte, kdfOut []byte, dRemaining int) {
	for i := 0; i != dRemaining; i++ {
		data[i] ^= kdfOut[i]
	}
}

// 表示SM2 Key的大数比较小时，直接通过Bytes()函数得到的字节数组可能不够32字节，这个时候要补齐成32字节
func bigIntTo32Bytes(bn *big.Int) []byte {
	byteArr := bn.Bytes()
	byteArrLen := len(byteArr)
	if byteArrLen == KeyBytes {
		return byteArr
	}
	byteArr = append(make([]byte, KeyBytes-byteArrLen), byteArr...)
	return byteArr
}

// kdf 是一个用于密钥派生的函数
// 它使用给定的哈希函数 digest、两个大整数 c1x 和 c1y 以及加密数据 encData 来生成派生密钥
// 参数：
//
//	digest: 用于生成派生密钥的哈希函数
//	c1x: 第一个大整数参数
//	c1y: 第二个大整数参数
//	encData: 加密数据，将被用于密钥派生过程中
func kdf(digest hash.Hash, c1x *big.Int, c1y *big.Int, encData []byte) {
	bufSize := 4
	if bufSize < digest.Size() {
		bufSize = digest.Size()
	}
	buf := make([]byte, bufSize)

	encDataLen := len(encData)
	c1xBytes := bigIntTo32Bytes(c1x)
	c1yBytes := bigIntTo32Bytes(c1y)
	off := 0
	ct := uint32(0)
	for off < encDataLen {
		digest.Reset()
		digest.Write(c1xBytes)
		digest.Write(c1yBytes)
		ct++
		binary.BigEndian.PutUint32(buf, ct)
		digest.Write(buf[:4])
		tmp := digest.Sum(nil)
		copy(buf[:bufSize], tmp[:bufSize])

		xorLen := encDataLen - off
		if xorLen > digest.Size() {
			xorLen = digest.Size()
		}
		xor(encData[off:], buf, xorLen)
		off += xorLen
	}
}

// notEncrypted 函数检查传入的加密数据（encData）是否未经加密（即是否与原始数据（in）相同）
// 如果加密数据与原始数据完全相同，则返回true，否则返回false
// encData：加密后的数据，类型为[]byte
// in：原始数据，与encData进行对比，类型为[]byte
// 返回值：如果encData与in完全相同，返回true，否则返回false
func notEncrypted(encData []byte, in []byte) bool {
	encDataLen := len(encData)
	for i := 0; i != encDataLen; i++ {
		if encData[i] != in[i] {
			return false
		}
	}
	return true
}
