package repository

import (
	"database/sql"
	"net/url"
	"service/internal/pkg/config"
	"service/internal/pkg/core"
	"service/internal/pkg/form"
	"service/internal/pkg/model"

	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	"gorm.io/gorm"
)

type ItemRepository interface {
	core.PaginateRepository[model.Item]
}

func NewItemRepository() ItemRepository {
	return itemRepository{}
}

type itemRepository struct {
	Transaction *gorm.DB
}

func (repo itemRepository) Paginate(parameters url.Values) ([]model.Item, interface{}, error) {
	fromDate, toDate := core.SetDateRange(parameters)

	query := config.PgSQL.Where(`"createdAt" BETWEEN ? AND ?`, fromDate, toDate)

	query = repo.prepareFilterForm(form.ItemFilterForm{
		Search:   parameters.Get("search"),
		Preloads: []string{"Type", "Category", "Brand", "Unit"},
	})

	items, pagination, err := xtrememodel.Paginate(query.Order("id DESC"), parameters, model.Item{})
	if err != nil {
		return nil, nil, err
	}

	return items, pagination, nil
}

/** --- Unexported Functions --- */

func (repo itemRepository) prepareFilterForm(form form.ItemFilterForm) *gorm.DB {
	query := config.PgSQL

	if form.IDs != nil && len(form.IDs) > 0 {
		query = query.Where("id IN (?)", form.IDs)
	}

	if form.Search != "" {
		searchVal := "%" + form.Search + "%"
		query = query.Where(`name ILIKE @search OR "SKU" ILIKE @search`, sql.Named("search", searchVal))
	}

	if form.Preloads != nil && len(form.Preloads) > 0 {
		for _, preload := range form.Preloads {
			query = query.Preload(preload)
		}
	}

	return query
}
