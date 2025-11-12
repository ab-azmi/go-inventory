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

type ItemComponentTypeRepository interface {
	core.TransactionRepository
	core.PaginateRepository[model.ItemComponentType]
	core.FirstByFormRepository[model.ItemComponentType, form.IdNameFilterForm]
	core.FindByFormRepository[model.ItemComponentType, form.IdNameFilterForm]

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

	query := repo.prepareFilterForm(form.IdNameFilterForm{
		Search: parameter.Get("search"),
	}).Order("id DESC")

	types, pagination, err := xtrememodel.Paginate(query, parameter, model.ItemComponentType{})
	if err != nil {
		return nil, nil, err
	}

	return types, pagination, nil
}

func (repo *itemComponentTypeRepository) FirstByForm(form form.IdNameFilterForm) model.ItemComponentType {
	var itemType model.ItemComponentType

	query := config.PgSQL

	query = repo.prepareFilterForm(form)

	err := query.First(&itemType).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentTypeGet(err.Error())
	}

	return itemType
}

func (repo *itemComponentTypeRepository) FindByForm(form form.IdNameFilterForm) []model.ItemComponentType {
	var itemTypes []model.ItemComponentType

	query := config.PgSQL

	query = repo.prepareFilterForm(form).Order("id DESC")

	err := query.Model(&itemTypes).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentTypeGet(err.Error())
	}

	return itemTypes
}

func (repo *itemComponentTypeRepository) Create(form form.SettingForm) model.ItemComponentType {
	itemType := model.ItemComponentType{
		Name: form.Name,
	}

	err := repo.transaction.Create(&itemType).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentTypeCreate(err.Error())
	}

	return itemType
}

func (repo *itemComponentTypeRepository) Update(itemType model.ItemComponentType, form form.SettingForm) model.ItemComponentType {
	itemType.Name = form.Name

	err := repo.transaction.Save(&itemType).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentTypeUpdate(err.Error())
	}

	return itemType
}

func (repo *itemComponentTypeRepository) Delete(itemType model.ItemComponentType) {
	err := repo.transaction.Delete(&itemType).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentTypeDelete(err.Error())
	}
}

/** --- Unexported Functions --- */

func (repo *itemComponentTypeRepository) prepareFilterForm(form form.IdNameFilterForm) *gorm.DB {
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
