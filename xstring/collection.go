package xstring

// Index 函数在字符串切片vs中查找字符串t，返回t在vs中的索引位置
// 如果未找到，则返回-1
// vs：待查找的字符串切片
// t：要查找的字符串
// 返回值：int类型，表示t在vs中的索引位置，若未找到则返回-1
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include 函数检查字符串 t 是否存在于字符串切片 vs 中，如果存在返回 true，否则返回 false
// 参数 vs 是待检查的字符串切片
// 参数 t 是待检查的字符串
// 返回值是布尔类型，表示字符串 t 是否存在于字符串切片 vs 中
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any 函数接受一个字符串切片和一个函数作为参数，判断是否存在某个元素满足函数条件
// 参数：
//
//	vs []string - 待判断的字符串切片
//	f func(string) bool - 判断函数，接受一个字符串参数，返回布尔值
//
// 返回值：
//
//	bool - 若存在满足条件的元素，则返回true，否则返回false
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All 函数接收一个字符串切片 vs 和一个函数 f 作为参数，
// 返回一个布尔值，用于判断切片 vs 中的所有元素是否都满足函数 f 的条件。
// 如果 vs 中所有元素都满足函数 f 的条件，则返回 true，否则返回 false。
// 函数 f 接收一个字符串参数，返回一个布尔值。
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter 函数接收一个字符串切片 vs 和一个函数 f，返回一个新的字符串切片 vsf
// vsf 包含所有满足函数 f 的字符串元素
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map 函数接受一个字符串切片vs和一个函数f作为参数，
// 将函数f应用到vs中的每个元素上，并将结果返回一个新的字符串切片。
// 参数vs：需要被处理的字符串切片。
// 参数f：一个接受字符串类型参数并返回字符串类型结果的函数。
// 返回值：一个新的字符串切片，包含函数f应用到vs中每个元素后的结果。
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
