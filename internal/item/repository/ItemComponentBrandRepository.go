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
	core.PaginateRepository[model.ItemComponentBrand]
	core.FirstByFormRepository[model.ItemComponentBrand, form.IdNameFilterForm]
	core.FindByFormRepository[model.ItemComponentBrand, form.IdNameFilterForm]

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

func (repo *itemComponentBrandRepository) Paginate(parameter url.Values) ([]model.ItemComponentBrand, interface{}, error) {
	var brands []model.ItemComponentBrand

	query := repo.prepareFilterForm(form.IdNameFilterForm{
		Search: parameter.Get("search"),
	}).Order("id DESC")

	brands, pagination, err := xtrememodel.Paginate(query, parameter, model.ItemComponentBrand{})
	if err != nil {
		return nil, nil, err
	}

	return brands, pagination, nil
}

func (repo *itemComponentBrandRepository) FirstByForm(form form.IdNameFilterForm) model.ItemComponentBrand {
	query := repo.prepareFilterForm(form)

	var brand model.ItemComponentBrand
	err := query.First(&brand).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentBrandGet(err.Error())
	}

	return brand
}

func (repo *itemComponentBrandRepository) FindByForm(form form.IdNameFilterForm) []model.ItemComponentBrand {
	query := repo.prepareFilterForm(form)

	var brands []model.ItemComponentBrand

	query = query.Order("id DESC")

	err := query.Model(&brands).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentBrandGet(err.Error())
	}

	return brands
}

func (repo *itemComponentBrandRepository) Create(form form.SettingForm) model.ItemComponentBrand {
	brand := model.ItemComponentBrand{
		Name: form.Name,
	}

	err := repo.transaction.Create(&brand).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentBrandCreate(err.Error())
	}

	return brand
}

func (repo *itemComponentBrandRepository) Update(brand model.ItemComponentBrand, form form.SettingForm) model.ItemComponentBrand {
	brand.Name = form.Name

	err := repo.transaction.Save(&brand).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentBrandUpdate(err.Error())
	}

	return brand
}

func (repo *itemComponentBrandRepository) Delete(brand model.ItemComponentBrand) {
	err := repo.transaction.Delete(&brand).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentBrandDelete(err.Error())
	}
}

/** --- Unexported Functions --- */

func (repo *itemComponentBrandRepository) prepareFilterForm(form form.IdNameFilterForm) *gorm.DB {
	query := config.PgSQL

	if form.IDs != nil && len(form.IDs) > 0 {
		query = query.Where("id IN (?)", form.IDs)
	}

	if form.Search != "" {
		search := "%" + form.Search + "%"
		query = query.Where("name LIKE ?", search)
	}

	if form.Limit > 0 {
		query = query.Limit(int(form.Limit))
	}

	return query
}
