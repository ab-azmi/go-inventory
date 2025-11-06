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

type ItemComponentBrandRepository interface {
	core.TransactionRepository

	Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemComponentBrand, interface{}, error)
	FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemComponentBrand
	Create(form form.SettingForm) model.ItemComponentBrand
	Update(brand model.ItemComponentBrand, form form.SettingForm) model.ItemComponentBrand
	Delete(brand model.ItemComponentBrand)
}

func NewItemComponentBrandRepository(args ...*gorm.DB) ItemComponentBrandRepository {
	repository := itemComponentBrandRepository{}

	if len(args) > 0 {
		repository.transaction = args[0]
	}

	return &repository
}

type itemComponentBrandRepository struct {
	transaction *gorm.DB
}

func (repo *itemComponentBrandRepository) SetTransaction(tx *gorm.DB) {
	repo.transaction = tx
}

func (repo *itemComponentBrandRepository) Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemComponentBrand, interface{}, error) {
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

func (repo *itemComponentBrandRepository) FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemComponentBrand {
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

func (repo *itemComponentBrandRepository) Create(form form.SettingForm) model.ItemComponentBrand {
	brand := model.ItemComponentBrand{
		Name: form.Name,
	}

	err := repo.transaction.Create(&brand).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemBrandCreate(err.Error())
	}

	return brand
}

func (repo *itemComponentBrandRepository) Update(brand model.ItemComponentBrand, form form.SettingForm) model.ItemComponentBrand {
	brand.Name = form.Name

	err := repo.transaction.Save(&brand).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemBrandUpdate(err.Error())
	}

	return brand
}

func (repo *itemComponentBrandRepository) Delete(brand model.ItemComponentBrand) {
	err := repo.transaction.Delete(&brand).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemBrandDelete(err.Error())
	}
}
