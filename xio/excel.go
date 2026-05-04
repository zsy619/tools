package xio

import "fmt"

// 用来匹配的字符A-Z
var Slice = []string{
	"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
	"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

// Div
/**
 * @description: 数字转字母
 * @param {int} num 需要转化的数字，如excel第几列
 * @return string
 */
// Deprecated: Div is no longer recommended,
// use ExcelColumnNumberToName instead.
func Div(num int) string {
	var (
		Str  string = ""
		k    int
		temp []int // 保存转化后每一位数据的值，然后通过索引的方式匹配A-Z
	)

	if num > 26 { // 数据大于26需要进行拆分
		for {
			k = num % 26 // 从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个
			if k == 0 {
				temp = append(temp, 26)
				k = 26
			} else {
				temp = append(temp, k)
			}
			num = (num - k) / 26 // 减去Num最后一位数的值，因为已经记录在temp中
			if num <= 26 {       // 小于等于26直接进行匹配，不需要进行数据拆分
				temp = append(temp, num)
				break
			}
		}
	} else {
		return Slice[num]
	}
	for _, value := range temp {
		Str = Slice[value] + Str // 因为数据切分后存储顺序是反的，所以Str要放在后面
	}
	return Str
}

func ExcelColumnNumberToName(columnNumber int) string {
	var columnName string
	for columnNumber > 0 {
		columnNumber--
		columnName = fmt.Sprint('A'+columnNumber%26) + columnName
		columnNumber = columnNumber / 26
	}
	return columnName
}

func ExcelColumnNumberToNextName(prevColumnNumber int) (int, string) {
	columnNumber := prevColumnNumber + 1
	var columnName string
	for columnNumber > 0 {
		columnNumber--
		columnName = fmt.Sprint('A'+columnNumber%26) + columnName
		columnNumber = columnNumber / 26
	}
	return columnNumber, columnName
}
