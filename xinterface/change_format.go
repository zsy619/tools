package xinterface

import (
	"bytes"
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
	switch i.(type) {
	case bool:
		str = strconv.FormatBool(i.(bool))
	case string:
		str = i.(string)
	case uint:
		str = strconv.Itoa(int(i.(uint)))
	case uint8:
		str = strconv.Itoa(int(i.(uint8)))
	case uint16:
		str = strconv.Itoa(int(i.(uint16)))
	case uint32:
		str = strconv.Itoa(int(i.(uint32)))
	case uint64:
		str = strconv.Itoa(int(i.(uint64)))
	case int:
		str = strconv.Itoa(i.(int))
	case int8:
		str = strconv.Itoa(int(i.(int8)))
	case int16:
		str = strconv.Itoa(int(i.(int16)))
	case int32:
		str = string(i.(int32))
	case int64:
		str = strconv.FormatInt(i.(int64), 10)
	case float32:
		str = fmt.Sprintf("%f", i.(float32))
	case float64:
		str = strconv.FormatFloat(i.(float64), 'f', -1, 32)
	case time.Time:
		str = i.(time.Time).Format("2006-01-02 15:04:05")
	case []byte:
		b := i.([]byte)
		str = *(*string)(unsafe.Pointer(&b))
	case error:
		str = i.(error).Error()
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
	switch i.(type) {
	case uint:
		num = i.(uint)
	case uint8:
		num = uint(i.(uint8))
	case uint16:
		num = uint(i.(uint16))
	case uint32:
		num = uint(i.(uint32))
	case uint64:
		// 有可能造成精度丢失
		num = uint(i.(uint64))
	case int:
		num = uint(i.(int))
	case int8:
		num = uint(i.(int8))
	case int16:
		num = uint(i.(int16))
	case int32:
		num = uint(i.(int32))
	case int64:
		// 有可能造成精度丢失
		num = uint(i.(int64))
	case float32:
		// 有可能造成精度丢失
		num = uint(i.(float32))
	case float64:
		// 有可能造成精度丢失
		num = uint(i.(float64))
	case string:
		n, e := strconv.Atoi(i.(string))
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
	switch i.(type) {
	case uint:
		num = uint32(i.(uint))
	case uint8:
		num = uint32(i.(uint8))
	case uint16:
		num = uint32(i.(uint16))
	case uint32:
		num = i.(uint32)
	case uint64:
		// 有可能造成精度丢失
		num = uint32(i.(uint64))
	case int:
		num = uint32(i.(int))
	case int8:
		num = uint32(i.(int8))
	case int16:
		num = uint32(i.(int16))
	case int32:
		num = uint32(i.(int32))
	case int64:
		// 有可能造成精度丢失
		num = uint32(i.(int64))
	case float32:
		// 有可能造成精度丢失
		num = uint32(i.(float32))
	case float64:
		// 有可能造成精度丢失
		num = uint32(i.(float64))
	case string:
		n, e := strconv.Atoi(i.(string))
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
	switch i.(type) {
	case uint:
		num = uint64(i.(uint))
	case uint8:
		num = uint64(i.(uint8))
	case uint16:
		num = uint64(i.(uint16))
	case uint32:
		num = uint64(i.(uint32))
	case uint64:
		num = i.(uint64)
	case int:
		num = uint64(i.(int))
	case int8:
		num = uint64(i.(int8))
	case int16:
		num = uint64(i.(int16))
	case int32:
		num = uint64(i.(int32))
	case int64:
		num = uint64(i.(int64))
	case float32:
		// 有可能造成精度丢失
		num = uint64(i.(float32))
	case float64:
		// 有可能造成精度丢失
		num = uint64(i.(float64))
	case string:
		n, e := strconv.Atoi(i.(string))
		num = uint64(n)
		err = e
	default:
		panic("该类型暂不支持")
	}
	return
}

// ToInt 转换为int类型
func ToInt(i interface{}) (num int, err error) {
	switch i.(type) {
	case uint:
		num = int(i.(uint))
	case uint8:
		num = int(i.(uint8))
	case uint16:
		num = int(i.(uint16))
	case uint32:
		num = int(i.(uint32))
	case uint64:
		// 有可能造成精度丢失
		num = int(i.(uint64))
	case int:
		num = i.(int)
	case int8:
		num = int(i.(int8))
	case int16:
		num = int(i.(int16))
	case int32:
		num = int(i.(int32))
	case int64:
		// 有可能造成精度丢失
		num = int(i.(int64))
	case float32:
		// 有可能造成精度丢失
		num = int(i.(float32))
	case float64:
		// 有可能造成精度丢失
		num = int(i.(float64))
	case string:
		n, e := strconv.Atoi(i.(string))
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
	switch i.(type) {
	case uint:
		num = int32(i.(uint))
	case uint8:
		num = int32(i.(uint8))
	case uint16:
		num = int32(i.(uint16))
	case uint32:
		num = int32(i.(uint32))
	case uint64:
		// 有可能造成精度丢失
		num = int32(i.(uint64))
	case int:
		num = int32(i.(int))
	case int8:
		num = int32(i.(int8))
	case int16:
		num = int32(i.(int16))
	case int32:
		num = i.(int32)
	case int64:
		// 有可能造成精度丢失
		num = int32(i.(int64))
	case float32:
		// 有可能造成精度丢失
		num = int32(i.(float32))
	case float64:
		// 有可能造成精度丢失
		num = int32(i.(float64))
	case string:
		n, e := strconv.Atoi(i.(string))
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
	switch i.(type) {
	case uint:
		num = int64(i.(uint))
	case uint8:
		num = int64(i.(uint8))
	case uint16:
		num = int64(i.(uint16))
	case uint32:
		num = int64(i.(uint32))
	case uint64:
		num = int64(i.(uint64))
	case int:
		num = int64(i.(int))
	case int8:
		num = int64(i.(int8))
	case int16:
		num = int64(i.(int16))
	case int32:
		num = int64(i.(int32))
	case int64:
		num = i.(int64)
	case float32:
		num = int64(i.(float32))
	case float64:
		num = int64(i.(float64))
	case string:
		num, err = strconv.ParseInt(i.(string), 10, 64)
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
	switch i.(type) {
	case string:
		// string无法直接转换float32，只能先转换为float64，再通过float64转float32
		var num64 float64
		num64, err = strconv.ParseFloat(i.(string), 32)
		num = float32(num64)
	case uint:
		num = float32(i.(uint))
	case uint8:
		num = float32(i.(uint8))
	case uint16:
		num = float32(i.(uint16))
	case uint32:
		num = float32(i.(uint32))
	case uint64:
		num = float32(i.(uint64))
	case int:
		num = float32(i.(int))
	case int8:
		num = float32(i.(int8))
	case int16:
		num = float32(i.(int16))
	case int32:
		num = float32(i.(int32))
	case int64:
		num = float32(i.(int64))
	case float32:
		num = i.(float32)
	case float64:
		// 可能造成精度丢失
		num = float32(i.(float64))
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
	switch i.(type) {
	case string:
		num, err = strconv.ParseFloat(i.(string), 64)
	case uint:
		num = float64(i.(uint))
	case uint8:
		num = float64(i.(uint8))
	case uint16:
		num = float64(i.(uint16))
	case uint32:
		num = float64(i.(uint32))
	case uint64:
		num = float64(i.(uint64))
	case int:
		num = float64(i.(int))
	case int8:
		num = float64(i.(int8))
	case int16:
		num = float64(i.(int16))
	case int32:
		num = float64(i.(int32))
	case int64:
		num = float64(i.(int64))
	case float32:
		num = float64(i.(float32))
	case float64:
		num = i.(float64)
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
	switch i.(type) {
	case int:
		input := i.(int)
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, input)
		if err != nil {
			panic(err)
		}
		return buf.Bytes()
	case int32:
		input := i.(int32)
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, uint32(input))
		return buf
	case string:
		str := i.(string)
		return *(*[]byte)(unsafe.Pointer(&str))
	default:
		panic("该类型暂不支持")
	}
}
