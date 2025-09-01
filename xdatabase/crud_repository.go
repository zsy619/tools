package xdatabase

import (
	"github.com/zsy619/tools"
	"gorm.io/gorm"
)

type CrudRepository[T Entity[ID], ID comparable] struct {
	store *gorm.DB
}

func NewDefaultRepository[T Entity[ID], ID comparable](store *gorm.DB) Repository[T, ID] {
	return &CrudRepository[T, ID]{store: store}
}

func (l *CrudRepository[T, ID]) Save(v T) error {
	return tools.SafeCall(func() error {
		return l.store.Model(new(T)).Create(v).Error
	})
}

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

func (l *CrudRepository[T, ID]) Find(v T) ([]T, error) {
	var result []T
	return result, tools.SafeCall(func() error {
		return l.store.Model(new(T)).Where(v).Find(&result).Error
	})
}

func (l *CrudRepository[T, ID]) FindByIdInBatch(ids []ID) ([]T, error) {
	var result []T
	return result, tools.SafeCall(func() error {
		return l.store.Model(new(T)).Find(new(T), ids).Error
	})
}

func (l *CrudRepository[T, ID]) Update(v T) error {
	return tools.SafeCall(func() error {
		return l.store.Model(new(T)).Where("id = ?", v.EntityID()).Updates(v).Error
	})
}

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

func (l *CrudRepository[T, ID]) Delete(v T) error {
	return tools.SafeCall(func() error {
		return l.store.Transaction(func(session *gorm.DB) error {
			return session.Model(new(T)).Delete(v).Error
		})
	})
}

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
