package repository

import (
	"database/sql"
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
	core.PaginateRepository[model.ItemComponentUnit]
	core.FirstByFormRepository[model.ItemComponentUnit, form.ItemComponentUnitFilterForm]
	core.FindByFormRepository[model.ItemComponentUnit, form.ItemComponentUnitFilterForm]

	Create(form form.ItemComponentUnitForm) model.ItemComponentUnit
	Update(itemUnit model.ItemComponentUnit, form form.ItemComponentUnitForm) model.ItemComponentUnit
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

func (repo *itemComponentUnitRepository) Paginate(parameter url.Values) ([]model.ItemComponentUnit, interface{}, error) {
	var units []model.ItemComponentUnit

	query := repo.prepareFilterForm(form.ItemComponentUnitFilterForm{
		Search: parameter.Get("search"),
	}).Order("id DESC")

	units, pagination, err := xtrememodel.Paginate(query, parameter, model.ItemComponentUnit{})
	if err != nil {
		return nil, nil, err
	}

	return units, pagination, nil
}

func (repo *itemComponentUnitRepository) FirstByForm(form form.ItemComponentUnitFilterForm) model.ItemComponentUnit {
	var itemUnit model.ItemComponentUnit

	query := repo.prepareFilterForm(form)

	err := query.First(&itemUnit).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentUnitGet(err.Error())
	}

	return itemUnit
}

func (repo *itemComponentUnitRepository) FindByForm(form form.ItemComponentUnitFilterForm) []model.ItemComponentUnit {
	var itemUnits []model.ItemComponentUnit

	query := repo.prepareFilterForm(form)

	err := query.Model(&itemUnits).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentUnitGet(err.Error())
	}

	return itemUnits
}

func (repo *itemComponentUnitRepository) Create(form form.ItemComponentUnitForm) model.ItemComponentUnit {
	itemUnit := model.ItemComponentUnit{
		Name:         form.Name,
		Abbreviation: form.Abbreviation,
		Type:         form.Type,
		IsBaseUnit:   form.IsBaseUnit,
		Conversion:   form.Conversion,
	}

	err := repo.transaction.Create(&itemUnit).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentUnitCreate(err.Error())
	}

	return itemUnit
}

func (repo *itemComponentUnitRepository) Update(itemUnit model.ItemComponentUnit, form form.ItemComponentUnitForm) model.ItemComponentUnit {
	itemUnit.Name = form.Name
	itemUnit.Abbreviation = form.Abbreviation
	itemUnit.Type = form.Type
	itemUnit.IsBaseUnit = form.IsBaseUnit
	itemUnit.Conversion = form.Conversion

	err := repo.transaction.Save(&itemUnit).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentUnitUpdate(err.Error())
	}

	return itemUnit
}

func (repo *itemComponentUnitRepository) Delete(itemUnit model.ItemComponentUnit) {
	err := repo.transaction.Delete(&itemUnit).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentUnitDelete(err.Error())
	}
}

/** --- Unexported Functions --- */

func (repo *itemComponentUnitRepository) prepareFilterForm(form form.ItemComponentUnitFilterForm) *gorm.DB {
	query := config.PgSQL

	if form.IDs != nil && len(form.IDs) > 0 {
		query = query.Where("id IN (?)", form.IDs)
	}

	if form.ID > 0 {
		query = query.Where("id =?", form.ID)
	}

	if form.Search != "" {
		search := "%" + form.Search + "%"
		query = query.Where("name LIKE @search OR abbreviation LIKE @search", sql.Named("search", search))
	}

	return query
}
