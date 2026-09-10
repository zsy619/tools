package xdatabase

// Entity 描述可被仓储管理的实体接口，必须能返回自身的主键 ID。
type Entity[ID comparable] interface {
	EntityID() ID
}

// Repository 通用仓储接口，定义了实体的 CRUD 操作。
//
// 泛型参数：
//   - T: 实体类型，必须实现 Entity[ID]。
//   - ID: 主键类型，必须 comparable。
type Repository[T Entity[ID], ID comparable] interface {
	// Save 保存单条实体。
	Save(T) error
	// SaveAll 在事务中批量保存多条实体。
	SaveAll([]T) error

	// Update 按主键更新单条实体。
	Update(T) error
	// UpdateAll 在事务中按主键批量更新多条实体。
	UpdateAll([]T) error

	// Delete 按主键删除单条实体。
	Delete(T) error
	// DeleteByIdInBatch 按主键集合批量删除。
	DeleteByIdInBatch([]ID) error

	// Find 按字段等值匹配查询多条记录。
	Find(T) ([]T, error)
	// FindByIdInBatch 按主键集合批量查询。
	FindByIdInBatch([]ID) ([]T, error)
}
