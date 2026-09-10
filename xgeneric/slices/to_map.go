package slices

// ToMap 通过 keyFunc 与 valueFunc 将切片中的每个元素转换为 map 中的键值对。
// 参数：slice 为输入切片，keyFunc 用于生成键，valueFunc 用于生成值。
// 返回：长度最多为 len(slice) 的 map[K]V；当 keyFunc 产生重复键时后者覆盖前者。
func ToMap[T any, K comparable, V any](slice []T, keyFunc func(item T, index int, slice []T) K,
	valueFunc func(item T, index int, slice []T) V) map[K]V {
	dict := make(map[K]V, len(slice))
	for index, value := range slice {
		dict[keyFunc(value, index, slice)] = valueFunc(value, index, slice)
	}
	return dict
}
