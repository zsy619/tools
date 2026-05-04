package xinterface

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

// ToString
/**
 * @Description: 转换为string类型
 * @param data
 * @return str
 */
func ToString(i interface{}) (str string) {
	switch i := i.(type) {
	case bool:
		str = strconv.FormatBool(i)
	case string:
		str = i
	case uint:
		str = strconv.Itoa(int(i))
	case uint8:
		str = strconv.Itoa(int(i))
	case uint16:
		str = strconv.Itoa(int(i))
	case uint32:
		str = strconv.Itoa(int(i))
	case uint64:
		str = strconv.Itoa(int(i))
	case int:
		str = strconv.Itoa(i)
	case int8:
		str = strconv.Itoa(int(i))
	case int16:
		str = strconv.Itoa(int(i))
	case int32:
		str = string(i)
	case int64:
		str = strconv.FormatInt(i, 10)
	case float32:
		str = fmt.Sprintf("%f", i)
	case float64:
		str = strconv.FormatFloat(i, 'f', -1, 32)
	case time.Time:
		str = i.Format("2006-01-02 15:04:05")
	case []byte:
		b := i
		str = *(*string)(unsafe.Pointer(&b))
	case error:
		str = i.Error()
	default:
		panic("该类型暂不支持")
	}
	return
}

// ToUint
/**
 * @description: 转换为uint类型
 * @param {interface{}} i
 * @return uint error
 */
func ToUint(i interface{}) (num uint, err error) {
	switch i := i.(type) {
	case uint:
		num = i
	case uint8:
		num = uint(i)
	case uint16:
		num = uint(i)
	case uint32:
		num = uint(i)
	case uint64:
		// 有可能造成精度丢失
		num = uint(i)
	case int:
		num = uint(i)
	case int8:
		num = uint(i)
	case int16:
		num = uint(i)
	case int32:
		num = uint(i)
	case int64:
		// 有可能造成精度丢失
		num = uint(i)
	case float32:
		// 有可能造成精度丢失
		num = uint(i)
	case float64:
		// 有可能造成精度丢失
		num = uint(i)
	case string:
		n, e := strconv.Atoi(i)
		num = uint(n)
		err = e
	default:
		panic("该类型暂不支持")
	}
	return
}

// ToUint32 转换为uint32类型
/**
 * @description: 转换为uint32类型
 * @param {interface{}} i
 * @return uint32 error
 */
func ToUint32(i interface{}) (num uint32, err error) {
	switch i := i.(type) {
	case uint:
		num = uint32(i)
	case uint8:
		num = uint32(i)
	case uint16:
		num = uint32(i)
	case uint32:
		num = i
	case uint64:
		// 有可能造成精度丢失
		num = uint32(i)
	case int:
		num = uint32(i)
	case int8:
		num = uint32(i)
	case int16:
		num = uint32(i)
	case int32:
		num = uint32(i)
	case int64:
		// 有可能造成精度丢失
		num = uint32(i)
	case float32:
		// 有可能造成精度丢失
		num = uint32(i)
	case float64:
		// 有可能造成精度丢失
		num = uint32(i)
	case string:
		n, e := strconv.Atoi(i)
		num = uint32(n)
		err = e
	default:
		panic("该类型暂不支持")
	}
	return
}

// ToUint64 *
/**
 * @description: 转换为uint64类型
 * @param {interface{}} i
 * @return uint64 error
 */
func ToUint64(i interface{}) (num uint64, err error) {
	switch i := i.(type) {
	case uint:
		num = uint64(i)
	case uint8:
		num = uint64(i)
	case uint16:
		num = uint64(i)
	case uint32:
		num = uint64(i)
	case uint64:
		num = i
	case int:
		num = uint64(i)
	case int8:
		num = uint64(i)
	case int16:
		num = uint64(i)
	case int32:
		num = uint64(i)
	case int64:
		num = uint64(i)
	case float32:
		// 有可能造成精度丢失
		num = uint64(i)
	case float64:
		// 有可能造成精度丢失
		num = uint64(i)
	case string:
		n, e := strconv.Atoi(i)
		num = uint64(n)
		err = e
	default:
		panic("该类型暂不支持")
	}
	return
}

// ToInt 转换为int类型
func ToInt(i interface{}) (num int, err error) {
	switch i := i.(type) {
	case uint:
		num = int(i)
	case uint8:
		num = int(i)
	case uint16:
		num = int(i)
	case uint32:
		num = int(i)
	case uint64:
		// 有可能造成精度丢失
		num = int(i)
	case int:
		num = i
	case int8:
		num = int(i)
	case int16:
		num = int(i)
	case int32:
		num = int(i)
	case int64:
		// 有可能造成精度丢失
		num = int(i)
	case float32:
		// 有可能造成精度丢失
		num = int(i)
	case float64:
		// 有可能造成精度丢失
		num = int(i)
	case string:
		n, e := strconv.Atoi(i)
		num = int(n)
		err = e
	default:
		panic("该类型暂不支持")
	}
	return
}

// ToInt32
/**
 * @Description: 将任意类型转为int32类型
 * @param i
 * @return num
 * @return err
 */
func ToInt32(i interface{}) (num int32, err error) {
	switch i := i.(type) {
	case uint:
		num = int32(i)
	case uint8:
		num = int32(i)
	case uint16:
		num = int32(i)
	case uint32:
		num = int32(i)
	case uint64:
		// 有可能造成精度丢失
		num = int32(i)
	case int:
		num = int32(i)
	case int8:
		num = int32(i)
	case int16:
		num = int32(i)
	case int32:
		num = i
	case int64:
		// 有可能造成精度丢失
		num = int32(i)
	case float32:
		// 有可能造成精度丢失
		num = int32(i)
	case float64:
		// 有可能造成精度丢失
		num = int32(i)
	case string:
		n, e := strconv.Atoi(i)
		num = int32(n)
		err = e
	default:
		panic("该类型暂不支持")
	}
	return
}

// ToInt64
/**
 * @Description: 将任意类型转为int64类型
 * @param i
 * @return num
 * @return err
 */
func ToInt64(i interface{}) (num int64, err error) {
	switch i := i.(type) {
	case uint:
		num = int64(i)
	case uint8:
		num = int64(i)
	case uint16:
		num = int64(i)
	case uint32:
		num = int64(i)
	case uint64:
		num = int64(i)
	case int:
		num = int64(i)
	case int8:
		num = int64(i)
	case int16:
		num = int64(i)
	case int32:
		num = int64(i)
	case int64:
		num = i
	case float32:
		num = int64(i)
	case float64:
		num = int64(i)
	case string:
		num, err = strconv.ParseInt(i, 10, 64)
	default:
		panic("该类型暂不支持")
	}
	return
}

// ToFloat32
/**
 * @Description: 将任意类型转为float32类型
 * @param i
 * @return num
 * @return err
 */
func ToFloat32(i interface{}) (num float32, err error) {
	switch i := i.(type) {
	case string:
		// string无法直接转换float32，只能先转换为float64，再通过float64转float32
		var num64 float64
		num64, err = strconv.ParseFloat(i, 32)
		num = float32(num64)
	case uint:
		num = float32(i)
	case uint8:
		num = float32(i)
	case uint16:
		num = float32(i)
	case uint32:
		num = float32(i)
	case uint64:
		num = float32(i)
	case int:
		num = float32(i)
	case int8:
		num = float32(i)
	case int16:
		num = float32(i)
	case int32:
		num = float32(i)
	case int64:
		num = float32(i)
	case float32:
		num = i
	case float64:
		// 可能造成精度丢失
		num = float32(i)
	default:
		panic("该类型暂不支持")
	}
	return
}

// ToFloat64
/**
 * @Description: 将任意类型转为float64类型
 * @param i
 * @return num
 * @return err
 */
func ToFloat64(i interface{}) (num float64, err error) {
	switch i := i.(type) {
	case string:
		num, err = strconv.ParseFloat(i, 64)
	case uint:
		num = float64(i)
	case uint8:
		num = float64(i)
	case uint16:
		num = float64(i)
	case uint32:
		num = float64(i)
	case uint64:
		num = float64(i)
	case int:
		num = float64(i)
	case int8:
		num = float64(i)
	case int16:
		num = float64(i)
	case int32:
		num = float64(i)
	case int64:
		num = float64(i)
	case float32:
		num = float64(i)
	case float64:
		num = i
	default:
		panic("该类型暂不支持")
	}
	return
}

// ToByteArray
/**
 *  @Description: 转换为[]byte
 *  @param i
 *  @return b
 */
func ToByteArray(i interface{}) (b []byte) {
	switch i := i.(type) {
	case int:
		input := i
		return IntToBytesBigEndian(input)
	case int32:
		input := i
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, uint32(input))
		return buf
	case string:
		str := i
		return *(*[]byte)(unsafe.Pointer(&str))
	default:
		panic("该类型暂不支持")
	}
}

func IntToBytesBigEndian(x int) []byte {
	var bytes []byte
	if unsafe.Sizeof(x) == 4 { // 检查系统架构是32位还是64位
		bytes = make([]byte, 4)
		// 大端序：最高位字节在前
		bytes[0] = byte(x >> 24)
		bytes[1] = byte(x >> 16)
		bytes[2] = byte(x >> 8)
		bytes[3] = byte(x)
	} else if unsafe.Sizeof(x) == 8 {
		bytes = make([]byte, 8)
		// 大端序：最高位字节在前
		bytes[0] = byte(x >> 56)
		bytes[1] = byte(x >> 48)
		bytes[2] = byte(x >> 40)
		bytes[3] = byte(x >> 32)
		bytes[4] = byte(x >> 24)
		bytes[5] = byte(x >> 16)
		bytes[6] = byte(x >> 8)
		bytes[7] = byte(x)
	}
	return bytes
}

func IntToBytesLittleEndian(x int) []byte {
	var bytes []byte
	if unsafe.Sizeof(x) == 4 {
		bytes = make([]byte, 4)
		// 小端序：最低位字节在前
		bytes[0] = byte(x)
		bytes[1] = byte(x >> 8)
		bytes[2] = byte(x >> 16)
		bytes[3] = byte(x >> 24)
	} else if unsafe.Sizeof(x) == 8 {
		bytes = make([]byte, 8)
		// 小端序：最低位字节在前
		bytes[0] = byte(x)
		bytes[1] = byte(x >> 8)
		bytes[2] = byte(x >> 16)
		bytes[3] = byte(x >> 24)
		bytes[4] = byte(x >> 32)
		bytes[5] = byte(x >> 40)
		bytes[6] = byte(x >> 48)
		bytes[7] = byte(x >> 56)
	}
	return bytes
}
