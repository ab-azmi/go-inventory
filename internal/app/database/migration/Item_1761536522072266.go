package migration

import (
	"os"
	"service/internal/pkg/config"
	"service/internal/pkg/model"

	xtremedb "github.com/globalxtreme/go-core/v2/database"
)

type Item struct{}

func (Item) Reference() string {
	return "Item"
}

func (Item) Tables() []xtremedb.Table {
	owner := os.Getenv("DB_OWNER")

	return []xtremedb.Table{
		{Connection: config.PgSQL, CreateTable: model.ItemComponentUnit{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: model.ItemComponentType{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: model.ItemComponentCategory{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: model.ItemComponentBrand{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: model.Item{}, Owner: owner},
	}
}

func (Item) Columns() []xtremedb.Column {
	return []xtremedb.Column{}
}
