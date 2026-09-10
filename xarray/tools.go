package xarray

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/zsy619/tools"
	"github.com/zsy619/tools/xgeneric"
	"github.com/zsy619/tools/xgeneric/maths"
	"github.com/zsy619/tools/xmap"
)

const (
	// SHUFFLE_THRESHOLD 打乱（Shuffle）相关操作所用的阈值常量，保留用于后续逻辑扩展。
	SHUFFLE_THRESHOLD = 5
)

// Sort 根据自定义 cmp 比较函数对切片进行原地排序。
// 参数：data 为待排序切片，cmp 为用户自定义的比较函数，约定返回负数表示 e1<e2、0 表示相等、正数表示 e1>e2。
// 返回值：无，原切片修改。空切片或 nil 切片调用安全（不执行任何操作）。
// 副作用：直接修改入参 data 中的元素顺序。
//
// 示例代码：
//
//	type Student struct {
//		Name string
//	}
//
//	students := []Student{{"xml"}, {"matthew"}, {"matt"}, {"xiemalin"}}
//	Sort(students, func(e1, e2 Student) int {
//		return strings.Compare(e1.Name, e2.Name)
//	})
func Sort[E any](data []E, cmp tools.CMP[E]) {
	sortobject := sortable[E]{data: data, cmp: cmp}
	sort.Sort(sortobject)
}

// SortOrdered 使用类型自身定义的天然顺序（xgeneric.Ordered）对切片进行原地排序。
// 参数：data 为待排序切片，asc 为 true 时按升序排列，false 时按降序排列。
// 返回值：无，原切片修改。空切片或 nil 切片调用安全（不执行任何操作）。
// 副作用：直接修改入参 data 中的元素顺序。
// 更多细节请参考 xgeneric.Ordered。
//
// 示例代码：
//
//	strArray := []string{"xml", "matthew", "matt", "xiemalin"}
//	SortOrdered(strArray, true)
func SortOrdered[E xgeneric.Ordered](data []E, asc bool) {
	sortobject := sortable[E]{data: data, cmp: func(e1, e2 E) int {
		ord := -1
		if !asc {
			ord *= -1
		}
		return ord * CompareTo(e1, e2)
	}}
	sort.Sort(sortobject)
}

// sortable 实现 sort.Interface 的泛型包装结构体，用于将任意切片与自定义比较函数适配到标准库排序接口。
type sortable[E any] struct {
	data []E          // 待排序的元素切片
	cmp  tools.CMP[E] // 用户提供的比较函数
}

// Len 返回切片中元素的数量，实现 sort.Interface。
func (s sortable[E]) Len() int { return len(s.data) }

// Swap 交换切片中索引 i 与 j 处的元素，实现 sort.Interface。
func (s sortable[E]) Swap(i, j int) { s.data[i], s.data[j] = s.data[j], s.data[i] }

// Less 当 data[i] 在自定义比较中“不小于”data[j] 时返回 true，实现 sort.Interface。
// 返回 s.cmp(...) >= 0 是为了兼容 SortOrdered 中通过正负号翻转实现的升降序逻辑。
func (s sortable[E]) Less(i, j int) bool {
	return s.cmp(s.data[i], s.data[j]) >= 0
}

// Shuffle 使用默认的随机源（基于切片长度作为种子）对切片进行原地随机打乱。
// 参数：data 为待打乱的切片。
// 返回值：无，原切片修改。空切片或 nil 切片调用安全（不执行任何操作）。
// 注意：以 len(data) 作为种子意味着相同长度的输入可能产生相同的打乱结果，因此不适合安全敏感场景。
func Shuffle[E any](data []E) {
	r := rand.New(rand.NewSource(int64(len(data))))
	ShuffleRandom(data, r)
}

// ShuffleRandom 使用调用方指定的随机源对切片进行原地随机打乱。
// 参数：data 为待打乱的切片，r 为用户提供的 *rand.Rand 随机源。
// 返回值：无，原切片修改。空切片或 nil 切片调用安全（不执行任何操作）。
func ShuffleRandom[E any](data []E, r *rand.Rand) {
	size := len(data)
	for i := 0; i < size; i++ {
		j := r.Intn(size)
		data[i], data[j] = data[j], data[i]
	}
}

// ShuffleRead 以随机顺序遍历切片中的每一个元素，并将元素交给 reader 回调处理。
// 参数：data 为待遍历的切片，reader 为针对每个元素执行的回调函数。
// 返回值：无。注意：data 本身不会被修改，但 reader 内可能对元素自身产生副作用。
// 空切片或 nil 切片调用安全（reader 不会被调用）。
func ShuffleRead[E any](data []E, reader func(e E)) {
	size := len(data)
	eleOrder := rand.Perm(size)
	for _, k := range eleOrder {
		reader(data[k])
	}
}

// Reverse 将切片中元素的顺序就地翻转。
// 参数：data 为待翻转的切片。
// 返回值：无，原切片修改。空切片或 nil 切片调用安全（不执行任何操作）。
func Reverse[E any](data []E) {
	size := len(data)
	mid := size >> 1
	j := size - 1
	for i := 0; i < mid; i++ {
		data[i], data[j] = data[j], data[i]
		j--
	}
}

// BinarySearch 使用二分查找算法在已排序切片中查找指定值。
// 参数：data 必须已按 cmp 所定义的顺序排好序，key 为目标值，cmp 为比较函数。
// 返回值：找到时返回 key 所在的索引；未找到时返回 -(low+1)，其中 low 为假设插入位置的索引（调用方可据此换算）。
// 注意：目标切片必须已排序；空切片或 nil 切片会返回 -1。
func BinarySearch[E any](data []E, key E, cmp tools.CMP[E]) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) >> 1
		midVal := data[mid]

		r := cmp(midVal, key)

		if r < 0 {
			low = mid + 1
		} else if r > 0 {
			high = mid - 1
		} else {
			return mid // key found
		}
	}
	return -(low + 1) // key not found.
}

// BinarySearchOrdered 使用类型天然顺序（xgeneric.Ordered）在已排序切片中二分查找指定值。
// 参数：data 必须已升序排好序，key 为目标值。
// 返回值：找到时返回 key 所在的索引；未找到时返回 -(low+1)。
// 内部直接复用 BinarySearch 并传入 CompareTo 作为比较函数。
func BinarySearchOrdered[E xgeneric.Ordered](data []E, key E) int {
	return BinarySearch(data, key, CompareTo[E])
}

// Contains 判断切片中是否包含与 key 相等的元素。
// 参数：data 为待搜索切片，key 为目标元素，equal 为自定义相等判定函数。
// 返回值：当切片中存在与 key 相等的元素时返回 true；切片为空/nil 或不存在匹配时返回 false。
func Contains[E any](data []E, key E, equal tools.EQL[E]) bool {
	size := len(data)
	if size == 0 {
		return false
	}
	for i := 0; i < size; i++ {
		if equal(data[i], key) {
			return true
		}
	}

	return false
}

// ContainsOrdered 使用 == 运算符判断已排序切片是否包含指定元素（适用于 xgeneric.Ordered 类型）。
// 参数：data 为待搜索切片，key 为目标元素。
// 返回值：找到时返回 true；切片为空/nil 或未匹配时返回 false。
func ContainsOrdered[E xgeneric.Ordered](data []E, key E) bool {
	return Contains(data, key, Equals[E])
}

// ContainsAny 判断切片 data 是否至少包含 other 中的一个元素。
// 参数：data 为待搜索切片，other 为候选元素切片，equal 为自定义相等判定函数。
// 返回值：当 other 中任意一个元素出现在 data 中时返回 true；当 other 为空或全部不存在时返回 false。
func ContainsAny[E any](data, other []E, equal tools.EQL[E]) bool {
	size := len(other)
	if size == 0 {
		return false
	}
	for i := 0; i < size; i++ {
		if Contains(data, other[i], equal) {
			return true
		}
	}

	return false
}

// ContainsAnyOrdered 使用 == 运算符判断切片 data 是否至少包含 other 中的一个元素（适用于 xgeneric.Ordered 类型）。
// 参数：data 为待搜索切片，other 为候选元素切片。
// 返回值：当 other 中任意一个元素出现在 data 中时返回 true；other 为空或全部不存在时返回 false。
func ContainsAnyOrdered[E xgeneric.Ordered](data, other []E) bool {
	return ContainsAny(data, other, Equals[E])
}

// Remove 从切片中移除第一个与 key 相等的元素并返回新切片。
// 参数：data 为待操作切片，key 为目标元素，equal 为自定义相等判定函数。
// 返回值：移除后的新切片与一个 bool 值（true 表示找到并已移除，false 表示未找到或切片为空）。
// 副作用：原切片不会被修改，总是返回一个新切片；返回的切片可能在原数组上复用底层存储。
func Remove[E any](data []E, key E, equal tools.EQL[E]) ([]E, bool) {
	return removeContional(data, key, equal, false)
}

// RemoveOrdered 使用 == 运算符从已排序切片中移除第一个与 key 相等的元素（适用于 xgeneric.Ordered 类型）。
// 参数：data 为待操作切片，key 为目标元素。
// 返回值：移除后的新切片与一个 bool 值（true 表示已移除，false 表示未找到或切片为空）。
func RemoveOrdered[E xgeneric.Ordered](data []E, key E) ([]E, bool) {
	return removeContional(data, key, Equals[E], false)
}

// RemoveAll 从切片中移除所有与 key 相等的元素并返回新切片。
// 参数：data 为待操作切片，key 为目标元素，equal 为自定义相等判定函数。
// 返回值：移除所有匹配元素后的新切片与一个 bool 值。
// 注意：当前实现中返回值总是 true（除非切片为空返回 false），并不严格反映是否实际移除了元素。
func RemoveAll[E any](data []E, key E, equal tools.EQL[E]) ([]E, bool) {
	return removeContional(data, key, equal, true)
}

// RemoveAllOrdered 使用 == 运算符从已排序切片中移除所有与 key 相等的元素（适用于 xgeneric.Ordered 类型）。
// 参数：data 为待操作切片，key 为目标元素。
// 返回值：移除所有匹配元素后的新切片与一个 bool 值。
func RemoveAllOrdered[E xgeneric.Ordered](data []E, key E) ([]E, bool) {
	return removeContional(data, key, Equals[E], true)
}

// removeContional 是 Remove 与 RemoveAll 共用的内部实现，按 all 控制是移除首个匹配项还是移除全部匹配项。
// 参数：data 为源切片，key 为目标元素，equal 为相等判定函数，all 为 true 时移除所有匹配项、为 false 时仅移除第一个。
// 返回值：处理后的新切片与 bool 值（true 表示至少移除过或切片非空；空切片返回 data 原值与 false）。
func removeContional[E any](data []E, key E, equal tools.EQL[E], all bool) ([]E, bool) {
	size := len(data)
	if size == 0 {
		return data, false
	}
	ret := make([]E, 0, size)
	for i := 0; i < size; i++ {
		if equal(data[i], key) {
			if !all {
				ret = append(ret, data[i+1:]...)
				return ret, true
			}
		} else {
			ret = append(ret, data[i])
		}
	}

	return ret, true
}

// RemoveIndex 按索引就地删除切片中的元素。
// 参数：data 为待操作切片，i 为待删除元素的索引。
// 返回值：删除指定索引后切片长度减 1 的新切片。
// 注意：当索引越界（i<0 或 i>=len(data)）时直接返回原切片，不做修改也不会 panic。
// 副作用：会就地修改入参底层数组中元素的位置与顺序。
func RemoveIndex[E any](data []E, i int) []E {
	size := len(data)
	if i >= size || i < 0 { // out of index, do nothing
		return data
	}

	if i < size-1 {
		copy(data[i:], data[i+1:])
	}
	return data[:len(data)-1]
}

// Min 在切片中查找最小元素及其首次出现的下标。
// 参数：data 为待搜索切片，cmp 为自定义比较函数（语义同 tools.CMP，返回值小于 0 表示 e1<e2）。
// 返回值：最小元素与该元素在切片中的位置（若存在多个最小元素，返回第一个的下标）。
// 注意：当前实现没有显式处理 size==0 的情况，调用空切片会导致下标越界 panic；size==1 时直接返回 (data[0], 0)。
func Min[E any](data []E, cmp tools.CMP[E]) (E, int) {
	size := len(data)
	if size == 1 {
		return data[0], 0
	}
	ret := data[0]
	pos := 0
	for i := 1; i < size; i++ {
		if cmp(ret, data[i]) > 0 {
			ret = data[i]
			pos = i
		}
	}

	return ret, pos
}

// MinOrdered 使用类型天然顺序（xgeneric.Ordered）在切片中查找最小元素及其首次出现的下标。
// 参数：data 为待搜索切片。
// 返回值：最小元素与其首次出现的下标。
// 内部直接复用 Min 并把 CompareTo 包装为比较函数。空切片的 panic 行为继承自 Min。
func MinOrdered[E xgeneric.Ordered](data []E) (E, int) {
	return Min(data, func(e1, e2 E) int {
		return CompareTo(e1, e2)
	})
}

// Max 在切片中查找最大元素及其首次出现的下标。
// 参数：data 为待搜索切片，cmp 为自定义比较函数（语义同 tools.CMP，返回值小于 0 表示 e1<e2）。
// 返回值：最大元素与其首次出现的下标。
// 注意：当前实现没有显式处理 size==0 的情况，调用空切片会导致下标越界 panic；size==1 时直接返回 (data[0], 0)。
func Max[E any](data []E, cmp tools.CMP[E]) (E, int) {
	size := len(data)
	if size == 1 {
		return data[0], 0
	}
	ret := data[0]
	pos := 0
	for i := 1; i < size; i++ {
		if cmp(ret, data[i]) < 0 {
			ret = data[i]
			pos = i
		}
	}
	return ret, pos
}

// MaxOrdered 使用类型天然顺序（xgeneric.Ordered）在切片中查找最大元素及其首次出现的下标。
// 参数：data 为待搜索切片。
// 返回值：最大元素与其首次出现的下标。
// 内部直接复用 Max 并把 CompareTo 作为比较函数。空切片的 panic 行为继承自 Max。
func MaxOrdered[E xgeneric.Ordered](data []E) (E, int) {
	return Max(data, CompareTo[E])
}

// ReplaceAll 将切片中所有与 oldVal 相等的元素替换为 newVal（就地操作）。
// 参数：data 为待操作切片，oldVal 为待替换值，newVal 为替换后的值，euqal 为自定义相等判定函数。
// 返回值：无。当 data 为 nil 时直接返回，不做任何处理；空切片调用安全。
func ReplaceAll[E any](data []E, oldVal, newVal E, euqal tools.EQL[E]) {
	if data == nil {
		return
	}
	size := len(data)
	for i := 0; i < size; i++ {
		if euqal(data[i], oldVal) {
			data[i] = newVal
		}
	}
}

// ReplaceOrderedAll 使用 == 运算符将切片中所有与 oldVal 相等的元素替换为 newVal（适用于 xgeneric.Ordered 类型）。
// 参数：data 为待操作切片，oldVal 为待替换值，newVal 为替换后的值。
// 返回值：无。内部直接调用 ReplaceAll 并传入 Equals 作为相等函数。
func ReplaceOrderedAll[E xgeneric.Ordered](data []E, oldVal, newVal E) {
	ReplaceAll(data, oldVal, newVal, Equals[E])
}

// EqualWith 判断两个切片在长度相同且对应位置元素逐一相等时为 true。
// 参数：data 与 other 为待比较切片，euqal 为自定义相等判定函数。
// 返回值：两个切片长度相等且对应下标元素均通过相等判定时返回 true，否则返回 false。
func EqualWith[E any](data, other []E, euqal tools.EQL[E]) bool {
	s1, s2 := len(data), len(other)
	if s1 != s2 {
		return false
	}

	for i := 0; i < s1; i++ {
		if !euqal(data[i], other[i]) {
			return false
		}
	}
	return true
}

// EqualWithOrdered 使用 == 运算符判断两个切片是否完全相等（适用于 xgeneric.Ordered 类型）。
// 参数：data 与 other 为待比较切片。
// 返回值：长度相同且对应位置元素都相等时返回 true，否则返回 false。
func EqualWithOrdered[E xgeneric.Ordered](data, other []E) bool {
	return EqualWith(data, other, Equals[E])
}

// Filter 按用户提供的判定函数过滤切片，保留 evaluate 返回 false 的元素（注意：原注释命名为 tester，但实现为“保留未通过的元素”）。
// 参数：data 为待过滤切片，evaluate 为针对每个元素的判定函数。
// 返回值：包含所有 evaluate(v) 返回 false 的元素的新切片（顺序与原切片一致）。空切片返回长度为 0 的非 nil 切片。
func Filter[E any](data []E, evaluate tools.Evaluate[E]) []E {
	ret := make([]E, 0)
	for _, v := range data {
		if !evaluate(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

// IndexOfSubArrayReturns the starting position of the first occurrence of the specified
//
//	target array within the specified source array
func IndexOfSubArray[E any](data, sub []E, euqal tools.EQL[E]) int {
	s1, s2 := len(data), len(sub)
	if s2 == 0 {
		return -1
	}

	if s2 > s1 {
		return -1
	}

	checkSize := s2
	beginPos := 0
	for beginPos+checkSize <= s1 {
		if EqualWith(data[beginPos:beginPos+checkSize], sub, euqal) {
			return beginPos
		}
		beginPos++
	}
	return -1
}

// IndexOfSubOrderedArray the starting position of the first occurrence of the specified
//
//	target array within the specified source array
func IndexOfSubOrderedArray[E xgeneric.Ordered](data, sub []E) int {
	return IndexOfSubArray(data, sub, Equals[E])
}

// LastIndexOfSubArray the last starting position of the first occurrence of the specified
//
//	target array within the specified source array
func LastIndexOfSubArray[E any](data, sub []E, euqal tools.EQL[E]) int {
	s1, s2 := len(data), len(sub)
	if s2 == 0 {
		return -1
	}

	if s2 > s1 {
		return -1
	}

	checkSize := s2
	beginPos := s1 - checkSize
	for beginPos >= 0 {
		if EqualWith(data[beginPos:beginPos+checkSize], sub, euqal) {
			return beginPos
		}
		beginPos--
	}

	return -1
}

// LastIndexOfSubArray the last starting position of the first occurrence of the specified
//
//	target array within the specified source array
func LastIndexOfSubOrderedArray[E xgeneric.Ordered](data, sub []E) int {
	return LastIndexOfSubArray(data, sub, Equals[E])
}

// Disjoint Returns true if the two specified collections have no
// elements in common.
func Disjoint[E any](data []E, other []E, euqal tools.EQL[E]) bool {
	s1, s2 := len(data), len(other)
	if s1 == 0 || s2 == 0 {
		return true
	}

	for i := 0; i < s1; i++ {
		if Contains(other, data[i], euqal) {
			return false
		}
	}

	return true
}

// Disjoint Returns true if the two specified collections have no
// elements in common.
func DisjointOrdered[E xgeneric.Ordered](data []E, other []E) bool {
	return Disjoint(data, other, Equals[E])
}

// Rotate Rotates the elements in the specified array by the specified distance.
// For example, suppose a string array arr := []string{"t", "a", "n", "k", "s"}.
// After invoking arrays.Rotate(arr, 1) (or arrays.Rotate(arr, -4)),  output is [s, t, a, n, k].
func Rotate[E any](data []E, distance int) {
	size := len(data)
	if size == 0 {
		return
	}
	mid := -distance % size
	if mid < 0 {
		mid += size
	}
	if mid == 0 {
		return
	}

	Reverse(data[0:mid])
	Reverse(data[mid:size])
	Reverse(data)
}

func getFreq[E xgeneric.Ordered](key E, mapa map[E]int) int {
	v, exist := mapa[key]
	if !exist {
		return 0
	}
	return v
}

// UnionOrdered returns a array containing the union
// of the given array.
func UnionOrdered[E xgeneric.Ordered](data, other []E) []E {
	ret := make([]E, 0)

	mapa := getCardinalityMap(data)
	mapb := getCardinalityMap(other)

	merged := xmap.AddAll(mapa, mapb)
	for k := range merged {
		i := 0
		for m := maths.Max(int(getFreq(k, mapa)), int(getFreq(k, mapb))); i < m; i++ {
			ret = append(ret, k)
		}
	}

	return ret
}

// IntersectionOrdered returns a array containing the intersection
// of the given array.
func IntersectionOrdered[E xgeneric.Ordered](data, other []E) []E {
	ret := make([]E, 0)

	mapa := getCardinalityMap(data)
	mapb := getCardinalityMap(other)

	merged := xmap.AddAll(mapa, mapb)
	for k := range merged {
		i := 0
		for m := maths.Min(int(getFreq(k, mapa)), int(getFreq(k, mapb))); i < m; i++ {
			ret = append(ret, k)
		}
	}

	return ret
}

// DisjunctionOrdered returns a array containing the exclusive disjunction
// (symmetric difference) of the given array
func DisjunctionOrdered[E xgeneric.Ordered](data, other []E) []E {
	ret := make([]E, 0)

	mapa := getCardinalityMap(data)
	mapb := getCardinalityMap(other)

	merged := xmap.AddAll(mapa, mapb)
	for k := range merged {
		i := 0
		m := maths.Max(int(getFreq(k, mapa)), int(getFreq(k, mapb))) - maths.Min(int(getFreq(k, mapa)), int(getFreq(k, mapb)))
		for ; i < m; i++ {
			ret = append(ret, k)
		}
	}

	return ret
}

// Subtract returns a new array containing data - other.
func Substract[E any](data, other []E, equal tools.EQL[E]) []E {
	ret := Clone(data)

	for _, e := range other {
		ret, _ = Remove(ret, e, equal)
	}
	return ret
}

// SubtractOrdered returns a new array containing data - other.
func SubstractOrdered[E xgeneric.Ordered](data, other []E) []E {
	ret := Clone(data)
	for _, v := range other {
		ret, _ = RemoveOrdered(ret, v)
	}

	return ret
}

func getCardinalityMap[E xgeneric.Ordered](data []E) map[E]int {
	ret := make(map[E]int)
	for _, e := range data {
		_, exist := ret[e]
		if exist {
			ret[e] = ret[e] + 1
		} else {
			ret[e] = 1
		}

	}
	return ret
}

// CompareTo Compares this object with the specified object for order.  Returns a
//
//	negative integer, zero, or a positive integer as this object is less
//	 than, equal to, or greater than the specified object.
func CompareTo[E xgeneric.Ordered](e1, e2 E) int {
	if e1 > e2 {
		return 1
	} else if e1 < e2 {
		return -1
	}
	return 0
}

// Clone Copies all of the elements and return a new array
func Clone[E any](data []E) []E {
	size := len(data)
	ret := make([]E, size, cap(data))
	copy(ret, data)
	return ret
}

// Insert target value in array, if index < 0 or > len(data) will do noting
func Insert[E any](data []E, index int, v E) []E {
	if index < 0 || index > len(data) {
		return data
	}
	if index == len(data) {
		return append(data, v)
	}
	var empty E
	ret := append(data, empty)
	copy(ret[index+1:], ret[index:])
	ret[index] = v
	return ret
}

// Equals to compare two value is equal
func Equals[E xgeneric.Ordered](e1, e2 E) bool {
	return e1 == e2
}

// CreateAndFill create array with the same element t
func CreateAndFill[E any](size int, defaultElementValue E) []E {
	ret := make([]E, size)
	for i := 0; i < size; i++ {
		ret[i] = defaultElementValue
	}
	return ret
}

// AsList convert to array
func AsList[E any](e ...E) []E {
	if e == nil {
		return nil
	}
	ret := make([]E, len(e))
	copy(ret, e)
	return ret
}

// Join element as string with split
func Join[E any](eles []E, split string) string {
	sz := len(eles)
	ret := ""
	if sz == 0 {
		return ret
	}

	for i := 0; i < sz; i++ {
		ret = ret + fmt.Sprintf("%v", eles[i])
		if i+1 < sz {
			ret = ret + split
		}
	}

	return ret
}

// Swap a series of elements in the given array.
func Swap[E any](array []E, offset1, offset2, l int) {
	if len(array) == 0 || offset1 >= len(array) || offset2 >= len(array) {
		return
	}
	if offset1 < 0 {
		offset1 = 0
	}
	if offset2 < 0 {
		offset2 = 0
	}
	if offset1 == offset2 {
		return
	}
	l = maths.Min(maths.Min(l, len(array)-offset1), len(array)-offset2)
	for i := 0; i < l; i++ {
		array[offset1], array[offset2] = array[offset2], array[offset1]
		offset1++
		offset2++
	}
}

// Addall add all the elements of the given arrays into a new array.
func Addall[E any](array []E, eles ...E) []E {
	ret := make([]E, len(array)+len(eles))
	copy(ret[:len(array)], array)
	if len(eles) > 0 {
		copy(ret[len(array):], eles)
	}
	return ret
}
