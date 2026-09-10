package xcrypto

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"hash/fnv"
	"io"
	"reflect"

	"github.com/cespare/xxhash"

	"github.com/zsy619/tools/xjson"
)

// 参考来源：https://github.com/pinpt/go-common/blob/master/hash/hash.go

// Values 会将所有对象转换为字符串，并返回拼接值经过哈希运算后的校验值。
// 它使用 xxhash 计算更快的哈希值；该哈希不具备密码学安全性，但对我们来说已经足够，
// 因为我们主要利用哈希生成一致的键值或进行相等性检查。
func Values(objects ...interface{}) string {
	return hashValues(objects...)
}

// Modulo 返回 sha 值对 num 取模后的余数
func Modulo(sha string, num int) int {
	hasher := fnv.New32a()
	_, _ = hasher.Write([]byte(sha))
	partition := int(hasher.Sum32()) % num
	if partition < 0 {
		partition = -partition
	}
	return partition
}

func hashValues(objects ...interface{}) string {
	h := xxhash.New()

	// 这是一个类型断言；遗憾的是，指针必须单独处理。
	// 注：看起来 fmt.Fprintf 是仅次于 _, _ = io.WriteString 的第二快写法。
	for _, o := range objects {
		if o == nil {
			_, _ = io.WriteString(h, "")
			continue
		}
		switch s := o.(type) {
		case string:
			_, _ = io.WriteString(h, s)
		case []byte:
			_, _ = h.Write(s)
		case []string:
			for _, v := range s {
				_, _ = io.WriteString(h, v)
			}
		case bool:
			if s {
				_, _ = io.WriteString(h, "true")
			} else {
				_, _ = io.WriteString(h, "false")
			}
		case int, int8, int16, int32, int64:
			fmt.Fprintf(h, "%d", s)
		case float32:
			// 如果是形如 123.00 的浮点数，则去掉小数部分
			if s == float32(int32(s)) {
				fmt.Fprintf(h, "%d", int32(s))
			} else {
				fmt.Fprintf(h, "%f", s)
			}
		case float64:
			// 如果是形如 123.00 的浮点数，则去掉小数部分
			if s == float64(int64(s)) {
				fmt.Fprintf(h, "%d", int64(s))
			} else {
				fmt.Fprintf(h, "%f", s)
			}
		case *string:
			if s == nil {
				_, _ = io.WriteString(h, "")
			} else {
				_, _ = io.WriteString(h, *s)
			}
		case *int:
			if s == nil {
				_, _ = io.WriteString(h, "")
			} else {
				fmt.Fprintf(h, "%d", *s)
			}
		case *int8:
			if s == nil {
				_, _ = io.WriteString(h, "")
			} else {
				fmt.Fprintf(h, "%d", *s)
			}
		case *int16:
			if s == nil {
				_, _ = io.WriteString(h, "")
			} else {
				fmt.Fprintf(h, "%d", *s)
			}
		case *int32:
			if s == nil {
				_, _ = io.WriteString(h, "")
			} else {
				fmt.Fprintf(h, "%d", *s)
			}
		case *int64:
			if s == nil {
				_, _ = io.WriteString(h, "")
			} else {
				fmt.Fprintf(h, "%d", *s)
			}
		case *float32:
			if s == nil {
				_, _ = io.WriteString(h, "")
			} else {
				// 如果是形如 123.00 的浮点数，则去掉小数部分
				if *s == float32(int32(*s)) {
					fmt.Fprintf(h, "%d", int32(*s))
				} else {
					fmt.Fprintf(h, "%f", *s)
				}
			}
		case *float64:
			if s == nil {
				_, _ = io.WriteString(h, "")
			} else {
				// 如果是形如 123.00 的浮点数，则去掉小数部分
				if *s == float64(int64(*s)) {
					fmt.Fprintf(h, "%d", int64(*s))
				} else {
					fmt.Fprintf(h, "%f", *s)
				}
			}
		case *bool:
			if s == nil {
				_, _ = io.WriteString(h, "")
			} else {
				fmt.Fprintf(h, "%v", *s)
			}
		default:
			t := reflect.TypeOf(s)
			k := t.Kind()
			if k == reflect.Ptr {
				s = reflect.ValueOf(s).Interface()
				k = reflect.Interface
			}
			switch k {
			case reflect.Struct, reflect.Slice, reflect.Interface:
				buf, _ := xjson.Marshal(s)
				_, _ = h.Write(buf)
			default:
				// fmt.Println(reflect.TypeOf(s), reflect.TypeOf(s).Kind(), fmt.Sprintf("%v", s))
				fmt.Fprintf(h, "%v", s)
			}
		}
	}

	return fmt.Sprintf("%016x", h.Sum64())
}

// Sha256Checksum 返回 r 的 sha256 校验和
func Sha256Checksum(r io.Reader) ([]byte, error) {
	_, sum, err := ChecksumFrom(r, sha256.New())
	return sum, err
}

// ChecksumFrom 每次从 r 读取 8096 字节写入 hasher，最后返回写入的字节总数与校验和
func ChecksumFrom(r io.Reader, hasher hash.Hash) (int64, []byte, error) {
	var written int
	for {
		buf := make([]byte, 8096)
		n, err := r.Read(buf)
		written += n
		if err == io.EOF || n == 0 {
			break
		} else if err != nil {
			return 0, nil, err
		}
		n, err = hasher.Write(buf[0:n])
		if err != nil {
			return 0, nil, err
		}
		fmt.Println("ChecksumFrom--->", n)
	}
	return int64(written), hasher.Sum(nil), nil
}

// ChecksumCopy 与 io.Copy 行为相同，但同时返回所读数据的 sha256 校验和
func ChecksumCopy(dst io.Writer, src io.Reader) (int64, []byte, error) {
	return ChecksumFrom(io.TeeReader(src, dst), sha256.New())
}
