package repository

import (
	"net/url"
	"service/internal/pkg/config"
	"service/internal/pkg/core"
	gxErr "service/internal/pkg/error"
	"service/internal/pkg/form"
	"service/internal/pkg/form/option"
	"service/internal/pkg/model"

	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	"gorm.io/gorm"
)

type SettingItemBrandRepository interface {
	core.TransactionRepository

	Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemComponentBrand, interface{}, error)
	FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemComponentBrand
	FirstByForm(form form.SettingItemBrandFilterForm) model.ItemComponentBrand
	Create(form form.SettingForm) model.ItemComponentBrand
	Update(brand model.ItemComponentBrand, form form.SettingForm) model.ItemComponentBrand
	Delete(brand model.ItemComponentBrand)
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

func (repo *settingItemBrandRepository) Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemComponentBrand, interface{}, error) {
	var brands []model.ItemComponentBrand

	fromDate, toDate := core.SetDateRange(parameter)
	query := config.PgSQL.Where(`"createdAt" BETWEEN ? AND ?`, fromDate, toDate)

	if search := parameter.Get("search"); len(search) > 3 {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query = query.Order("id DESC")

	brands, pagination, err := xtrememodel.Paginate(query, parameter, model.ItemComponentBrand{})
	if err != nil {
		return nil, nil, err
	}

	return brands, pagination, nil
}

func (repo *settingItemBrandRepository) FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemComponentBrand {
	var brand model.ItemComponentBrand

	query := config.PgSQL
	if len(args) > 0 {
		query = args[0](query)
	}

	err := query.First(&brand, "id = ?", id).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemBrandGet(err.Error())
	}

	return brand
}

func (repo *settingItemBrandRepository) FirstByForm(form form.SettingItemBrandFilterForm) model.ItemComponentBrand {
	query := repo.prepareAndQuery(form)

	var brand model.ItemComponentBrand
	err := query.First(&brand).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemBrandGet(err.Error())
	}

	return brand
}

func (repo *settingItemBrandRepository) Create(form form.SettingForm) model.ItemComponentBrand {
	brand := model.ItemComponentBrand{
		Name: form.Name,
	}

	err := repo.transaction.Create(&brand).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemBrandCreate(err.Error())
	}

	return brand
}

func (repo *settingItemBrandRepository) CreateV2(form form.SettingForm, opt option.SettingItemBrandCreateOption) model.ItemComponentBrand {
	brand := model.ItemComponentBrand{
		Name: form.Name,
	}

	err := repo.transaction.Create(&brand).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemBrandCreate(err.Error())
	}

	return brand
}

func (repo *settingItemBrandRepository) Update(brand model.ItemComponentBrand, form form.SettingForm) model.ItemComponentBrand {
	brand.Name = form.Name

	err := repo.transaction.Save(&brand).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemBrandUpdate(err.Error())
	}

	return brand
}

func (repo *settingItemBrandRepository) Delete(brand model.ItemComponentBrand) {
	err := repo.transaction.Delete(&brand).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemBrandDelete(err.Error())
	}
}

func (repo *settingItemBrandRepository) prepareAndQuery(form form.SettingItemBrandFilterForm) *gorm.DB {
	var query *gorm.DB
	if form.UseTransaction {
		query = repo.transaction
	} else {
		query = config.PgSQL
	}

	if form.ID > 0 {
		query = query.Where("id = ?", form.ID)
	}

	if len(form.Preloads) > 0 {
		for _, preload := range form.Preloads {
			query = query.Preload(preload)
		}
	}

	return query
}
