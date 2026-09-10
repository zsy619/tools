package xgeneric

// Contains 判断切片中是否包含指定元素。
// 参数：collection 为待搜索的切片，element 为要查找的元素。
// 返回：找到则返回 true，否则返回 false。
func Contains[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}

	return false
}

// ContainsBy 判断切片中是否存在满足谓词函数的元素。
// 参数：collection 为待搜索的切片，predicate 为判断元素是否满足条件的函数。
// 返回：存在则返回 true，否则返回 false。
func ContainsBy[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}

	return false
}

// Every 判断 subset 中的所有元素是否全部存在于 collection 中。
// 参数：collection 为被包含的目标切片，subset 为候选元素切片。
// 返回：subset 中的每个元素都能在 collection 中找到时返回 true，否则返回 false。空 subset 总是返回 true。
func Every[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if !Contains(collection, elem) {
			return false
		}
	}

	return true
}

// Some 判断 subset 中是否至少有一个元素存在于 collection 中。
// 参数：collection 为被包含的目标切片，subset 为候选元素切片。
// 返回：subset 中存在任一元素可在 collection 中找到时返回 true，否则返回 false。空 subset 总是返回 false。
func Some[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if Contains(collection, elem) {
			return true
		}
	}

	return false
}

// Intersect 返回两个切片的交集（仅保留 list1 中出现过的元素，按 list2 顺序）。
// 参数：list1 与 list2 为输入切片，元素必须可比较。
// 返回：包含交集元素的新切片（list2 中顺序）。若两者均无重复元素，结果中不会包含重复项。
func Intersect[T comparable](list1 []T, list2 []T) []T {
	result := []T{}
	seen := map[T]struct{}{}

	for _, elem := range list1 {
		seen[elem] = struct{}{}
	}

	for _, elem := range list2 {
		if _, ok := seen[elem]; ok {
			result = append(result, elem)
		}
	}

	return result
}

// Difference 返回两个切片之间的差集。
// 参数：list1 与 list2 为输入切片，元素必须可比较。
// 返回：两个切片，第一个是存在于 list1 但不在 list2 中的元素（按 list1 顺序），第二个是存在于 list2 但不在 list1 中的元素（按 list2 顺序）。
func Difference[T comparable](list1 []T, list2 []T) ([]T, []T) {
	left := []T{}
	right := []T{}

	seenLeft := map[T]struct{}{}
	seenRight := map[T]struct{}{}

	for _, elem := range list1 {
		seenLeft[elem] = struct{}{}
	}

	for _, elem := range list2 {
		seenRight[elem] = struct{}{}
	}

	for _, elem := range list1 {
		if _, ok := seenRight[elem]; !ok {
			left = append(left, elem)
		}
	}

	for _, elem := range list2 {
		if _, ok := seenLeft[elem]; !ok {
			right = append(right, elem)
		}
	}

	return left, right
}

// Union 返回两个切片的所有不同元素的并集，相对顺序保持不变。
// 参数：list1 与 list2 为输入切片，元素必须可比较。
// 返回：去重后的并集切片，元素顺序为 list1 中先出现，再追加仅出现在 list2 中的元素。
func Union[T comparable](list1 []T, list2 []T) []T {
	result := []T{}

	seen := map[T]struct{}{}
	hasAdd := map[T]struct{}{}

	for _, e := range list1 {
		seen[e] = struct{}{}
	}

	for _, e := range list2 {
		seen[e] = struct{}{}
	}

	for _, e := range list1 {
		if _, ok := seen[e]; ok {
			result = append(result, e)
			hasAdd[e] = struct{}{}
		}
	}

	for _, e := range list2 {
		if _, ok := hasAdd[e]; ok {
			continue
		}
		if _, ok := seen[e]; ok {
			result = append(result, e)
		}
	}

	return result
}
