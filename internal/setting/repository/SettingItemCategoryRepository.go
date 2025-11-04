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

type SettingItemCategoryRepository interface {
	core.TransactionRepository

	Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemCategory, interface{}, error)
	FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemCategory
	Create(form form.SettingItemCategoryForm) model.ItemCategory
	Update(itemCategory model.ItemCategory, form form.SettingItemCategoryForm) model.ItemCategory
	Delete(itemCategory model.ItemCategory)
}

func NewSettingItemCategoryRepository(args ...*gorm.DB) SettingItemCategoryRepository {
	repository := settingItemCategoryRepository{}

	if len(args) > 0 {
		repository.transaction = args[0]
	}

	return &repository
}

type settingItemCategoryRepository struct {
	transaction *gorm.DB
}

func (repo *settingItemCategoryRepository) SetTransaction(tx *gorm.DB) {
	repo.transaction = tx
}

func (repo *settingItemCategoryRepository) Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemCategory, interface{}, error) {
	var categories []model.ItemCategory

	fromDate, toDate := core.SetDateRange(parameter)
	query := config.PgSQL.Where(`"createdAt" BETWEEN ? AND ?`, fromDate, toDate)

	if search := parameter.Get("search"); len(search) > 3 {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query = query.Order("id DESC")

	categories, pagination, err := xtrememodel.Paginate(query, parameter, model.ItemCategory{})
	if err != nil {
		return nil, nil, err
	}

	return categories, pagination, nil
}

func (repo *settingItemCategoryRepository) FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemCategory {
	var itemCategory model.ItemCategory

	query := config.PgSQL
	if len(args) > 0 {
		query = args[0](query)
	}

	err := query.First(&itemCategory, "id = ?", id).Error
	if err != nil {
		gxErr.ErrXtremeSettingGet(itemCategory.FeatureName(), err.Error())
	}

	return itemCategory
}

func (repo *settingItemCategoryRepository) Create(form form.SettingItemCategoryForm) model.ItemCategory {
	itemCategory := model.ItemCategory{
		Name:      form.Name,
		IsForSale: form.IsForSale,
	}

	err := repo.transaction.Create(&itemCategory).Error
	if err != nil {
		gxErr.ErrXtremeSettingCreate(itemCategory.FeatureName(), err.Error())
	}

	return itemCategory
}

func (repo *settingItemCategoryRepository) Update(itemCategory model.ItemCategory, form form.SettingItemCategoryForm) model.ItemCategory {
	itemCategory.Name = form.Name
	itemCategory.IsForSale = form.IsForSale

	err := repo.transaction.Save(&itemCategory).Error
	if err != nil {
		gxErr.ErrXtremeSettingUpdate(itemCategory.FeatureName(), err.Error())
	}

	return itemCategory
}

func (repo *settingItemCategoryRepository) Delete(itemCategory model.ItemCategory) {
	err := repo.transaction.Delete(&itemCategory).Error
	if err != nil {
		gxErr.ErrXtremeSettingDelete(itemCategory.FeatureName(), err.Error())
	}
}
