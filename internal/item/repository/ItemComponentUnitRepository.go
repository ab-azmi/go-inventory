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

type ItemComponentUnitRepository interface {
	core.TransactionRepository

	Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemComponentUnit, interface{}, error)
	FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemComponentUnit
	Create(form form.SettingItemUnitForm) model.ItemComponentUnit
	Update(itemUnit model.ItemComponentUnit, form form.SettingItemUnitForm) model.ItemComponentUnit
	Delete(itemUnit model.ItemComponentUnit)
}

func NewItemComponentUnitRepository(args ...*gorm.DB) ItemComponentUnitRepository {
	repository := itemComponentUnitRepository{}

	if len(args) > 0 {
		repository.transaction = args[0]
	}

	return &repository
}

type itemComponentUnitRepository struct {
	transaction *gorm.DB
}

func (repo *itemComponentUnitRepository) SetTransaction(tx *gorm.DB) {
	repo.transaction = tx
}

func (repo *itemComponentUnitRepository) Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemComponentUnit, interface{}, error) {
	var units []model.ItemComponentUnit

	fromDate, toDate := core.SetDateRange(parameter)
	query := config.PgSQL.Where(`"createdAt" BETWEEN ? AND ?`, fromDate, toDate)

	if search := parameter.Get("search"); len(search) > 3 {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query = query.Order("id DESC")

	units, pagination, err := xtrememodel.Paginate(query, parameter, model.ItemComponentUnit{})
	if err != nil {
		return nil, nil, err
	}

	return units, pagination, nil
}

func (repo *itemComponentUnitRepository) FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemComponentUnit {
	var itemUnit model.ItemComponentUnit

	query := config.PgSQL
	if len(args) > 0 {
		query = args[0](query)
	}

	err := query.First(&itemUnit, "id = ?", id).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemUnitGet(err.Error())
	}

	return itemUnit
}

func (repo *itemComponentUnitRepository) Create(form form.SettingItemUnitForm) model.ItemComponentUnit {
	itemUnit := model.ItemComponentUnit{
		Name:         form.Name,
		Abbreviation: form.Abbreviation,
		Type:         form.Type,
		IsBaseUnit:   form.IsBaseUnit,
		Conversion:   form.Conversion,
	}

	err := repo.transaction.Create(&itemUnit).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemUnitCreate(err.Error())
	}

	return itemUnit
}

func (repo *itemComponentUnitRepository) Update(itemUnit model.ItemComponentUnit, form form.SettingItemUnitForm) model.ItemComponentUnit {
	itemUnit.Name = form.Name
	itemUnit.Abbreviation = form.Abbreviation
	itemUnit.Type = form.Type
	itemUnit.IsBaseUnit = form.IsBaseUnit
	itemUnit.Conversion = form.Conversion

	err := repo.transaction.Save(&itemUnit).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemUnitUpdate(err.Error())
	}

	return itemUnit
}

func (repo *itemComponentUnitRepository) Delete(itemUnit model.ItemComponentUnit) {
	err := repo.transaction.Delete(&itemUnit).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemUnitDelete(err.Error())
	}
}
