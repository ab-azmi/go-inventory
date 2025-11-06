package migration

import (
	xtremedb "github.com/globalxtreme/go-core/v2/database"
	"os"
	"service/internal/pkg/config"
	"service/internal/pkg/model"
)

type Item struct{}

func (Item) Reference() string {
	return "Item"
}

func (Item) Tables() []xtremedb.Table {
	owner := os.Getenv("DB_OWNER")

	return []xtremedb.Table{
		{Connection: config.PgSQL, CreateTable: model.ItemUnit{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: model.ItemType{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: model.ItemCategory{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: model.ItemComponentBrand{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: model.Item{}, Owner: owner},
	}
}

func (Item) Columns() []xtremedb.Column {
	return []xtremedb.Column{}
}
