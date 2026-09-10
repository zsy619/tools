// Package xinterface/change_format.go 提供一组将 interface{} 强转为各种
// 目标类型的工具函数，覆盖字符串、整型、浮点、时间、字节等基础类型。
//
// 设计动机：在弱类型场景（如解析 JSON、读取数据库可变字段）下，
// 调用方拿到的是 interface{}，需要做一系列类型断言才能得到想要的值。
// 本文件将这些常用断言集中到 ToXxx 函数中，配合类型 switch 给出明确
// 的支持范围与失败时的 panic 信息。
//
// 注意：
//   - 这些函数对未支持的类型会 panic，请仅在已知类型集合的代码中使用；
//   - 在公开 API 边界（解析外部输入）建议改为返回 (T, error) 的版本。
package xinterface

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

// ToString 将任意 interface{} 转换为字符串。
//
// 支持的类型：bool、string、uint/uint8/16/32/64、int/int8/16/32/64、
// float32、float64、time.Time、[]byte、error。
//
// 行为细节：
//   - 数值类型使用 strconv.Itoa/FormatInt/FormatFloat 转换；
//   - time.Time 使用 "2006-01-02 15:04:05" 格式；
//   - []byte 通过 unsafe 零拷贝转 string，要求调用方保证 b 不可变；
//   - error 直接调用 Error()。
//
// 未支持的类型会触发 panic("该类型暂不支持")。
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

// ToUint 将任意类型转换为 uint。
//
// 支持的来源类型：uint/uint8/16/32/64、int/int8/16/32/64、float32/64、
// string（通过 strconv.Atoi 解析十进制）。
//
// 注意事项：
//   - 负数 -> uint 会按位补码解释为一个非常大的正数；
//   - float 截断小数部分；NaN / Inf 行为未定义；
//   - 64 位源类型向 uint 收窄时可能发生精度丢失。
//
// 仅 string 分支会返回错误（解析失败），其余分支始终返回 nil error。
// 未支持的类型触发 panic("该类型暂不支持")。
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

// ToUint32 将任意类型转换为 uint32。
//
// 支持的来源类型与 ToUint 类似，但目标宽度收窄到 32 位。
// 当来源类型为 64 位（uint64/int64/float64）时可能发生精度或溢出
// 截断。
//
// 仅 string 分支会返回错误，未支持的类型触发 panic。
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

// ToUint64 将任意类型转换为 uint64。
//
// 支持的来源类型与 ToUint 类似，目标是 64 位无符号整数，
// 一般不会出现溢出（除非 int64 负值）。
//
// 仅 string 分支会返回错误，未支持的类型触发 panic。
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

// ToInt 将任意类型转换为 int。
//
// 支持的来源类型与 ToUint 类似，但目标是有符号整型。
//
// 注意事项：
//   - 64 位源类型向 int 收窄时可能发生精度丢失；
//   - float 截断小数部分；NaN / Inf 行为未定义。
//
// 仅 string 分支会返回错误，未支持的类型触发 panic。
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

// ToInt32 将任意类型转换为 int32。
//
// 支持的来源类型与 ToInt 类似，目标宽度收窄到 32 位。
// 当来源类型为 64 位（uint64/int64/float64）时可能发生精度或溢出
// 截断。
//
// 仅 string 分支会返回错误，未支持的类型触发 panic。
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

// ToInt64 将任意类型转换为 int64。
//
// 支持的来源类型：所有整型、浮点型以及 string（通过 strconv.ParseInt
// 按十进制解析）。
//
// 推荐用于存储数据库主键、雪花 ID、时间戳等场景，因为 int64 是 Go
// 在大多数架构上最自然的「整数容器」。
//
// 仅 string 分支会返回错误，未支持的类型触发 panic。
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

// ToFloat32 将任意类型转换为 float32。
//
// 支持的来源类型：所有整型、float32/64、string（通过 strconv.ParseFloat
// 以 float32 精度解析）。
//
// 从 float64 收窄到 float32 时可能发生精度丢失；string 解析失败时
// 返回错误，未支持的类型触发 panic。
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

// ToFloat64 将任意类型转换为 float64。
//
// 支持的来源类型：所有整型、float32/64、string（通过 strconv.ParseFloat
// 按 64 位精度解析）。
//
// float64 是 Go 默认的浮点宽度，推荐在不需要极致性能或节省内存的场景
// 使用；string 解析失败时返回错误，未支持的类型触发 panic。
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

// ToByteArray 将任意类型转换为 []byte。
//
// 支持的来源类型：
//   - int：调用 IntToBytesBigEndian 输出大端字节序；
//   - int32：使用 encoding/binary.BigEndian 写入 4 字节；
//   - string：通过 unsafe 零拷贝切片化，要求 s 不可变。
//
// 未支持的类型触发 panic("该类型暂不支持")。
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

// IntToBytesBigEndian 将 int 转换为大端序字节序列。
//
// 根据 unsafe.Sizeof(x) 自动选择 4 字节（32 位平台）或 8 字节
// （64 位平台）的输出长度。
//
// 大端序在网络协议、文件格式（如 JPEG、PNG）以及跨平台传输中非常常用。
//
// 注意：返回的切片是新分配的，调用方可以自由修改。
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

// IntToBytesLittleEndian 将 int 转换为小端序字节序列。
//
// 与 IntToBytesBigEndian 对称，适用于 x86/ARM 等原生小端平台的数据
// 读写或与其它小端系统交互的场景。
//
// 注意：返回的切片是新分配的，调用方可以自由修改。
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
