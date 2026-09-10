package slices

import "sort"

// SortStable 使用 condition 比较函数对切片进行原地稳定排序（相等元素保持原有顺序）。
// 参数：slice 为输入切片（会被就地修改），condition 为比较函数（返回 true 表示 item1 应排在 item2 之前）。
// 返回：已排序的原切片本身。
func SortStable[T any](slice []T, condition func(item1, item2 T) bool) []T {
	sort.SliceStable(slice, func(i, j int) bool {
		return condition(slice[i], slice[j])
	})
	return slice
}
