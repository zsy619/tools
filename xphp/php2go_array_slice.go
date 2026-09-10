package xphp

import "reflect"

// ArrayUnique 实现 PHP array_unique()：去除字符串切片中的重复元素，保持首次出现的顺序。
func ArrayUnique(arr []string) []string {
	size := len(arr)
	result := make([]string, 0, size)
	temp := map[string]struct{}{}
	for i := 0; i < size; i++ {
		if _, ok := temp[arr[i]]; !ok {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}
	return result
}

// ArraySlice 实现 PHP array_slice()：从 s 的 offset 起取 length 个元素。
// offset 越界时 panic；length 越界时取到末尾。
func ArraySlice(s []interface{}, offset, length int) []interface{} {
	if offset > len(s) {
		panic("offset: the offset is less than the length of s")
	}
	end := offset + length
	if end < len(s) {
		return s[offset:end]
	}
	return s[offset:]
}

// ArraySliceString 是 ArraySlice 针对字符串切片的便利版本。
func ArraySliceString(s []string, offset, length int) []string {
	if offset > len(s) {
		panic("offset: the offset is less than the length of s")
	}
	end := offset + length
	if end < len(s) {
		return s[offset:end]
	}
	return s[offset:]
}

// ArrayDiff 实现 PHP array_diff()：返回在 array1 但不在任意 arrayOthers 中的字符串元素。
// 输出顺序由 map 遍历顺序决定，不保证稳定。
func ArrayDiff(array1 []string, arrayOthers ...[]string) []string {
	c := make(map[string]bool)
	for i := 0; i < len(array1); i++ {
		if _, hasKey := c[array1[i]]; hasKey {
			c[array1[i]] = true
		} else {
			c[array1[i]] = false
		}
	}
	for i := 0; i < len(arrayOthers); i++ {
		for j := 0; j < len(arrayOthers[i]); j++ {
			if _, hasKey := c[arrayOthers[i][j]]; hasKey {
				c[arrayOthers[i][j]] = true
			} else {
				c[arrayOthers[i][j]] = false
			}
		}
	}
	result := make([]string, 0)
	for k, v := range c {
		if !v {
			result = append(result, k)
		}
	}
	return result
}

// ArrayIntersect 实现 PHP array_intersect()：返回同时存在于 array1 与任意 arrayOthers 的字符串元素。
func ArrayIntersect(array1 []string, arrayOthers ...[]string) []string {
	c := make(map[string]bool)
	for i := 0; i < len(array1); i++ {
		if _, hasKey := c[array1[i]]; hasKey {
			c[array1[i]] = true
		} else {
			c[array1[i]] = false
		}
	}
	for i := 0; i < len(arrayOthers); i++ {
		for j := 0; j < len(arrayOthers[i]); j++ {
			if _, hasKey := c[arrayOthers[i][j]]; hasKey {
				c[arrayOthers[i][j]] = true
			} else {
				c[arrayOthers[i][j]] = false
			}
		}
	}
	result := make([]string, 0)
	for k, v := range c {
		if v {
			result = append(result, k)
		}
	}
	return result
}

// ArraySearch 实现 PHP array_search()：在 hystack 中查找 needle。
// 仅支持切片（reflect.Slice），使用 reflect.DeepEqual 比较元素。
// 命中返回首个匹配下标，未命中返回 -1。
func ArraySearch(needle interface{}, hystack interface{}) (index int) {
	index = -1

	switch reflect.TypeOf(hystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(hystack)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) {
				index = i
				return
			}
		}
	}
	return
}
