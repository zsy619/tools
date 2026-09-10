package slices

import (
	"bytes"
	"encoding/gob"
	"encoding/json"

	"github.com/zsy619/tools/xgeneric"
)

// Map 对切片中的每个元素应用函数 f，返回转换后的新切片。
// 参数：s 为输入切片，f 为转换函数。
// 返回：长度与 s 相同、由 f 转换后元素组成的新 []D。
func Map[S, D any](s []S, f func(S) D) []D {
	r := make([]D, 0, len(s))
	for _, e := range s {
		r = append(r, f(e))
	}
	return r
}

// Keys 返回包含 map 所有键的切片（顺序随机）。
// 参数：in 为输入 map。
// 返回：以任意顺序排列的所有键组成的 []K。
func Keys[K comparable, V any](in map[K]V) []K {
	result := make([]K, 0, len(in))

	for k := range in {
		result = append(result, k)
	}

	return result
}

// Values 返回包含 map 所有值的切片（顺序随机）。
// 参数：in 为输入 map。
// 返回：以任意顺序排列的所有值组成的 []V。
func Values[K comparable, V any](in map[K]V) []V {
	result := make([]V, 0, len(in))

	for _, v := range in {
		result = append(result, v)
	}

	return result
}

// PickBy 根据谓词函数过滤 map，仅保留满足条件的键值对。
// 参数：in 为输入 map，predicate 为过滤函数（对每个 (k,v) 调用，返回 true 时保留）。
// 返回：与 in 同类型的新 map。
func PickBy[K comparable, V any](in map[K]V, predicate func(K, V) bool) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if predicate(k, v) {
			r[k] = v
		}
	}
	return r
}

// PickByKeys 仅保留 keys 中存在的键对应的 map 元素。
// 参数：in 为输入 map，keys 为要保留的键列表。
// 返回：包含所有键在 keys 中出现过的元素的新 map。
func PickByKeys[K comparable, V any](in map[K]V, keys []K) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if xgeneric.Contains(keys, k) {
			r[k] = v
		}
	}
	return r
}

// PickByValues 仅保留值出现在 values 列表中的 map 元素。
// 参数：in 为输入 map，values 为要保留的值列表。
// 返回：包含所有值在 values 中出现过的元素的新 map。
func PickByValues[K comparable, V comparable](in map[K]V, values []V) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if xgeneric.Contains(values, v) {
			r[k] = v
		}
	}
	return r
}

// OmitBy 根据谓词函数过滤 map，移除满足条件的键值对（即 PickBy 的反义）。
// 参数：in 为输入 map，predicate 为过滤函数（返回 true 时移除）。
// 返回：与 in 同类型的新 map。
func OmitBy[K comparable, V any](in map[K]V, predicate func(K, V) bool) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if !predicate(k, v) {
			r[k] = v
		}
	}
	return r
}

// OmitByKeys 移除键出现在 keys 列表中的 map 元素（即 PickByKeys 的反义）。
// 参数：in 为输入 map，keys 为要移除的键列表。
// 返回：不包含 keys 中键的新 map。
func OmitByKeys[K comparable, V any](in map[K]V, keys []K) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if !xgeneric.Contains(keys, k) {
			r[k] = v
		}
	}
	return r
}

// OmitByValues 移除值出现在 values 列表中的 map 元素（即 PickByValues 的反义）。
// 参数：in 为输入 map，values 为要移除的值列表。
// 返回：不包含 values 中值的新 map。
func OmitByValues[K comparable, V comparable](in map[K]V, values []V) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if !xgeneric.Contains(values, v) {
			r[k] = v
		}
	}
	return r
}

// Entries 将 map 转换为由键值对组成的切片。
// 参数：in 为输入 map。
// 返回：包含所有 (key, value) 对的 []xgeneric.Entry[K, V]（顺序随机）。
func Entries[K comparable, V any](in map[K]V) []xgeneric.Entry[K, V] {
	entries := make([]xgeneric.Entry[K, V], 0, len(in))

	for k, v := range in {
		entries = append(entries, xgeneric.Entry[K, V]{
			Key:   k,
			Value: v,
		})
	}

	return entries
}

// FromEntries 将 (key, value) 对切片转换为 map。
// 参数：entries 为输入的 []xgeneric.Entry[K, V]。
// 返回：由 entries 生成的 map；当出现重复键时后者覆盖前者。
func FromEntries[K comparable, V any](entries []xgeneric.Entry[K, V]) map[K]V {
	out := map[K]V{}

	for _, v := range entries {
		out[v.Key] = v.Value
	}

	return out
}

// Invert 创建一个键值互换的新 map（即 (k,v) → (v,k)）。
// 如果原 map 包含重复的值，后出现的键会覆盖之前的键。
// 参数：in 为输入 map（K、V 均必须可比较）。
// 返回：键为 V、值为 K 的 map[V]K。
func Invert[K comparable, V comparable](in map[K]V) map[V]K {
	out := map[V]K{}

	for k, v := range in {
		out[v] = k
	}

	return out
}

// Assign 从左到右合并多个 map，后出现的键值会覆盖先前的。
// 参数：maps 为可变数量的 map[K]V。
// 返回：合并后的新 map；任一参数为 nil 时该 map 不参与合并。
func Assign[K comparable, V any](maps ...map[K]V) map[K]V {
	out := map[K]V{}

	for _, m := range maps {
		for k, v := range m {
			out[k] = v
		}
	}

	return out
}

// MapKeys 通过 iteratee 函数重新生成每个键，得到键类型为 R 的新 map（值不变）。
// 参数：in 为输入 map，iteratee 为键转换函数（接收值与原键，返回新键）。
// 返回：键类型为 R、值不变的 map[R]V；若 iteratee 返回重复键则后者覆盖前者。
func MapKeys[K comparable, V any, R comparable](in map[K]V, iteratee func(V, K) R) map[R]V {
	result := map[R]V{}

	for k, v := range in {
		result[iteratee(v, k)] = v
	}

	return result
}

// MapValues 通过 iteratee 函数重新生成每个值，得到值类型为 R 的新 map（键不变）。
// 参数：in 为输入 map，iteratee 为值转换函数（接收值与键，返回新值）。
// 返回：键不变、值类型为 R 的 map[K]R。
func MapValues[K comparable, V any, R any](in map[K]V, iteratee func(V, K) R) map[K]R {
	result := map[K]R{}

	for k, v := range in {
		result[k] = iteratee(v, k)
	}

	return result
}

// DeepCopy 创建 map 的浅拷贝（仅复制键值引用，不递归深拷贝）。
// 返回：包含相同键值的新 map[K]V；修改新 map 不会影响原 map，但键值本身的引用仍共享。
func DeepCopy[K comparable, V any](value map[K]V) map[K]V {
	newMap := make(map[K]V)
	for k, v := range value {
		newMap[k] = v
	}

	return newMap
}

// DeepCopyByGob 通过 gob 编码/解码对 map 进行深拷贝（要求 K、V 实现了 gob 编解码）。
// 返回：与 value 等值的新 map[K]V 及 nil 错误；编码或解码失败时返回 nil 与 error。
func DeepCopyByGob[K comparable, V any](value map[K]V) (dst map[K]V, err error) {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(value); err != nil {
		return nil, err
	}

	err = gob.NewDecoder(&buffer).Decode(dst)
	return
}

// DeepCopyByJson 通过 JSON 序列化/反序列化对 map 进行深拷贝（要求 K、V 可被 JSON 编解码）。
// 返回：与 value 等值的新 map[K]V 及 nil 错误；序列化或反序列化失败时返回 nil 与 error。
func DeepCopyByJson[K comparable, V any](value map[K]V) (dst map[K]V, err error) {
	b, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, dst)
	return dst, err
}
