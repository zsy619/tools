package xphp

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

// ArrayMap 是字符串键到任意值的映射，等价于 PHP 中带字符串键的关联数组。
type ArrayMap map[string]interface{}

// Case 表示大小写转换方向，供 Array_change_key_case 使用。
type Case int

const (
	// CASE_UPPER 表示把 key 转为大写。
	CASE_UPPER Case = 1
	// CASE_LOWER 表示把 key 转为小写。
	CASE_LOWER Case = 0
)

// Array_diff 计算数组差集：返回在第一个 ArrayMap 中但不在其它 ArrayMap 中的元素。
// 当前为占位实现，TODO。
func Array_diff(m ...ArrayMap) ArrayMap {
	// todo
	return ArrayMap{}
}

// Array_change_key_case 按 cs 指定的大小写方向转换 ArrayMap 的所有 key 并返回新 map。
// cs 取值 CASE_LOWER 或 CASE_UPPER，其它值返回空 map。
func Array_change_key_case(arr ArrayMap, cs Case) ArrayMap {
	var tmp ArrayMap = ArrayMap{}
	if cs == CASE_LOWER {
		for k, v := range arr {
			tmp[strings.ToLower(k)] = v
		}
	} else if cs == CASE_UPPER {
		for k, v := range arr {
			tmp[strings.ToUpper(k)] = v
		}
	}
	return tmp
}

// ArrayFill 实现 PHP array_fill()：从 startIndex 起填充 num 个 value。
// 返回 map[int]interface{}，key 连续递增。
func ArrayFill(startIndex int, num uint, value interface{}) map[int]interface{} {
	m := make(map[int]interface{})
	var i uint
	for i = 0; i < num; i++ {
		m[startIndex] = value
		startIndex++
	}
	return m
}

// ArrayFlip 实现 PHP array_flip()：交换 map 的 key 与 value。
// 若原 value 不唯一，结果中只保留最后一个 key。
func ArrayFlip(m map[interface{}]interface{}) map[interface{}]interface{} {
	n := make(map[interface{}]interface{})
	for i, v := range m {
		n[v] = i
	}
	return n
}

// ArrayFlipString 是 ArrayFlip 针对字符串键 map 的便利包装。
func ArrayFlipString(m map[string]interface{}) map[interface{}]interface{} {
	n := make(map[interface{}]interface{})
	for i, v := range m {
		n[v] = i
	}
	return n
}

// ArrayFlipInt 是 ArrayFlip 针对整型键 map 的便利包装。
func ArrayFlipInt(m map[int]interface{}) map[interface{}]interface{} {
	n := make(map[interface{}]interface{})
	for i, v := range m {
		n[v] = i
	}
	return n
}

// ArrayFlipInt64 是 ArrayFlip 针对 int64 键 map 的便利包装。
func ArrayFlipInt64(m map[int64]interface{}) map[interface{}]interface{} {
	n := make(map[interface{}]interface{})
	for i, v := range m {
		n[v] = i
	}
	return n
}

// ArrayKeys 返回 map 的所有 key；返回切片顺序不固定。
func ArrayKeys(elements map[interface{}]interface{}) []interface{} {
	i, keys := 0, make([]interface{}, len(elements))
	for key := range elements {
		keys[i] = key
		i++
	}
	return keys
}

// ArrayValues 返回 map 的所有 value；返回切片顺序不固定。
func ArrayValues(elements map[interface{}]interface{}) []interface{} {
	i, vals := 0, make([]interface{}, len(elements))
	for _, val := range elements {
		vals[i] = val
		i++
	}
	return vals
}

// ArrayMerge 将多个切片按顺序拼接为一个新切片；任一输入为 nil 也能正确处理。
func ArrayMerge(ss ...[]interface{}) []interface{} {
	n := 0
	for _, v := range ss {
		n += len(v)
	}
	s := make([]interface{}, 0, n)
	for _, v := range ss {
		s = append(s, v...)
	}
	return s
}

// ArrayChunk 实现 PHP array_chunk()：将切片按 size 切分。
// size < 1 时 panic。
func ArrayChunk(s []interface{}, size int) [][]interface{} {
	if size < 1 {
		panic("size: cannot be less than 1")
	}
	length := len(s)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]interface{}
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		n = append(n, s[i*size:end])
		i++
	}
	return n
}

// ArrayPad 实现 PHP array_pad()：用 val 把切片填充到 size 长度。
// size == 0、size 超出当前长度方向时直接返回原切片。
func ArrayPad(s []interface{}, size int, val interface{}) []interface{} {
	if size == 0 || (size > 0 && size < len(s)) || (size < 0 && size > -len(s)) {
		return s
	}
	n := size
	if size < 0 {
		n = -size
	}
	n -= len(s)
	tmp := make([]interface{}, n)
	for i := 0; i < n; i++ {
		tmp[i] = val
	}
	if size > 0 {
		return append(s, tmp...)
	}
	return append(tmp, s...)
}

// ArrayRand 实现 PHP array_rand()：以随机顺序返回切片的元素。
// 使用当前时间作为随机种子。
func ArrayRand(elements []interface{}) []interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := make([]interface{}, len(elements))
	for i, v := range r.Perm(len(elements)) {
		n[i] = elements[v]
	}
	return n
}

// ArrayColumn 实现 PHP array_column()：从 input 中提取 columnKey 列的值。
// 元素中不包含 columnKey 时跳过。
func ArrayColumn(input map[string]map[string]interface{}, columnKey string) []interface{} {
	columns := make([]interface{}, 0, len(input))
	for _, val := range input {
		if v, ok := val[columnKey]; ok {
			columns = append(columns, v)
		}
	}
	return columns
}

// ArrayPush 实现 PHP array_push()：把元素追加到切片末尾，返回新长度。
func ArrayPush(s *[]interface{}, elements ...interface{}) int {
	*s = append(*s, elements...)
	return len(*s)
}

// ArrayPop 实现 PHP array_pop()：弹出切片末尾元素；空切片返回 nil。
func ArrayPop(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]
	return e
}

// ArrayUnshift 实现 PHP array_unshift()：把元素插入到切片头部，返回新长度。
func ArrayUnshift(s *[]interface{}, elements ...interface{}) int {
	*s = append(elements, *s...)
	return len(*s)
}

// ArrayShift 实现 PHP array_shift()：移除并返回切片首元素；空切片返回 nil。
func ArrayShift(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	f := (*s)[0]
	*s = (*s)[1:]
	return f
}

// ArrayKeyExists 判断 key 是否存在于 map 中。
func ArrayKeyExists(key interface{}, m map[interface{}]interface{}) bool {
	_, ok := m[key]
	return ok
}

// ArrayCombine 实现 PHP array_combine()：使用 s1 的元素作 key、s2 的元素作 value。
// 长度不一致时 panic。
func ArrayCombine(s1, s2 []interface{}) map[interface{}]interface{} {
	if len(s1) != len(s2) {
		panic("the number of elements for each slice isn't equal")
	}
	m := make(map[interface{}]interface{}, len(s1))
	for i, v := range s1 {
		m[v] = s2[i]
	}
	return m
}

// ArrayReverse 反转切片并返回（原地修改）。
func ArrayReverse(s []interface{}) []interface{} {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// ArrayStringReverse 反转字符串切片（原地修改）。
func ArrayStringReverse(arr *[]string) {
	length := len(*arr)
	var temp string
	for i := 0; i < length/2; i++ {
		temp = (*arr)[i]
		(*arr)[i] = (*arr)[length-1-i]
		(*arr)[length-1-i] = temp
	}
}

// Implode 实现 PHP implode()：使用 seq 连接 list 切片的字符串形式。
// list 不是切片时返回空串；元素会自动调用 fmt.Sprint 转为字符串。
func Implode(seq string, list interface{}) string {
	listValue := reflect.Indirect(reflect.ValueOf(list))
	if listValue.Kind() != reflect.Slice {
		return ""
	}
	count := listValue.Len()
	listStr := make([]string, 0, count)
	for i := 0; i < count; i++ {
		v := listValue.Index(i)
		if str, err := getValue(v); err == nil {
			listStr = append(listStr, str)
		}
	}
	return strings.Join(listStr, seq)
}

// getValue 将 reflect.Value 转为字符串：指针解引用，其余类型用 fmt.Sprint。
func getValue(value reflect.Value) (res string, err error) {
	switch value.Kind() {
	case reflect.Ptr:
		res, err = getValue(value.Elem())
	default:
		res = fmt.Sprint(value.Interface())
	}
	return
}

// InArray 实现 PHP in_array()：判断 needle 是否出现在 haystack 中。
// haystack 必须是切片、数组或 map；其它类型会 panic。
// 使用 reflect.DeepEqual 比较元素；map 情况下比较 value 而非 key。
func InArray(needle interface{}, haystack interface{}) bool {
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
	default:
		panic("haystack: haystack type muset be slice, array or map")
	}

	return false
}
