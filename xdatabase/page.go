package xdatabase

import "gorm.io/gorm"

// Page 通用分页结果结构。
//
// 字段说明：
//   - CurrentPage: 当前页码（从 1 开始）。
//   - PageSize: 每页条数。
//   - Total: 命中总数。
//   - Pages: 总页数（向上取整）。
//   - Data: 当前页数据。
type Page[T any] struct {
	CurrentPage int64
	PageSize    int64
	Total       int64
	Pages       int64
	Data        []T
}

// SelectPage 一次性执行 count 和分页查询，并将结果回填到 page 中。
//
// 参数：
//   - db: 已初始化的 *gorm.DB。
//   - wrapper: 查询条件，等价于 gorm 的 Where 参数。
//
// 返回值：当查询过程中 gorm 报错时返回该错误；wrapper 为空或命中 0 条记录时 page.Data 会被置为空切片且不返回错误。
func (page *Page[T]) SelectPage(db *gorm.DB, wrapper map[string]interface{}) (e error) {
	e = nil
	var model T
	db.Model(&model).Where(wrapper).Count(&page.Total)
	if page.Total == 0 {
		// 没有符合条件的数据，直接返回一个T类型的空列表
		page.Data = []T{}
		return
	}
	// 查询结果可以直接存到Page的Data字段中，因为编译的时候page.Data是有确定类型的
	e = db.Model(&model).Where(wrapper).Scopes(Paginate(page)).Find(&page.Data).Error
	return
}

// Paginate加上T就行
// Paginate 返回一个 gorm Scope，用于在查询中应用分页参数。
//
// 参数：
//   - page: 已填充 CurrentPage、PageSize 的分页对象；执行后会回填 Pages。
//
// 返回值：返回 func(db *gorm.DB) *gorm.DB，可在链式调用中通过 Scopes 注入。
//
// 副作用：会修正 CurrentPage <= 0、PageSize <= 0、PageSize > 100 等异常输入。
func Paginate[T any](page *Page[T]) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page.CurrentPage <= 0 {
			page.CurrentPage = 0
		}
		switch {
		case page.PageSize > 100:
			page.PageSize = 100
		case page.PageSize <= 0:
			page.PageSize = 10
		}
		page.Pages = page.Total / page.PageSize
		if page.Total%page.PageSize != 0 {
			page.Pages++
		}
		p := page.CurrentPage
		if page.CurrentPage > page.Pages {
			p = page.Pages
		}
		size := page.PageSize
		offset := int((p - 1) * size)
		return db.Offset(offset).Limit(int(size))
	}
}
