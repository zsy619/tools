package xmap

// DeepCopy 递归深拷贝任意 interface{} 类型的 map 或切片。
//
// 仅对 map[string]interface{} 与 []interface{} 进行递归拷贝：
//   - 遇到 map[string]interface{} 时对每个 value 再调用本函数；
//   - 遇到 []interface{} 时对每个元素再调用本函数；
//   - 其它类型直接返回原值（浅拷贝）。
//
// 因此对于包含指针、结构体等非基础类型字段的 map，本函数不会真正
// 隔离引用；如需通用的深拷贝请改用 encoding/gob、JSON 序列化等方案。
//
// 入参为 nil 时返回 nil。
func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}
		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}
		return newSlice
	}
	return value
}

// Clone 浅拷贝 map：返回一个新的 map，键值与原 map 一致。
//
// 注意是浅拷贝：value 本身不会被深拷贝，修改原 map 不会影响新 map，
// 但修改 value 指向的对象会影响两者。
//
// 入参为 nil 时返回空 map（不是 nil），便于调用方继续使用。
func Clone[E comparable, V any](mapa map[E]V) map[E]V {
	size := len(mapa)
	ret := make(map[E]V, size)
	for k, v := range mapa {
		ret[k] = v
	}
	return ret
}

// AddAll 将 mapb 中「不存在于 mapa」的键值合并到 mapa 的副本并返回。
//
// 即 mapb 中已有的键不会覆盖 mapa 的同键值，保证 "先来先得"。
// 该函数不修改任何入参，总是返回一个新 map。
func AddAll[E comparable, V any](mapa, mapb map[E]V) map[E]V {
	ret := Clone(mapa)
	for k, v := range mapb {
		_, exist := ret[k]
		if !exist {
			ret[k] = v
		}
	}
	return ret
}

// Clear 删除 map 中的全部键值对，使其变为空 map。
//
// 与 maps.Clear（Go 1.21+）行为一致，但本函数兼容 Go 1.20 及更早版本。
//
// 注意：对 nil map 调用 Clear 是安全 no-op，不会触发 panic。
func Clear[E comparable, V any](mapa map[E]V) {
	if mapa == nil {
		return
	}
	for k := range mapa {
		delete(mapa, k)
	}
}
