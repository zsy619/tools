package xdatabase

import (
	"github.com/zsy619/tools"
	"gorm.io/gorm"
)

// CrudRepository 基于 gorm.DB 的通用 CRUD 仓储实现。
//
// 字段说明：
//   - store: 底层 gorm.DB 连接。
type CrudRepository[T Entity[ID], ID comparable] struct {
	store *gorm.DB
}

// NewDefaultRepository 构造一个基于默认实现的 Repository。
//
// 参数：
//   - store: 已初始化的 *gorm.DB。
//
// 返回值：实现 Repository[T, ID] 接口的实例。
func NewDefaultRepository[T Entity[ID], ID comparable](store *gorm.DB) Repository[T, ID] {
	return &CrudRepository[T, ID]{store: store}
}

// Save 在数据库中插入单条记录。
//
// 参数：
//   - v: 待保存的实体。
//
// 返回值：底层 gorm 错误，已被 SafeCall 包装。
func (l *CrudRepository[T, ID]) Save(v T) error {
	return tools.SafeCall(func() error {
		return l.store.Model(new(T)).Create(v).Error
	})
}

// SaveAll 在同一个事务中批量保存多条记录，任一失败将回滚。
//
// 参数：
//   - entities: 待保存的实体列表。
//
// 返回值：底层 gorm 错误或事务错误。
func (l *CrudRepository[T, ID]) SaveAll(entities []T) error {
	return tools.SafeCall(func() error {
		return l.store.Transaction(func(session *gorm.DB) error {
			for _, entity := range entities {
				err := session.Model(new(T)).Save(entity).Error
				if nil != err {
					return err
				}
			}
			return nil
		})
	})
}

// Find 按字段等值匹配查询多条记录。
//
// 参数：
//   - v: 用作查询条件的实体模板（其非零字段会被当作 WHERE 条件）。
//
// 返回值：
//   - []T: 匹配的实体列表，未命中时为空切片。
//   - error: 查询失败时返回错误。
func (l *CrudRepository[T, ID]) Find(v T) ([]T, error) {
	var result []T
	return result, tools.SafeCall(func() error {
		return l.store.Model(new(T)).Where(v).Find(&result).Error
	})
}

// FindByIdInBatch 根据主键集合批量查询。
//
// 参数：
//   - ids: 主键列表。
//
// 返回值：
//   - []T: 命中的实体列表。
//   - error: 查询失败时返回错误。
func (l *CrudRepository[T, ID]) FindByIdInBatch(ids []ID) ([]T, error) {
	var result []T
	return result, tools.SafeCall(func() error {
		return l.store.Model(new(T)).Find(new(T), ids).Error
	})
}

// Update 按主键更新一条记录（依赖 Entity.EntityID() 返回主键）。
//
// 参数：
//   - v: 包含新字段值与主键的实体。
//
// 返回值：底层 gorm 错误。
func (l *CrudRepository[T, ID]) Update(v T) error {
	return tools.SafeCall(func() error {
		return l.store.Model(new(T)).Where("id = ?", v.EntityID()).Updates(v).Error
	})
}

// UpdateAll 在事务中按主键批量更新多条记录。
//
// 参数：
//   - entities: 待更新的实体列表。
//
// 返回值：事务执行过程中产生的错误。
func (l *CrudRepository[T, ID]) UpdateAll(entities []T) error {
	return tools.SafeCall(func() error {
		return l.store.Transaction(func(session *gorm.DB) error {
			for _, entity := range entities {
				err := session.Model(new(T)).Where("id = ?", entity.EntityID()).Updates(entity).Error
				if nil != err {
					return err
				}
			}
			return nil
		})
	})
}

// Delete 在事务中按主键删除一条记录。
//
// 参数：
//   - v: 包含主键的实体。
//
// 返回值：事务执行过程中产生的错误。
func (l *CrudRepository[T, ID]) Delete(v T) error {
	return tools.SafeCall(func() error {
		return l.store.Transaction(func(session *gorm.DB) error {
			return session.Model(new(T)).Delete(v).Error
		})
	})
}

// DeleteByIdInBatch 在事务中按主键集合批量删除。
//
// 参数：
//   - ids: 待删除的主键列表。
//
// 返回值：底层 gorm 错误。
func (l *CrudRepository[T, ID]) DeleteByIdInBatch(ids []ID) error {
	return tools.SafeCall(func() error {
		return l.store.Transaction(func(session *gorm.DB) error {
			return l.store.Model(new(T)).Delete(new(T), ids).Error
		})
	})
}

/*

type TenantRepository struct {
	infrastructure.Repository[*tenantModel.Tenant, string]
	store *gorm.DB
}

func NewTenantRepository(store *gorm.DB) *TenantRepository {
	store.AutoMigrate(&tenantModel.Tenant{})
	return &TenantRepository{
		Repository: infrastructure.NewCrudRepository[*tenantModel.Tenant, string](store),
		store:      store,
	}
}

*/
