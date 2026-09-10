package xphp

import (
	"math"
	"reflect"
	"sort"
)

// ArrayUniqueItf 去除切片/数组中的重复元素并返回与原类型相同的去重结果。
//
// 行为：
//   - array 为 nil：返回 nil；
//   - array 不是切片/数组类型：原样返回 array；
//   - array 长度为 0：原样返回 array；
//   - 其它情况：根据首元素类型返回相应类型的切片（bool/int*/uint*/float*/complex*/string）；
//     不支持的类型退化为 []interface{}。
//
// 提示：可通过类型断言将返回值转换为原输入的具体切片类型。
func ArrayUniqueItf(array interface{}) interface{} {
	type empty struct{}

	if array == nil {
		return nil
	}

	// if it's not a slice then return the original input
	kind := reflect.TypeOf(array).Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return array
	}

	s := reflect.ValueOf(array)
	if s.Len() == 0 {
		return array
	}

	set := make(map[interface{}]empty)
	for i := 0; i < s.Len(); i++ {
		set[s.Index(i).Interface()] = empty{}
	}

	switch s.Index(0).Kind() {
	case reflect.Bool:
		result := make([]bool, 0, s.Len())
		for k := range set {
			result = append(result, k.(bool))
		}
		return result
	case reflect.Int:
		result := make([]int, 0, s.Len())
		for k := range set {
			result = append(result, k.(int))
		}
		return result
	case reflect.Int8:
		result := make([]int8, 0, s.Len())
		for k := range set {
			result = append(result, k.(int8))
		}
		return result
	case reflect.Int16:
		result := make([]int16, 0, s.Len())
		for k := range set {
			result = append(result, k.(int16))
		}
		return result
	case reflect.Int32:
		result := make([]int32, 0, s.Len())
		for k := range set {
			result = append(result, k.(int32))
		}
		return result
	case reflect.Int64:
		result := make([]int64, 0, s.Len())
		for k := range set {
			result = append(result, k.(int64))
		}
		return result
	case reflect.Uint:
		result := make([]uint, 0, s.Len())
		for k := range set {
			result = append(result, k.(uint))
		}
		return result
	case reflect.Uint8:
		result := make([]uint8, 0, s.Len())
		for k := range set {
			result = append(result, k.(uint8))
		}
		return result
	case reflect.Uint16:
		result := make([]uint16, 0, s.Len())
		for k := range set {
			result = append(result, k.(uint16))
		}
		return result
	case reflect.Uint32:
		result := make([]uint32, 0, s.Len())
		for k := range set {
			result = append(result, k.(uint32))
		}
		return result
	case reflect.Uint64:
		result := make([]uint64, 0, s.Len())
		for k := range set {
			result = append(result, k.(uint64))
		}
		return result
	case reflect.Float32:
		result := make([]float32, 0, s.Len())
		for k := range set {
			result = append(result, k.(float32))
		}
		return result
	case reflect.Float64:
		result := make([]float64, 0, s.Len())
		for k := range set {
			result = append(result, k.(float64))
		}
		return result
	case reflect.Complex64:
		result := make([]complex64, 0, s.Len())
		for k := range set {
			result = append(result, k.(complex64))
		}
		return result
	case reflect.Complex128:
		result := make([]complex128, 0, s.Len())
		for k := range set {
			result = append(result, k.(complex128))
		}
		return result
	case reflect.String:
		result := make([]string, 0, s.Len())
		for k := range set {
			result = append(result, k.(string))
		}
		return result
	default:
		result := make([]interface{}, 0, s.Len())
		for k := range set {
			result = append(result, k)
		}
		return result
	}
}

// InArrayItf 判断 needle 是否存在于 haystack（切片、数组或 map 的值）中。
// 使用 reflect.DeepEqual 比较元素；map 情况下比较的是 value 而非 key。
// haystack 为 nil 或其它不支持的类型时返回 false。
func InArrayItf(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	}

	return false
}

// ArrayChunkItf 按 size 将 array 切分为若干子切片。
// size < 1 时返回 nil；最后一个子切片可能不足 size 个元素。
func ArrayChunkItf(array []interface{}, size int) [][]interface{} {
	if size < 1 {
		return nil
	}
	length := len(array)
	chunkNum := int(math.Ceil(float64(length) / float64(size)))
	var chunks [][]interface{}
	for i, end := 0, 0; chunkNum > 0; chunkNum-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		chunks = append(chunks, array[i*size:end])
		i++
	}
	return chunks
}

// ArrayColumnItf 从 input 中提取每个 map 在 columnKey 列上的值，返回扁平 []interface{}。
// 元素中若不存在 columnKey 则跳过；不会触发 panic。
func ArrayColumnItf(input []map[string]interface{}, columnKey string) []interface{} {
	columns := make([]interface{}, 0, len(input))
	for _, val := range input {
		if v, ok := val[columnKey]; ok {
			columns = append(columns, v)
		}
	}
	return columns
}

// ArrayCombineItf 使用 keys 作为键、values 作为值构造 map[interface{}]interface{}。
// 当 keys 与 values 长度不一致时返回 nil。
func ArrayCombineItf(keys, values []interface{}) map[interface{}]interface{} {
	if len(keys) != len(values) {
		return nil
	}
	m := make(map[interface{}]interface{}, len(keys))
	for i, v := range keys {
		m[v] = values[i]
	}
	return m
}

// ArrayDiffItf 返回在 array1 但不在 array2 中的元素（保持 array1 顺序）。
func ArrayDiffItf(array1, array2 []interface{}) []interface{} {
	var res []interface{}
	for _, v := range array1 {
		if !InArray(v, array2) {
			res = append(res, v)
		}
	}
	return res
}

// ArrayIntersectItf 返回同时存在于 array1 与 array2 中的元素（保持 array1 顺序）。
func ArrayIntersectItf(array1, array2 []interface{}) []interface{} {
	var res []interface{}
	for _, v := range array1 {
		if InArray(v, array2) {
			res = append(res, v)
		}
	}
	return res
}

// ArrayFlipItf 交换切片/数组的下标与元素值，或交换 map 的 key 与 value。
// 输入为 nil、空切片或非切片/数组/map 类型时返回 nil。
func ArrayFlipItf(input interface{}) interface{} {
	if input == nil {
		return nil
	}
	val := reflect.ValueOf(input)
	if val.Len() == 0 {
		return nil
	}
	res := make(map[interface{}]interface{}, val.Len())
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			res[val.Index(i).Interface()] = i
		}
		return res
	case reflect.Map:
		for _, k := range val.MapKeys() {
			res[val.MapIndex(k).Interface()] = k.Interface()
		}
		return res
	}
	return nil
}

// ArrayKeysItf 返回切片/数组的所有下标（[]int）或 map 的所有 key（排序后的 []string）。
// 输入为 nil 或长度为 0 时返回 nil。
func ArrayKeysItf(input interface{}) interface{} {
	if input == nil {
		return nil
	}
	val := reflect.ValueOf(input)
	if val.Len() == 0 {
		return nil
	}
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		var res []int
		for i := 0; i < val.Len(); i++ {
			res = append(res, i)
		}
		return res
	case reflect.Map:
		var res []string
		for _, k := range val.MapKeys() {
			res = append(res, k.String())
		}
		sort.SliceStable(res, func(i, j int) bool {
			return res[i] < res[j]
		})
		return res
	}
	return nil
}

// ArrayKeyExistsItf 是 KeyExistsItf 的别名，判断给定 key 是否存在于 map 中。
func ArrayKeyExistsItf(k interface{}, m map[interface{}]interface{}) bool {
	return KeyExistsItf(k, m)
}

// KeyExistsItf 判断给定 key 是否存在于 map 中。
func KeyExistsItf(k interface{}, m map[interface{}]interface{}) bool {
	_, ok := m[k]
	return ok
}

// CountItf 返回数组/切片/map 中的元素数量；nil 返回 0。
func CountItf(v interface{}) int {
	if v == nil {
		return 0
	}
	return reflect.ValueOf(v).Len()
}

// ArrayFilterItf 使用 callback 对切片/数组或 map 进行过滤。
// 输入为 nil/空或非支持类型时返回 nil；callback 为 nil 时使用默认的“非 nil 即保留”规则。
func ArrayFilterItf(input interface{}, callback func(interface{}) bool) interface{} {
	if input == nil {
		return nil
	}
	val := reflect.ValueOf(input)
	if val.Len() == 0 {
		return nil
	}
	if callback == nil {
		callback = func(v interface{}) bool {
			return v != nil
		}
	}
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		var res []interface{}
		for i := 0; i < val.Len(); i++ {
			v := val.Index(i).Interface()
			if callback(v) {
				res = append(res, v)
			}
		}
		return res
	case reflect.Map:
		res := make(map[interface{}]interface{})
		for _, k := range val.MapKeys() {
			v := val.MapIndex(k).Interface()
			if callback(v) {
				res[k.Interface()] = v
			}
		}
		return res
	}

	return input
}

// ArrayPadItf 使用 value 将 array 填充到指定长度 size。
//   - size == 0、size > 0 但小于 len(array)、size < 0 但大于 -len(array)：原样返回；
//   - size > 0：在尾部填充；
//   - size < 0：在头部填充。
func ArrayPadItf(array []interface{}, size int, value interface{}) []interface{} {
	if size == 0 || (size > 0 && size < len(array)) || (size < 0 && size > -len(array)) {
		return array
	}
	n := size
	if size < 0 {
		n = -size
	}
	n -= len(array)
	tmp := make([]interface{}, n)
	for i := 0; i < n; i++ {
		tmp[i] = value
	}
	if size > 0 {
		return append(array, tmp...)
	}
	return append(tmp, array...)
}

// ArrayPopItf 弹出切片末尾元素并从切片中移除；s 为 nil 或空时返回 nil。
func ArrayPopItf(s *[]interface{}) interface{} {
	if s == nil || len(*s) == 0 {
		return nil
	}
	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]
	return e
}

// ArrayPushItf 将一个或多个元素追加到切片末尾，返回追加后切片的新长度；s 为 nil 时返回 0。
func ArrayPushItf(s *[]interface{}, elements ...interface{}) int {
	if s == nil {
		return 0
	}
	*s = append(*s, elements...)
	return len(*s)
}

// ArrayShiftItf 移除并返回切片首元素；s 为 nil 或空时返回 nil。
func ArrayShiftItf(s *[]interface{}) interface{} {
	if s == nil || len(*s) == 0 {
		return nil
	}
	f := (*s)[0]
	*s = (*s)[1:]
	return f
}

// ArrayUnshiftItf 将一个或多个元素插入到切片头部，返回插入后切片的新长度；s 为 nil 时返回 0。
func ArrayUnshiftItf(s *[]interface{}, elements ...interface{}) int {
	if s == nil {
		return 0
	}
	*s = append(elements, *s...)
	return len(*s)
}

// ArrayReverseItf 反转切片并返回（原地修改）。
func ArrayReverseItf(s []interface{}) []interface{} {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// ArraySliceItf 按 offset 与 length 从 array 中截取子切片。
// offset 大于 len(array) 时返回 nil；length 越界时取到末尾。
func ArraySliceItf(array []interface{}, offset, length uint) []interface{} {
	if offset > uint(len(array)) {
		return nil
	}
	end := offset + length
	if end < uint(len(array)) {
		return array[offset:end]
	}
	return array[offset:]
}

// ArraySumItf 对切片/数组元素求和，返回与元素类型相对应的求和值。
// 整型返回 int64，无符号整型返回 uint64，浮点返回 float64，字符串返回拼接结果；
// 不支持的类型返回 nil。
func ArraySumItf(array interface{}) interface{} {
	if array == nil {
		return nil
	}
	kind := reflect.TypeOf(array).Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil
	}

	s := reflect.ValueOf(array)
	if s.Len() == 0 {
		return 0
	}

	switch s.Index(0).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var sum int64
		for i := 0; i < s.Len(); i++ {
			sum += s.Index(i).Int()
		}
		return sum
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var sum uint64
		for i := 0; i < s.Len(); i++ {
			sum += s.Index(i).Uint()
		}
		return sum
	case reflect.Float32, reflect.Float64:
		var sum float64
		for i := 0; i < s.Len(); i++ {
			sum += s.Index(i).Float()
		}
		return sum
	case reflect.String:
		var sum string
		for i := 0; i < s.Len(); i++ {
			sum += s.Index(i).String()
		}
		return sum
	}

	return nil
}

// SortItf 对切片/数组按从小到大排序，返回按其元素类型排序后的新切片。
// array 为 nil、非切片/数组类型或长度为 0 时返回原值；不支持的类型原样返回。
func SortItf(array interface{}) interface{} {
	if array == nil {
		return nil
	}
	kind := reflect.TypeOf(array).Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil
	}

	s := reflect.ValueOf(array)
	if s.Len() == 0 {
		return array
	}

	switch s.Index(0).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		res := make([]int64, s.Len())
		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).Int()
		}
		sort.Slice(res, func(i int, j int) bool {
			return res[i] < res[j]
		})
		return res
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		res := make([]uint64, s.Len())
		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).Uint()
		}
		sort.Slice(res, func(i int, j int) bool {
			return res[i] < res[j]
		})
		return res
	case reflect.Float32, reflect.Float64:
		res := make([]float64, s.Len())
		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).Float()
		}
		sort.Slice(res, func(i int, j int) bool {
			return res[i] < res[j]
		})
		return res
	case reflect.String:
		res := make([]string, s.Len())
		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).String()
		}
		sort.Slice(res, func(i int, j int) bool {
			return res[i] < res[j]
		})
		return res
	}

	return array
}

// RsortItf 对切片/数组按从大到小排序，返回按其元素类型排序后的新切片。
// array 为 nil、非切片/数组类型或长度为 0 时返回原值；不支持的类型原样返回。
func RsortItf(array interface{}) interface{} {
	if array == nil {
		return nil
	}
	kind := reflect.TypeOf(array).Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil
	}

	s := reflect.ValueOf(array)
	if s.Len() == 0 {
		return array
	}

	switch s.Index(0).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		res := make([]int64, s.Len())
		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).Int()
		}
		sort.Slice(res, func(i int, j int) bool {
			return res[i] > res[j]
		})
		return res
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		res := make([]uint64, s.Len())
		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).Uint()
		}
		sort.Slice(res, func(i int, j int) bool {
			return res[i] > res[j]
		})
		return res
	case reflect.Float32, reflect.Float64:
		res := make([]float64, s.Len())
		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).Float()
		}
		sort.Slice(res, func(i int, j int) bool {
			return res[i] > res[j]
		})
		return res
	case reflect.String:
		res := make([]string, s.Len())
		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).String()
		}
		sort.Slice(res, func(i int, j int) bool {
			return res[i] > res[j]
		})
		return res
	}

	return array
}
