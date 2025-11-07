package seeder

import (
	"log"
	"service/internal/pkg/config"
	"service/internal/pkg/constant"
	"service/internal/pkg/core"
	"service/internal/pkg/model"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	"gorm.io/gorm/clause"
)

type ItemWarehouseSeeder struct {
	itemWarehouseCounter uint
}

func (seed *ItemWarehouseSeeder) Seed() {
	seed.truncateTables()

	batchSize := 100

	warehouses, warehouseIds := seed.getWarehouseData(8)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&warehouses, batchSize)
	core.ResetAutoIncrement(model.SettingWarehouse{}.TableName())

	seed.itemWarehouseCounter = 1

	seed.seedItemWarehouseWithoutSN(batchSize, warehouseIds)

	seed.seedItemWarehouseWithSN(batchSize, warehouseIds)
}

func (seed *ItemWarehouseSeeder) truncateTables() {
	config.PgSQL.Exec("SET FOREIGN_KEY_CHECKS = 0")
	result := config.PgSQL.Exec(`TRUNCATE TABLE 
		item_warehouse_stocks,
		item_warehouse_serial_numbers,
		item_warehouse_stock_histories,
		item_warehouse_serial_number_histories,
		item_warehouses
	RESTART IDENTITY CASCADE`)
	if result.Error != nil {
		log.Println(result.Error)
	}
	config.PgSQL.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func (seed *ItemWarehouseSeeder) seedItemWarehouseWithoutSN(batchSize int, warehouseIds []uint) {
	var itemIds []uint
	err := config.PgSQL.Model(&model.Item{}).Where("isTrackSerialNumber", false).Pluck("id", &itemIds).Error
	if err != nil {
		log.Println(err.Error())
	}

	itemWarehouses, itemWarehouseIds := seed.getItemWarehouseData(itemIds, warehouseIds)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&itemWarehouses, batchSize)
	core.ResetAutoIncrement(model.ItemWarehouse{}.TableName())

	stocks, histories := seed.getItemWarehouseStockData(itemWarehouseIds, nil)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&stocks, batchSize)
	core.ResetAutoIncrement(model.ItemWarehouseStock{}.TableName())
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&histories, batchSize)
	core.ResetAutoIncrement(model.ItemWarehouseStockHistory{}.TableName())
}

func (seed ItemWarehouseSeeder) seedItemWarehouseWithSN(batchSize int, warehouseIds []uint) {
	var itemIds []uint
	err := config.PgSQL.Model(&model.Item{}).Where("isTrackSerialNumber", true).Pluck("id", &itemIds).Error
	if err != nil {
		log.Println(err.Error())
	}

	itemWarehouses, itemWarehouseIds := seed.getItemWarehouseData(itemIds, warehouseIds)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&itemWarehouses, batchSize)
	core.ResetAutoIncrement(model.ItemWarehouse{}.TableName())

	numbers, numberHistories := seed.getItemWarehouseSerialNumberData(100, itemWarehouseIds)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&numbers, batchSize)
	core.ResetAutoIncrement(model.ItemWarehouseSerialNumber{}.TableName())
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&numberHistories, batchSize)
	core.ResetAutoIncrement(model.ItemWarehouseSerialNumberHistory{}.TableName())

	stocks, histories := seed.getItemWarehouseStockData(itemWarehouseIds, &numbers)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&stocks, batchSize)
	core.ResetAutoIncrement(model.ItemWarehouseStock{}.TableName())
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&histories, batchSize)
	core.ResetAutoIncrement(model.ItemWarehouseStockHistory{}.TableName())
}

func (seed *ItemWarehouseSeeder) getItemWarehouseData(itemIds []uint, warehouseIds []uint) ([]model.ItemWarehouse, []uint) {
	var itemWarehouses []model.ItemWarehouse
	var ids []uint

	for _, i := range itemIds {
		for _, j := range warehouseIds {
			itemWarehouses = append(itemWarehouses, model.ItemWarehouse{
				BaseModelUUID: xtrememodel.BaseModelUUID{
					ID:        seed.itemWarehouseCounter,
					UUID:      gofakeit.UUID(),
					Timezone:  "Asia/Makassar",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				WarehouseId:  j,
				ItemId:       i,
				OrderType:    gofakeit.RandomString(constant.OrderType{}.OptionCodeNames()),
				SellingPrice: gofakeit.Product().Price,
				IsIncludeTax: gofakeit.Bool(),
				Location:     core.StrPtr(gofakeit.AdverbPlace()),
			})

			ids = append(ids, seed.itemWarehouseCounter)

			seed.itemWarehouseCounter += 1
		}
	}

	return itemWarehouses, ids
}

func (seed *ItemWarehouseSeeder) getWarehouseData(rows uint) ([]model.SettingWarehouse, []uint) {
	var warehouses []model.SettingWarehouse
	var ids []uint

	for i := uint(1); i < rows; i++ {
		warehouses = append(warehouses, model.SettingWarehouse{
			BaseModel: xtrememodel.BaseModel{
				ID:        i,
				Timezone:  "Asia/Makassar",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			BranchOfficeId: gofakeit.RandomUint([]uint{1, 2, 3, 4, 5}),
			Name:           gofakeit.SongName(),
			Address:        core.StrPtr(gofakeit.Address().Address),
		})

		ids = append(ids, i)
	}

	return warehouses, ids
}

func (seed *ItemWarehouseSeeder) getItemWarehouseStockData(itemWarehouseIds []uint, numbers *[]model.ItemWarehouseSerialNumber) ([]model.ItemWarehouseStock, []model.ItemWarehouseStockHistory) {
	var stocks []model.ItemWarehouseStock
	var histories []model.ItemWarehouseStockHistory
	mapNumbers := make(map[uint][]string)
	mapQty := make(map[uint]float64)

	if numbers != nil {
		for _, num := range *numbers {
			mapNumbers[num.ItemWarehouseId] = append(mapNumbers[num.ItemWarehouseId], num.Number)

			if num.Status == constant.SerialNumberType.IN {
				mapQty[num.ItemWarehouseId] += 1
			}
		}
	}

	for _, id := range itemWarehouseIds {
		qty, ok := mapQty[id]
		if !ok {
			qty = float64(0)
		}

		numbers, ok := mapNumbers[id]
		if !ok {
			numbers = []string{}
			qty = gofakeit.Price(1, 20)
		}

		stocks = append(stocks, model.ItemWarehouseStock{
			BaseModel: xtrememodel.BaseModel{
				ID:        id,
				Timezone:  "Asia/Makassar",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			ItemWarehouseId:   id,
			Physical:          qty,
			PhysicalAllocated: 0.00,
		})

		histories = append(histories, model.ItemWarehouseStockHistory{
			BaseModel: xtrememodel.BaseModel{
				ID:        id,
				Timezone:  "Asia/Makassar",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			ItemWarehouseId: id,
			Adjusted:        qty,
			NewQuantity:     qty,
			Description:     &gofakeit.Job().Descriptor,
			SerialNumbers:   (*xtrememodel.ArrayStringColumn)(core.ArrayStringPtr(numbers)),
		})
	}

	return stocks, histories
}

func (seed *ItemWarehouseSeeder) getItemWarehouseSerialNumberData(rows uint, itemWarehouseIds []uint) ([]model.ItemWarehouseSerialNumber, []model.ItemWarehouseSerialNumberHistory) {
	var numbers []model.ItemWarehouseSerialNumber
	var histories []model.ItemWarehouseSerialNumberHistory

	for i := uint(1); i < rows; i++ {
		snType := gofakeit.RandomString(constant.SerialNumberType.OPTION())

		numbers = append(numbers, model.ItemWarehouseSerialNumber{
			BaseModel: xtrememodel.BaseModel{
				ID:        i,
				Timezone:  "Asia/Makassar",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			ItemWarehouseId: gofakeit.RandomUint(itemWarehouseIds),
			Number:          gofakeit.DigitN(10),
			Status:          snType,
		})

		histories = append(histories, model.ItemWarehouseSerialNumberHistory{
			BaseModel: xtrememodel.BaseModel{
				ID:        i,
				Timezone:  "Asia/Makassar",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			SerialNumberId: i,
			Type:           snType,
			Action:         "adjustment",
		})
	}

	return numbers, histories
}
