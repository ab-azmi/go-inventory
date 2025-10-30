package migration

import (
	"os"
	"service/internal/pkg/config"
	ItemModel "service/internal/pkg/model/Item"
	SettingModel "service/internal/pkg/model/Setting"

	xtremedb "github.com/globalxtreme/go-core/v2/database"
)

type ItemWarehouse_1761804063029332 struct{}

func (ItemWarehouse_1761804063029332) Reference() string {
	return "ItemWarehouse_1761804063029332"
}

func (ItemWarehouse_1761804063029332) Tables() []xtremedb.Table {
	owner := os.Getenv("DB_OWNER")

	return []xtremedb.Table{
		{Connection: config.PgSQL, CreateTable: SettingModel.Warehouse{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: ItemModel.ItemWarehouse{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: ItemModel.ItemWarehouseStock{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: ItemModel.ItemWarehouseStockHistory{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: ItemModel.ItemWarehouseSerialNumber{}, Owner: owner},
		{Connection: config.PgSQL, CreateTable: ItemModel.ItemWarehouseSerialNumberHistory{}, Owner: owner},
	}
}

func (ItemWarehouse_1761804063029332) Columns() []xtremedb.Column {
	return []xtremedb.Column{}
}
