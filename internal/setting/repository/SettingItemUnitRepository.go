package repository

import (
	"net/url"
	"service/internal/pkg/config"
	"service/internal/pkg/core"
	gxErr "service/internal/pkg/error"
	"service/internal/pkg/form"
	"service/internal/pkg/model"

	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	"gorm.io/gorm"
)

type SettingItemUnitRepository interface {
	core.TransactionRepository

	Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemUnit, interface{}, error)
	FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemUnit
	Create(form form.SettingItemUnitForm) model.ItemUnit
	Update(itemUnit model.ItemUnit, form form.SettingItemUnitForm) model.ItemUnit
	Delete(itemUnit model.ItemUnit)
}

func NewSettingItemUnitRepository(args ...*gorm.DB) SettingItemUnitRepository {
	repository := settingItemUnitRepository{}

	if len(args) > 0 {
		repository.transaction = args[0]
	}

	return &repository
}

type settingItemUnitRepository struct {
	transaction *gorm.DB
}

func (repo *settingItemUnitRepository) SetTransaction(tx *gorm.DB) {
	repo.transaction = tx
}

func (repo *settingItemUnitRepository) Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemUnit, interface{}, error) {
	var units []model.ItemUnit

	fromDate, toDate := core.SetDateRange(parameter)
	query := config.PgSQL.Where(`"createdAt" BETWEEN ? AND ?`, fromDate, toDate)

	if search := parameter.Get("search"); len(search) > 3 {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query = query.Order("id DESC")

	units, pagination, err := xtrememodel.Paginate(query, parameter, model.ItemUnit{})
	if err != nil {
		return nil, nil, err
	}

	return units, pagination, nil
}

func (repo *settingItemUnitRepository) FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemUnit {
	var itemUnit model.ItemUnit

	query := config.PgSQL
	if len(args) > 0 {
		query = args[0](query)
	}

	err := query.First(&itemUnit, "id = ?", id).Error
	if err != nil {
		gxErr.ErrXtremeSettingGet(itemUnit.FeatureName(), err.Error())
	}

	return itemUnit
}

func (repo *settingItemUnitRepository) Create(form form.SettingItemUnitForm) model.ItemUnit {
	itemUnit := model.ItemUnit{
		Name:         form.Name,
		Abbreviation: form.Abbreviation,
		Type:         form.Type,
		IsBaseUnit:   form.IsBaseUnit,
		Conversion:   form.Conversion,
	}

	err := repo.transaction.Create(&itemUnit).Error
	if err != nil {
		gxErr.ErrXtremeSettingCreate(itemUnit.FeatureName(), err.Error())
	}

	return itemUnit
}

func (repo *settingItemUnitRepository) Update(itemUnit model.ItemUnit, form form.SettingItemUnitForm) model.ItemUnit {
	itemUnit.Name = form.Name
	itemUnit.Abbreviation = form.Abbreviation
	itemUnit.Type = form.Type
	itemUnit.IsBaseUnit = form.IsBaseUnit
	itemUnit.Conversion = form.Conversion

	err := repo.transaction.Save(&itemUnit).Error
	if err != nil {
		gxErr.ErrXtremeSettingUpdate(itemUnit.FeatureName(), err.Error())
	}

	return itemUnit
}

func (repo *settingItemUnitRepository) Delete(itemUnit model.ItemUnit) {
	err := repo.transaction.Delete(&itemUnit).Error
	if err != nil {
		gxErr.ErrXtremeSettingDelete(itemUnit.FeatureName(), err.Error())
	}
}
