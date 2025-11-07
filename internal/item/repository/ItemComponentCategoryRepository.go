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
	core.PaginateRepository[model.ItemComponentCategory]
	core.FirstByFormRepository[model.ItemComponentCategory, form.ItemComponentCategoryFilterForm]
	core.FindByFormRepository[model.ItemComponentCategory, form.ItemComponentCategoryFilterForm]

	Create(form form.ItemComponentCategoryForm) model.ItemComponentCategory
	Update(itemCategory model.ItemComponentCategory, form form.ItemComponentCategoryForm) model.ItemComponentCategory
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

func (repo *itemComponentCategoryRepository) Paginate(parameter url.Values) ([]model.ItemComponentCategory, interface{}, error) {
	var categories []model.ItemComponentCategory

	fromDate, toDate := core.SetDateRange(parameter)
	query := config.PgSQL.Where(`"createdAt" BETWEEN ? AND ?`, fromDate, toDate)

	query = repo.prepare(form.ItemComponentCategoryFilterForm{
		Search:    parameter.Get("search"),
		IsForSale: parameter.Get("isForSale") == "true",
	}).Order("id DESC")

	categories, pagination, err := xtrememodel.Paginate(query, parameter, model.ItemComponentCategory{})
	if err != nil {
		return nil, nil, err
	}

	return categories, pagination, nil
}

func (repo *itemComponentCategoryRepository) FirstByForm(form form.ItemComponentCategoryFilterForm) model.ItemComponentCategory {
	var itemCategory model.ItemComponentCategory

	query := config.PgSQL
	query = repo.prepare(form)

	err := query.Model(&itemCategory).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentCategoryGet(err.Error())
	}

	return itemCategory
}

func (repo *itemComponentCategoryRepository) FindByForm(form form.ItemComponentCategoryFilterForm) []model.ItemComponentCategory {
	var categories []model.ItemComponentCategory

	query := config.PgSQL
	query = repo.prepare(form)

	err := query.Model(&categories).Error
	if err != nil {
		gxErr.ErrXtremeItemComponentCategoryGet(err.Error())
	}

	return categories
}

func (repo *itemComponentCategoryRepository) Create(form form.ItemComponentCategoryForm) model.ItemComponentCategory {
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

func (repo *itemComponentCategoryRepository) Update(itemCategory model.ItemComponentCategory, form form.ItemComponentCategoryForm) model.ItemComponentCategory {
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

/** --- Unexported Functions --- */

func (repo *itemComponentCategoryRepository) prepare(form form.ItemComponentCategoryFilterForm) *gorm.DB {
	query := config.PgSQL

	if form.IDs != nil && len(form.IDs) > 0 {
		query = query.Where("id id IN (?)", form.IDs)
	}

	if form.Search != "" {
		query = query.Where("name LIKE ?", "%"+form.Search+"%")
	}

	return query
}
