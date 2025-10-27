package migration

import (
	"os"
	"service/internal/pkg/config"
	ItemModel "service/internal/pkg/model/Item"

	xtremedb "github.com/globalxtreme/go-core/v2/database"
)

type Item struct{}

func (Item) Reference() string {
	return "Item"
}

func (Item) Tables() []xtremedb.Table {
	owner := os.Getenv("DB_OWNER")

	return []xtremedb.Table{
		{Connection: config.PgSQL, CreateTable: ItemModel.ItemUnit{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: ItemModel.ItemType{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: ItemModel.ItemCategory{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: ItemModel.ItemBrand{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: ItemModel.Item{}, Owner: owner},
	}
}

func (Item) Columns() []xtremedb.Column {
	return []xtremedb.Column{}
}
