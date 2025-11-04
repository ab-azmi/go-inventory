package migration

import (
	"os"
	"service/internal/pkg/config"
	"service/internal/pkg/model"

	xtremedb "github.com/globalxtreme/go-core/v2/database"
)

type ItemWarehouse_1761804063029332 struct{}

func (ItemWarehouse_1761804063029332) Reference() string {
	return "ItemWarehouse_1761804063029332"
}

func (ItemWarehouse_1761804063029332) Tables() []xtremedb.Table {
	owner := os.Getenv("DB_OWNER")

	return []xtremedb.Table{
		{Connection: config.PgSQL, CreateTable: model.SettingWarehouse{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: model.ItemWarehouse{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: model.ItemWarehouseStock{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: model.ItemWarehouseStockHistory{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: model.ItemWarehouseSerialNumber{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: model.ItemWarehouseSerialNumberHistory{}, Owner: owner},
	}
}

func (ItemWarehouse_1761804063029332) Columns() []xtremedb.Column {
	return []xtremedb.Column{}
}
