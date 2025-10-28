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

	brands := seed.getBrandData(10)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&brands, batchSize)

	types := seed.getTypeData(10)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&types, batchSize)

	units := seed.getUnitData(5)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&units, batchSize)

	categories := seed.getCategoryData(5)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&categories, batchSize)

	items := seed.getItemData(20)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&items, batchSize)
}

func (seed *ItemSeeder) getItemData(rows uint) []ItemModel.Item {
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
			TypeId:        gofakeit.RandomUint([]uint{1, 2, 3, 4, 5}),
			CategoryId:    gofakeit.RandomUint([]uint{1, 2, 3, 4, 5}),
			UnitId:        gofakeit.RandomUint([]uint{1, 2, 3, 4, 5}),
			BrandId:       core.UintPtr(gofakeit.RandomUint([]uint{1, 2, 3, 4, 5})),
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

func (seed *ItemSeeder) getBrandData(rows uint) []ItemModel.ItemBrand {
	var brands []ItemModel.ItemBrand

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
	}

	return brands
}

func (seed *ItemSeeder) getTypeData(rows uint) []ItemModel.ItemType {
	var types []ItemModel.ItemType

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
	}

	return types
}

func (seed *ItemSeeder) getUnitData(rows uint) []ItemModel.ItemUnit {
	var units []ItemModel.ItemUnit
	symbols := []string{"kg", "g", "mg", "l", "ml", "m", "cm", "mm", "pcs"}

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
	}

	return units
}

func (seed *ItemSeeder) getCategoryData(rows uint) []ItemModel.ItemCategory {
	var categories []ItemModel.ItemCategory

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
	}

	return categories
}
