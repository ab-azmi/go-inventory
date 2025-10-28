package repository

import (
	"database/sql"
	"net/url"
	"service/internal/pkg/config"
	"service/internal/pkg/core"
	ItemModel "service/internal/pkg/model/Item"

	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	"gorm.io/gorm"
)

type ItemRepository interface {
	Find(parameters url.Values) ([]ItemModel.Item, interface{}, error)
}

func NewItemRepository() ItemRepository {
	return itemRepository{}
}

type itemRepository struct {
	Transaction *gorm.DB
}

func (repo itemRepository) Find(parameters url.Values) ([]ItemModel.Item, interface{}, error) {
	query := repo.filterByParam(parameters).Preload("Type").Preload("Category").Preload("Brand").Preload("Unit")

	items, pagination, err := xtrememodel.Paginate(query.Order("id DESC"), parameters, ItemModel.Item{})
	if err != nil {
		return nil, nil, err
	}

	return items, pagination, nil
}

func (repo itemRepository) filterByParam(parameters url.Values) *gorm.DB {
	fromDate, toDate := core.SetDateRange(parameters)

	query := config.PgSQL.Where(`"createdAt" BETWEEN ? AND ?`, fromDate, toDate)

	if searchReq := parameters.Get("search"); len(searchReq) > 3 {
		searchVal := "%" + searchReq + "%"
		query = query.Where(`name ILIKE @search OR "SKU" ILIKE @search`, sql.Named("search", searchVal))
	}

	return query
}
