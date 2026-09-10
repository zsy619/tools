package xstring

// Index 在字符串切片 vs 中线性查找字符串 t，并返回其下标。
//
// 参数说明：
//   - vs：待查找的字符串切片。允许为 nil 或空切片，均视为未命中。
//   - t：要查找的目标字符串。
//
// 返回值：
//   - 命中时返回 t 在 vs 中第一次出现的下标（>=0）。
//   - 未命中时返回 -1。
//
// 使用简单的线性扫描（O(n)）。比较使用 == 运算符，对字节敏感。
//
// 示例：
//
//	i := xstring.Index([]string{"a", "b", "c"}, "b") // i == 1
//	j := xstring.Index([]string{}, "x")               // j == -1
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include 判断字符串 t 是否出现在字符串切片 vs 中。
//
// 参数说明：
//   - vs：待检查的字符串切片。允许为 nil 或空切片，此时一定返回 false。
//   - t：要查找的字符串。
//
// 返回值：
//   - true 表示 t 在 vs 中至少出现一次；否则 false。
//
// 内部基于 Index 实现，对空切片与 nil 切片均安全。
//
// 示例：
//
//	xstring.Include([]string{"foo", "bar"}, "foo") // true
//	xstring.Include([]string{}, "foo")             // false
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any 判断字符串切片 vs 中是否存在至少一个元素满足函数 f 的判定。
//
// 参数说明：
//   - vs：待扫描的字符串切片。允许为 nil 或空切片，此时直接返回 false。
//   - f：判定函数，对每个元素调用一次；返回 true 即视为命中，函数立即被短路。
//
// 返回值：
//   - true 表示至少存在一个元素满足 f；false 表示没有任何元素满足，或切片为空。
//
// 短路求值：一旦命中即停止遍历；f 不会被多余调用。
//
// 示例：
//
//	xstring.Any([]string{"1", "22", "333"}, func(s string) bool { return len(s) > 2 }) // true
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All 判断字符串切片 vs 中的所有元素是否都满足函数 f 的判定。
//
// 参数说明：
//   - vs：待扫描的字符串切片。允许为 nil 或空切片，按"空集全称命题为真"的约定返回 true。
//   - f：判定函数，对每个元素调用一次；返回 false 即视为不满足，函数立即被短路。
//
// 返回值：
//   - true 表示全部元素满足 f，或切片为空；false 表示存在不满足 f 的元素。
//
// 短路求值：一旦遇到不满足的元素即停止遍历。
//
// 示例：
//
//	xstring.All([]string{"3", "33", "333"}, func(s string) bool { return len(s) > 0 }) // true
//	xstring.All([]string{}, func(s string) bool { return false })                      // true（空集）
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter 返回一个仅包含满足函数 f 的元素的新切片。
//
// 参数说明：
//   - vs：待过滤的字符串切片。允许为 nil 或空切片，结果均为长度为 0 的切片（非 nil）。
//   - f：判定函数，对每个元素调用一次，返回 true 表示保留该元素。
//
// 返回值：
//   - 一个新分配的 []string，按 vs 中的原始顺序保留通过 f 的元素。
//   - 即使没有任何元素满足，结果也是非 nil 的空切片（已通过 make([]string, 0) 初始化）。
//
// 不修改入参 vs；不会 panic。
//
// 示例：
//
//	evens := xstring.Filter([]string{"1", "2", "3", "4"}, func(s string) bool {
//	    n, _ := strconv.Atoi(s)
//	    return n%2 == 0
//	}) // evens == []string{"2", "4"}
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map 对字符串切片 vs 中的每个元素应用函数 f，返回一个等长的新切片。
//
// 参数说明：
//   - vs：待处理的字符串切片。允许为 nil 或空切片，结果分别为 nil 和长度为 0 的切片。
//   - f：转换函数，接收字符串返回字符串；返回值即为新切片对应位置的元素。
//
// 返回值：
//   - 与 vs 等长的新切片，每个元素是 f 对应位置元素的结果。
//
// 注意：与 Filter 不同，这里会预分配 len(vs) 长度的切片，因此即使 f 在某些位置上
// 返回空字符串，结果切片中也会保留这些位置（不会"压缩"）。
//
// 不修改入参 vs；不会 panic。
//
// 示例：
//
//	upper := xstring.Map([]string{"a", "b", "c"}, strings.ToUpper) // []string{"A","B","C"}
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
