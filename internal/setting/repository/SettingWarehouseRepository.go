package repository

import (
	"net/url"
	"service/internal/pkg/config"
	"service/internal/pkg/core"
	"service/internal/pkg/form"
	"service/internal/pkg/model"

	"gorm.io/gorm"
)

type SettingWarehouseRepository interface {
	core.TransactionRepository
	core.PaginateRepository[model.SettingWarehouse]
	core.FirstByFormRepository[model.SettingWarehouse, form.ComponentFilterForm]
	core.FindByFormRepository[model.SettingWarehouse, form.ComponentFilterForm]

	Create(form form.SettingWarehouseForm) model.SettingWarehouse
	Update(warehouse model.SettingWarehouse, warehouseForm form.SettingWarehouseForm) model.SettingWarehouse
	Delete(warehouse model.SettingWarehouse)
}

func NewSettingWarehouseRepository(args ...*gorm.DB) SettingWarehouseRepository {
	repo := settingWarehouseRepository{}

	if len(args) > 0 {
		repo.transaction = args[0]
	}

	return &repo
}

type settingWarehouseRepository struct {
	transaction *gorm.DB
}

func (s *settingWarehouseRepository) SetTransaction(tx *gorm.DB) {
	s.transaction = tx
}

func (s *settingWarehouseRepository) Paginate(parameter url.Values) ([]model.SettingWarehouse, interface{}, error) {
	var warehouses []model.SettingWarehouse

	return warehouses, warehouses, nil
}

func (s *settingWarehouseRepository) FirstByForm(form form.ComponentFilterForm) model.SettingWarehouse {
	//TODO implement me
	panic("implement me")
}

func (s *settingWarehouseRepository) FindByForm(form form.ComponentFilterForm) []model.SettingWarehouse {
	//TODO implement me
	panic("implement me")
}

func (s *settingWarehouseRepository) Create(form form.SettingWarehouseForm) model.SettingWarehouse {
	//TODO implement me
	panic("implement me")
}

func (s *settingWarehouseRepository) Update(warehouse model.SettingWarehouse, warehouseForm form.SettingWarehouseForm) model.SettingWarehouse {
	//TODO implement me
	panic("implement me")
}

func (s *settingWarehouseRepository) Delete(warehouse model.SettingWarehouse) {
	//TODO implement me
	panic("implement me")
}

/** --- Unexported Functions --- */

func (s *settingWarehouseRepository) prepareFilterForm(form form.ComponentFilterForm) *gorm.DB {
	query := config.PgSQL

	if form.IDs != nil && len(form.IDs) > 0 {
		query = query.Where("id IN (?)", form.IDs)
	}

	if form.ID > 0 {
		query = query.Where("id =?", form.ID)
	}

	if form.Search != "" {
		query = query.Where("name LIKE ?", "%"+form.Search+"%")
	}

	return query
}
