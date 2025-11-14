package repository

import (
	"fmt"
	"net/url"
	"service/internal/pkg/config"
	"service/internal/pkg/core"
	error2 "service/internal/pkg/error"
	"service/internal/pkg/form"
	"service/internal/pkg/model"

	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	"gorm.io/gorm"
)

type ItemComponentTypeRepository interface {
	core.TransactionRepository
	core.PaginateRepository[model.ItemComponentType]
	core.FirstByFormRepository[model.ItemComponentType, form.ComponentFilterForm]
	core.FindByFormRepository[model.ItemComponentType, form.ComponentFilterForm]

	Create(form form.SettingForm) model.ItemComponentType
	Update(itemType model.ItemComponentType, form form.SettingForm) model.ItemComponentType
	Delete(itemType model.ItemComponentType)
}

func NewItemComponentTypeRepository(args ...*gorm.DB) ItemComponentTypeRepository {
	repository := itemComponentTypeRepository{}

	if len(args) > 0 {
		repository.transaction = args[0]
	}

	return &repository
}

type itemComponentTypeRepository struct {
	transaction *gorm.DB
}

func (repo *itemComponentTypeRepository) SetTransaction(tx *gorm.DB) {
	repo.transaction = tx
}

func (repo *itemComponentTypeRepository) Paginate(parameter url.Values) ([]model.ItemComponentType, interface{}, error) {
	var types []model.ItemComponentType

	query := repo.prepareFilterForm(form.ComponentFilterForm{
		Search:  parameter.Get("search"),
		OrderBy: parameter.Get("orderBy"),
	})

	types, pagination, err := xtrememodel.Paginate(query, parameter, model.ItemComponentType{})
	if err != nil {
		return nil, nil, err
	}

	return types, pagination, nil
}

func (repo *itemComponentTypeRepository) FirstByForm(form form.ComponentFilterForm) model.ItemComponentType {
	var itemType model.ItemComponentType

	query := repo.prepareFilterForm(form)

	err := query.First(&itemType).Error
	if err != nil {
		error2.ErrXtremeItemComponentTypeGet(err.Error())
	}

	return itemType
}

func (repo *itemComponentTypeRepository) FindByForm(form form.ComponentFilterForm) []model.ItemComponentType {
	var itemTypes []model.ItemComponentType

	query := repo.prepareFilterForm(form)

	err := query.Find(&itemTypes).Error
	if err != nil {
		error2.ErrXtremeItemComponentTypeGet(err.Error())
	}

	return itemTypes
}

func (repo *itemComponentTypeRepository) Create(form form.SettingForm) model.ItemComponentType {
	itemType := model.ItemComponentType{
		Name: form.Name,
	}

	err := repo.transaction.Create(&itemType).Error
	if err != nil {
		error2.ErrXtremeItemComponentTypeCreate(err.Error())
	}

	return itemType
}

func (repo *itemComponentTypeRepository) Update(itemType model.ItemComponentType, form form.SettingForm) model.ItemComponentType {
	itemType.Name = form.Name

	err := repo.transaction.Save(&itemType).Error
	if err != nil {
		error2.ErrXtremeItemComponentTypeUpdate(err.Error())
	}

	return itemType
}

func (repo *itemComponentTypeRepository) Delete(itemType model.ItemComponentType) {
	err := repo.transaction.Delete(&itemType).Error
	if err != nil {
		error2.ErrXtremeItemComponentTypeDelete(err.Error())
	}
}

/** --- Unexported Functions --- */

func (repo *itemComponentTypeRepository) prepareFilterForm(form form.ComponentFilterForm) *gorm.DB {
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

	if form.OrderBy != "" {
		field, direction, err := core.GetOrderBy(form.OrderBy)
		if err != nil {
			error2.ErrXtremeItemComponentTypeGet(err.Error())
		}

		query = query.Order(fmt.Sprintf("%s %s", field, direction))
	} else {
		query = query.Order("id DESC")
	}

	return query
}
