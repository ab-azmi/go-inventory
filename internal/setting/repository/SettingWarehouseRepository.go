package repository

import (
	"net/url"
	"service/internal/pkg/config"
	"service/internal/pkg/core"
	error2 "service/internal/pkg/error"
	"service/internal/pkg/form"
	"service/internal/pkg/model"

	xtrememodel "github.com/globalxtreme/go-core/v2/model"
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

func (repo *settingWarehouseRepository) SetTransaction(tx *gorm.DB) {
	repo.transaction = tx
}

func (repo *settingWarehouseRepository) Paginate(parameter url.Values) ([]model.SettingWarehouse, interface{}, error) {
	query := repo.prepareFilterForm(form.ComponentFilterForm{
		Search: parameter.Get("search"),
	})

	warehouses, pagination, err := xtrememodel.Paginate(query, parameter, model.SettingWarehouse{})
	if err != nil {
		return nil, nil, err
	}

	return warehouses, pagination, nil
}

func (repo *settingWarehouseRepository) FirstByForm(form form.ComponentFilterForm) model.SettingWarehouse {
	var warehouse model.SettingWarehouse

	query := repo.prepareFilterForm(form)

	err := query.First(&warehouse).Error
	if err != nil {
		error2.ErrXtremeSettingWarehouseGet(err.Error())
	}

	return warehouse
}

func (repo *settingWarehouseRepository) FindByForm(form form.ComponentFilterForm) []model.SettingWarehouse {
	var warehouses []model.SettingWarehouse

	query := repo.prepareFilterForm(form)

	err := query.Find(&warehouses).Error
	if err != nil {
		error2.ErrXtremeSettingWarehouseGet(err.Error())
	}

	return warehouses
}

func (repo *settingWarehouseRepository) Create(form form.SettingWarehouseForm) model.SettingWarehouse {
	warehouse := model.SettingWarehouse{
		Name:           form.Name,
		Address:        core.StrPtr(form.Address),
		ParentId:       core.UintPtr(form.ParentId),
		BranchOfficeId: form.BranchOfficeId,
	}

	err := repo.transaction.Create(&warehouse).Error
	if err != nil {
		error2.ErrXtremeSettingWarehouseCreate(err.Error())
	}

	return warehouse
}

func (repo *settingWarehouseRepository) Update(warehouse model.SettingWarehouse, form form.SettingWarehouseForm) model.SettingWarehouse {
	warehouse.Name = form.Name
	warehouse.Address = core.StrPtr(form.Address)
	warehouse.ParentId = core.UintPtr(form.ParentId)
	warehouse.BranchOfficeId = form.BranchOfficeId

	err := repo.transaction.Save(&warehouse).Error
	if err != nil {
		error2.ErrXtremeSettingWarehouseUpdate(err.Error())
	}

	return warehouse
}

func (repo *settingWarehouseRepository) Delete(warehouse model.SettingWarehouse) {
	err := repo.transaction.Delete(&warehouse).Error
	if err != nil {
		error2.ErrXtremeSettingWarehouseDelete(err.Error())
	}
}

/** --- Unexported Functions --- */

func (repo *settingWarehouseRepository) prepareFilterForm(form form.ComponentFilterForm) *gorm.DB {
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
