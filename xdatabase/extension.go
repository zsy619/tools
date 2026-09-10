package xdatabase

import "fmt"

// Limit 生成分页 SQL 的 LIMIT 子串，形如 "offset,size"，可直接拼到 SQL 末尾。
//
// 参数：
//   - page: 当前页码（从 1 开始），<=0 时按 1 处理。
//   - pageSize: 每页条数，<=0 时按 20 处理。
//
// 返回值：形如 "offset,size" 的字符串。
func Limit[T int16 | int | int32 | int64](page, pageSize T) (limit string) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	limit = fmt.Sprintf("%d,%d", (page-1)*pageSize, pageSize)
	return
}

// Offset 计算分页查询的偏移量（从 0 开始）。
//
// 参数：
//   - page: 当前页码（从 1 开始），<=0 时按 1 处理。
//   - pageSize: 每页条数，<=0 时按 20 处理。
//
// 返回值：应跳过的记录数，等价于 (page-1)*pageSize。
func Offset[T int16 | int | int32 | int64](page, pageSize T) (offset T) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	return (page - 1) * pageSize
}
