package xdatabase

import "fmt"

func Limit(page, pageSize int) (limit string) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	limit = fmt.Sprintf("%d,%d", (page-1)*pageSize, pageSize)
	return
}

func Offset(page, pageSize int) (offset int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	return (page - 1) * pageSize
}
