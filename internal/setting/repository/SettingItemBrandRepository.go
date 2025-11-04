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

type SettingItemBrandRepository interface {
	core.TransactionRepository

	Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemBrand, interface{}, error)
	FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemBrand
	Create(form form.SettingForm) model.ItemBrand
	Update(brand model.ItemBrand, form form.SettingForm) model.ItemBrand
	Delete(brand model.ItemBrand)
}

func NewSettingItemBrandRepository(args ...*gorm.DB) SettingItemBrandRepository {
	repository := settingItemBrandRepository{}

	if len(args) > 0 {
		repository.transaction = args[0]
	}

	return &repository
}

type settingItemBrandRepository struct {
	transaction *gorm.DB
}

func (repo *settingItemBrandRepository) SetTransaction(tx *gorm.DB) {
	repo.transaction = tx
}

func (repo *settingItemBrandRepository) Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemBrand, interface{}, error) {
	var brands []model.ItemBrand

	fromDate, toDate := core.SetDateRange(parameter)
	query := config.PgSQL.Where(`"createdAt" BETWEEN ? AND ?`, fromDate, toDate)

	if search := parameter.Get("search"); len(search) > 3 {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query = query.Order("id DESC")

	brands, pagination, err := xtrememodel.Paginate(query, parameter, model.ItemBrand{})
	if err != nil {
		return nil, nil, err
	}

	return brands, pagination, nil
}

func (repo *settingItemBrandRepository) FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemBrand {
	var brand model.ItemBrand

	query := config.PgSQL
	if len(args) > 0 {
		query = args[0](query)
	}

	err := query.First(&brand, "id = ?", id).Error
	if err != nil {
		gxErr.ErrXtremeSettingGet(brand.FeatureName(), err.Error())
	}

	return brand
}

func (repo *settingItemBrandRepository) Create(form form.SettingForm) model.ItemBrand {
	brand := model.ItemBrand{
		Name: form.Name,
	}

	err := repo.transaction.Create(&brand).Error
	if err != nil {
		gxErr.ErrXtremeSettingCreate(brand.FeatureName(), err.Error())
	}

	return brand
}

func (repo *settingItemBrandRepository) Update(brand model.ItemBrand, form form.SettingForm) model.ItemBrand {
	brand.Name = form.Name

	err := repo.transaction.Save(&brand).Error
	if err != nil {
		gxErr.ErrXtremeSettingUpdate(brand.FeatureName(), err.Error())
	}

	return brand
}

func (repo *settingItemBrandRepository) Delete(brand model.ItemBrand) {
	err := repo.transaction.Delete(&brand).Error
	if err != nil {
		gxErr.ErrXtremeSettingDelete(brand.FeatureName(), err.Error())
	}
}
