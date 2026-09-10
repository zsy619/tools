// Package xmap 提供一组常用的 map 工具类型与函数。
//
// 该包主要面向“字符串键 -> interface{} 值”以及“泛型键值对”
// 两种使用场景：
//
//   - define.go 中导出别名类型 H 及其一系列类型安全的 Get/Set
//     辅助方法，适合在解析配置文件、HTTP 请求参数等弱类型 JSON
//     数据时使用；
//   - map.go 中导出线程安全的泛型 Map[K, V]，是对 Go 内置
//     map 的常见增强（线程安全、Keys/Values、Min/Max 等）；
//   - tool.go 中提供与具体类型无关的 DeepCopy、Clone、AddAll
//     与 Clear 等通用工具函数。
//
// 该包刻意保持轻量，仅依赖标准库与同仓库内的 xreflect / tools。
package xmap

import "errors"

// H 是 map[string]interface{} 的语义化别名，配套提供了一组类型安全的
// 读取与写入方法，便于在解析 JSON、配置或 HTTP 参数等弱类型场景中
// 直接使用。
//
// 由于底层是普通 map，H 本身不是线程安全的；如需并发读写，请使用
// map.go 中的 Map[K, V] 或自行加锁。
type H map[string]interface{}

// Get 返回 key 对应的原始值，不做类型转换。
//
// 当 key 不存在或 h 为 nil 时返回 nil（与原生 map 行为一致）。
//
// 示例：
//
//	h := xmap.H{"name": "alice", "age": 18}
//	v := h.Get("name") // v == "alice"
//	v = h.Get("miss")  // v == nil
func (h H) Get(key string) interface{} {
	return h[key]
}

// GetString 返回 key 对应的字符串值。
//
// 如果 key 不存在或对应值的类型不是 string，会通过类型断言触发 panic，
// 因此仅在可以确保类型正确时使用；不确定时建议改用类型安全、带 error
// 返回的 GetXxx 系列方法。
func (h H) GetString(key string) string {
	return h[key].(string)
}

// GetInt 读取 key 对应的 int 值。
//
// 当 key 不存在、h 为 nil、或者对应值的类型不是 int 时，
// 返回 (0, errors.New("not int"))。
//
// 注意：JSON 反序列化时数字通常被解析为 float64，因此直接读取
// 数字字段时更推荐 GetFloat64 或先做类型转换。
func (h H) GetInt(key string) (int, error) {
	out := h[key]
	outInt, ok := out.(int)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not int")
}

// GetInt32 读取 key 对应的 int32 值。
//
// 行为与 GetInt 一致：类型不匹配时返回 (0, errors.New("not int32"))。
func (h H) GetInt32(key string) (int32, error) {
	out := h[key]
	outInt, ok := out.(int32)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not int32")
}

// GetInt16 读取 key 对应的 int16 值。
//
// 行为与 GetInt 一致：类型不匹配时返回 (0, errors.New("not int16"))。
func (h H) GetInt16(key string) (int16, error) {
	out := h[key]
	outInt, ok := out.(int16)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not int16")
}

// GetInt8 读取 key 对应的 int8 值。
//
// 行为与 GetInt 一致：类型不匹配时返回 (0, errors.New("not int8"))。
func (h H) GetInt8(key string) (int8, error) {
	out := h[key]
	outInt, ok := out.(int8)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not int8")
}

// GetUint 读取 key 对应的 uint 值。
//
// 行为与 GetInt 一致：类型不匹配时返回 (0, errors.New("not uint"))。
func (h H) GetUint(key string) (uint, error) {
	out := h[key]
	outInt, ok := out.(uint)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not uint")
}

// GetUint32 读取 key 对应的 uint32 值。
//
// 行为与 GetInt 一致：类型不匹配时返回 (0, errors.New("not uint32"))。
func (h H) GetUint32(key string) (uint32, error) {
	out := h[key]
	outInt, ok := out.(uint32)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not uint32")
}

// GetUint16 读取 key 对应的 uint16 值。
//
// 行为与 GetInt 一致：类型不匹配时返回 (0, errors.New("not uint16"))。
func (h H) GetUint16(key string) (uint16, error) {
	out := h[key]
	outInt, ok := out.(uint16)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not uint16")
}

// GetUint8 读取 key 对应的 uint8 值。
//
// 行为与 GetInt 一致：类型不匹配时返回 (0, errors.New("not uint8"))。
func (h H) GetUint8(key string) (uint8, error) {
	out := h[key]
	outInt, ok := out.(uint8)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not uint8")
}

// GetInt64 读取 key 对应的 int64 值。
//
// 行为与 GetInt 一致：类型不匹配时返回 (0, errors.New("not int64"))。
// 当 JSON 中数字可能为小数时，更推荐使用 GetFloat64。
func (h H) GetInt64(key string) (int64, error) {
	out := h[key]
	outInt, ok := out.(int64)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not int64")
}

func (h H) GetFloat64(key string) (float64, error) {
	out := h[key]
	outInt, ok := out.(float64)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not float64")
}

func (h H) GetBool(key string) (bool, error) {
	out := h[key]
	outInt, ok := out.(bool)
	if ok {
		return outInt, nil
	}
	return false, errors.New("not bool")
}

func (h H) GetMap(key string) (H, error) {
	out := h[key]
	outInt, ok := out.(H)
	if ok {
		return outInt, nil
	}
	return nil, errors.New("not H")
}

func (h H) GetSlice(key string) ([]interface{}, error) {
	out := h[key]
	outInt, ok := out.([]interface{})
	if ok {
		return outInt, nil
	}
	return nil, errors.New("not []interface{}")
}

func (h H) Set(key string, value interface{}) {
	h[key] = value
}

func (h H) Keys() []string {
	keys := make([]string, 0, len(h))
	for key := range h {
		keys = append(keys, key)
	}

	return keys
}

func (h H) Values() []interface{} {
	values := make([]interface{}, 0, len(h))
	for _, value := range h {
		values = append(values, value)
	}

	return values
}

func (h H) Len() int {
	return len(h)
}

func (h H) DeepCopy() H {
	newMap := make(H)
	for k, v := range h {
		newMap[k] = DeepCopy(v)
	}

	return newMap
}
