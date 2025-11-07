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

type ItemComponentCategoryRepository interface {
	core.TransactionRepository

	Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemComponentCategory, interface{}, error)
	FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemComponentCategory
	Create(form form.SettingItemCategoryForm) model.ItemComponentCategory
	Update(itemCategory model.ItemComponentCategory, form form.SettingItemCategoryForm) model.ItemComponentCategory
	Delete(itemCategory model.ItemComponentCategory)
}

func NewItemComponentCategoryRepository(args ...*gorm.DB) ItemComponentCategoryRepository {
	repository := itemComponentCategoryRepository{}

	if len(args) > 0 {
		repository.transaction = args[0]
	}

	return &repository
}

type itemComponentCategoryRepository struct {
	transaction *gorm.DB
}

func (repo *itemComponentCategoryRepository) SetTransaction(tx *gorm.DB) {
	repo.transaction = tx
}

func (repo *itemComponentCategoryRepository) Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemComponentCategory, interface{}, error) {
	var categories []model.ItemComponentCategory

	fromDate, toDate := core.SetDateRange(parameter)
	query := config.PgSQL.Where(`"createdAt" BETWEEN ? AND ?`, fromDate, toDate)

	if search := parameter.Get("search"); len(search) > 3 {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query = query.Order("id DESC")

	categories, pagination, err := xtrememodel.Paginate(query, parameter, model.ItemComponentCategory{})
	if err != nil {
		return nil, nil, err
	}

	return categories, pagination, nil
}

func (repo *itemComponentCategoryRepository) FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemComponentCategory {
	var itemCategory model.ItemComponentCategory

	query := config.PgSQL
	if len(args) > 0 {
		query = args[0](query)
	}

	err := query.First(&itemCategory, "id = ?", id).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentCategoryGet(err.Error())
	}

	return itemCategory
}

func (repo *itemComponentCategoryRepository) Create(form form.SettingItemCategoryForm) model.ItemComponentCategory {
	itemCategory := model.ItemComponentCategory{
		Name:      form.Name,
		IsForSale: form.IsForSale,
	}

	err := repo.transaction.Create(&itemCategory).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentCategoryCreate(err.Error())
	}

	return itemCategory
}

func (repo *itemComponentCategoryRepository) Update(itemCategory model.ItemComponentCategory, form form.SettingItemCategoryForm) model.ItemComponentCategory {
	itemCategory.Name = form.Name
	itemCategory.IsForSale = form.IsForSale

	err := repo.transaction.Save(&itemCategory).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentCategoryUpdate(err.Error())
	}

	return itemCategory
}

func (repo *itemComponentCategoryRepository) Delete(itemCategory model.ItemComponentCategory) {
	err := repo.transaction.Delete(&itemCategory).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentCategoryDelete(err.Error())
	}
}
