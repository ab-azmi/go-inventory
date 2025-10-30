package seeder

import (
	"service/internal/pkg/config"
	"service/internal/pkg/core"
	ItemModel "service/internal/pkg/model/Item"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	"gorm.io/gorm/clause"
)

type ItemSeeder struct{}

func (seed *ItemSeeder) Seed() {
	batchSize := 100

	brands, brandIds := seed.getBrandData(10)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&brands, batchSize)
	config.PgSQL.Exec("SELECT setval(pg_get_serial_sequence('item_brands', 'id'), (SELECT MAX(id) FROM item_brands))")

	types, typeIds := seed.getTypeData(10)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&types, batchSize)
	config.PgSQL.Exec("SELECT setval(pg_get_serial_sequence('item_types', 'id'), (SELECT MAX(id) FROM item_types))")

	units, unitIds := seed.getUnitData(5)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&units, batchSize)
	config.PgSQL.Exec("SELECT setval(pg_get_serial_sequence('item_units', 'id'), (SELECT MAX(id) FROM item_units))")

	categories, categoryIds := seed.getCategoryData(5)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&categories, batchSize)
	config.PgSQL.Exec("SELECT setval(pg_get_serial_sequence('item_categories', 'id'), (SELECT MAX(id) FROM item_categories))")

	items := seed.getItemData(20, brandIds, typeIds, unitIds, categoryIds)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&items, batchSize)
	config.PgSQL.Exec("SELECT setval(pg_get_serial_sequence('items', 'id'), (SELECT MAX(id) FROM items))")
}

func (seed *ItemSeeder) getItemData(rows uint, brandIds []uint, typeIds []uint, unitIds []uint, categoryIds []uint) []ItemModel.Item {
	var items []ItemModel.Item

	for i := uint(1); i < rows; i++ {
		items = append(items, ItemModel.Item{
			BaseModelUUID: xtrememodel.BaseModelUUID{
				ID:        i,
				UUID:      gofakeit.UUID(),
				Timezone:  "Asia/Makassar",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			TypeId:        gofakeit.RandomUint(typeIds),
			CategoryId:    gofakeit.RandomUint(categoryIds),
			UnitId:        gofakeit.RandomUint(unitIds),
			BrandId:       core.UintPtr(gofakeit.RandomUint(brandIds)),
			Name:          gofakeit.ProductName(),
			SKU:           gofakeit.DigitN(20),
			IsForSale:     gofakeit.Bool(),
			IsQualified:   gofakeit.Bool(),
			PurchasedCost: gofakeit.Price(5, 70),
			CreatedBy:     core.StrPtr(gofakeit.UUID()),
			CreatedByName: core.StrPtr(gofakeit.FirstName()),
		})
	}

	return items
}

func (seed *ItemSeeder) getBrandData(rows uint) ([]ItemModel.ItemBrand, []uint) {
	var brands []ItemModel.ItemBrand
	var ids []uint

	for i := uint(1); i <= rows; i++ {
		brands = append(brands, ItemModel.ItemBrand{
			Name: gofakeit.LastName(),
			BaseModel: xtrememodel.BaseModel{
				ID:        i,
				Timezone:  "Asia/Makassar",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		})

		ids = append(ids, i)
	}

	return brands, ids
}

func (seed *ItemSeeder) getTypeData(rows uint) ([]ItemModel.ItemType, []uint) {
	var types []ItemModel.ItemType
	var ids []uint

	for i := uint(1); i <= rows; i++ {
		types = append(types, ItemModel.ItemType{
			Name: gofakeit.LastName(),
			BaseModel: xtrememodel.BaseModel{
				ID:        i,
				Timezone:  "Asia/Makassar",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		})

		ids = append(ids, i)
	}

	return types, ids
}

func (seed *ItemSeeder) getUnitData(rows uint) ([]ItemModel.ItemUnit, []uint) {
	var units []ItemModel.ItemUnit
	symbols := []string{"kg", "g", "mg", "l", "ml", "m", "cm", "mm", "pcs"}
	var ids []uint

	for i := uint(1); i <= rows; i++ {
		units = append(units, ItemModel.ItemUnit{
			Name:          gofakeit.RandomString(symbols),
			Abbreviation:  gofakeit.RandomString(symbols),
			Type:          gofakeit.FirstName(),
			IsBaseUnit:    gofakeit.Bool(),
			CreatedBy:     core.StrPtr(gofakeit.UUID()),
			CreatedByName: core.StrPtr(gofakeit.Name()),
			BaseModel: xtrememodel.BaseModel{
				ID:        i,
				Timezone:  "Asia/Makassar",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		})

		ids = append(ids, i)
	}

	return units, ids
}

func (seed *ItemSeeder) getCategoryData(rows uint) ([]ItemModel.ItemCategory, []uint) {
	var categories []ItemModel.ItemCategory
	var ids []uint

	for i := uint(1); i <= rows; i++ {
		categories = append(categories, ItemModel.ItemCategory{
			Name:      gofakeit.LastName(),
			IsForSale: gofakeit.Bool(),
			BaseModel: xtrememodel.BaseModel{
				ID:        i,
				Timezone:  "Asia/Makassar",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		})

		ids = append(ids, i)
	}

	return categories, ids
}
