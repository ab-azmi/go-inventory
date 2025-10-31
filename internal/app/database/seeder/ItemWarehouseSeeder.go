package seeder

import (
	"log"
	"service/internal/pkg/config"
	ItemConstant "service/internal/pkg/constant/Item"
	"service/internal/pkg/core"
	ItemModel "service/internal/pkg/model/Item"
	SettingModel "service/internal/pkg/model/Setting"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	"gorm.io/gorm/clause"
)

type ItemWarehouseSeeder struct {
	itemWarehouseCounter uint
}

func (seed *ItemWarehouseSeeder) Seed() {
	batchSize := 100

	warehouses, warehouseIds := seed.getWarehouseData(8)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&warehouses, batchSize)
	core.ResetAutoIncrement(SettingModel.Warehouse{}.TableName())

	seed.itemWarehouseCounter = 1

	seed.seedItemWarehouseWithoutSN(batchSize, warehouseIds)

	seed.seedItemWarehouseWithSN(batchSize, warehouseIds)
}

func (seed ItemWarehouseSeeder) seedItemWarehouseWithoutSN(batchSize int, warehouseIds []uint) {
	var itemIds []uint
	err := config.PgSQL.Model(&ItemModel.Item{}).Where("isTrackSerialNumber", false).Pluck("id", &itemIds).Error
	if err != nil {
		log.Println(err.Error())
	}

	itemWarehouses, itemWarehouseIds := seed.getItemWarehouseData(itemIds, warehouseIds)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&itemWarehouses, batchSize)
	core.ResetAutoIncrement(ItemModel.ItemWarehouse{}.TableName())

	stocks, histories := seed.getItemWarehouseStockData(itemWarehouseIds, nil)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&stocks, batchSize)
	core.ResetAutoIncrement(ItemModel.ItemWarehouseStock{}.TableName())
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&histories, batchSize)
	core.ResetAutoIncrement(ItemModel.ItemWarehouseStockHistory{}.TableName())
}

func (seed ItemWarehouseSeeder) seedItemWarehouseWithSN(batchSize int, warehouseIds []uint) {
	var itemIds []uint
	err := config.PgSQL.Model(&ItemModel.Item{}).Where("isTrackSerialNumber", true).Pluck("id", &itemIds).Error
	if err != nil {
		log.Println(err.Error())
	}

	itemWarehouses, itemWarehouseIds := seed.getItemWarehouseData(itemIds, warehouseIds)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&itemWarehouses, batchSize)
	core.ResetAutoIncrement(ItemModel.ItemWarehouse{}.TableName())

	numbers, numberHistories := seed.getItemWarehouseSerialNumberData(100, itemWarehouseIds)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&numbers, batchSize)
	core.ResetAutoIncrement(ItemModel.ItemWarehouseSerialNumber{}.TableName())
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&numberHistories, batchSize)
	core.ResetAutoIncrement(ItemModel.ItemWarehouseSerialNumberHistory{}.TableName())

	stocks, histories := seed.getItemWarehouseStockData(itemWarehouseIds, &numbers)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&stocks, batchSize)
	core.ResetAutoIncrement(ItemModel.ItemWarehouseStock{}.TableName())
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&histories, batchSize)
	core.ResetAutoIncrement(ItemModel.ItemWarehouseStockHistory{}.TableName())
}

func (seed *ItemWarehouseSeeder) getItemWarehouseData(itemIds []uint, warehouseIds []uint) ([]ItemModel.ItemWarehouse, []uint) {
	var itemWarehouses []ItemModel.ItemWarehouse
	var ids []uint

	for _, i := range itemIds {
		itemWarehouses = append(itemWarehouses, ItemModel.ItemWarehouse{
			BaseModelUUID: xtrememodel.BaseModelUUID{
				ID:        seed.itemWarehouseCounter,
				UUID:      gofakeit.UUID(),
				Timezone:  "Asia/Makassar",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			WarehouseId:  gofakeit.RandomUint(warehouseIds),
			ItemId:       i,
			OrderType:    gofakeit.RandomString(ItemConstant.OrderType.OPTION()),
			SellingPrice: gofakeit.Product().Price,
			IsIncludeTax: gofakeit.Bool(),
			Location:     core.StrPtr(gofakeit.AdverbPlace()),
		})

		ids = append(ids, seed.itemWarehouseCounter)

		seed.itemWarehouseCounter += 1
	}

	return itemWarehouses, ids
}

func (seed *ItemWarehouseSeeder) getWarehouseData(rows uint) ([]SettingModel.Warehouse, []uint) {
	var warehouses []SettingModel.Warehouse
	var ids []uint

	for i := uint(1); i < rows; i++ {
		warehouses = append(warehouses, SettingModel.Warehouse{
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

func (seed *ItemWarehouseSeeder) getItemWarehouseStockData(itemWarehouseIds []uint, numbers *[]ItemModel.ItemWarehouseSerialNumber) ([]ItemModel.ItemWarehouseStock, []ItemModel.ItemWarehouseStockHistory) {
	var stocks []ItemModel.ItemWarehouseStock
	var histories []ItemModel.ItemWarehouseStockHistory
	mapNumbers := make(map[uint][]string)
	mapQty := make(map[uint]float64)

	if numbers != nil {
		for _, num := range *numbers {
			mapNumbers[num.ItemWarehouseId] = append(mapNumbers[num.ItemWarehouseId], num.Number)

			if num.Status == ItemConstant.SerialNumberType.IN {
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

		stocks = append(stocks, ItemModel.ItemWarehouseStock{
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

		histories = append(histories, ItemModel.ItemWarehouseStockHistory{
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

func (seed *ItemWarehouseSeeder) getItemWarehouseSerialNumberData(rows uint, itemWarehouseIds []uint) ([]ItemModel.ItemWarehouseSerialNumber, []ItemModel.ItemWarehouseSerialNumberHistory) {
	var numbers []ItemModel.ItemWarehouseSerialNumber
	var histories []ItemModel.ItemWarehouseSerialNumberHistory

	for i := uint(1); i < rows; i++ {
		snType := gofakeit.RandomString(ItemConstant.SerialNumberType.OPTION())

		numbers = append(numbers, ItemModel.ItemWarehouseSerialNumber{
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

		histories = append(histories, ItemModel.ItemWarehouseSerialNumberHistory{
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
