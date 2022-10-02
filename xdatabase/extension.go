package xdatabase

import "fmt"

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

func Offset[T int16 | int | int32 | int64](page, pageSize T) (offset T) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	return (page - 1) * pageSize
}
