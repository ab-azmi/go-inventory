package seeder

import (
	"service/internal/pkg/config"
	"service/internal/pkg/core"
	"service/internal/pkg/model"
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
	core.ResetAutoIncrement(model.ItemComponentBrand{}.TableName())

	types, typeIds := seed.getTypeData(10)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&types, batchSize)
	core.ResetAutoIncrement(model.ItemComponentType{}.TableName())

	units, unitIds := seed.getUnitData(5)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&units, batchSize)
	core.ResetAutoIncrement(model.ItemComponentUnit{}.TableName())

	categories, categoryIds := seed.getCategoryData(5)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&categories, batchSize)
	core.ResetAutoIncrement(model.ItemComponentCategory{}.TableName())

	items := seed.getItemData(20, brandIds, typeIds, unitIds, categoryIds)
	config.PgSQL.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&items, batchSize)
	core.ResetAutoIncrement(model.Item{}.TableName())
}

func (seed *ItemSeeder) getItemData(rows uint, brandIds []uint, typeIds []uint, unitIds []uint, categoryIds []uint) []model.Item {
	var items []model.Item

	for i := uint(1); i < rows; i++ {
		items = append(items, model.Item{
			BaseModelUUID: xtrememodel.BaseModelUUID{
				ID:        i,
				UUID:      gofakeit.UUID(),
				Timezone:  "Asia/Makassar",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			TypeId:              gofakeit.RandomUint(typeIds),
			CategoryId:          gofakeit.RandomUint(categoryIds),
			UnitId:              gofakeit.RandomUint(unitIds),
			BrandId:             core.UintPtr(gofakeit.RandomUint(brandIds)),
			Name:                gofakeit.ProductName(),
			SKU:                 gofakeit.DigitN(20),
			IsForSale:           gofakeit.Bool(),
			IsQualified:         gofakeit.Bool(),
			IsTrackSerialNumber: gofakeit.Bool(),
			PurchasedCost:       gofakeit.Price(5, 70),
			CreatedBy:           core.StrPtr(gofakeit.UUID()),
			CreatedByName:       core.StrPtr(gofakeit.FirstName()),
		})
	}

	return items
}

func (seed *ItemSeeder) getBrandData(rows uint) ([]model.ItemComponentBrand, []uint) {
	var brands []model.ItemComponentBrand
	var ids []uint

	for i := uint(1); i <= rows; i++ {
		brands = append(brands, model.ItemComponentBrand{
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

func (seed *ItemSeeder) getTypeData(rows uint) ([]model.ItemComponentType, []uint) {
	var types []model.ItemComponentType
	var ids []uint

	for i := uint(1); i <= rows; i++ {
		types = append(types, model.ItemComponentType{
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

func (seed *ItemSeeder) getUnitData(rows uint) ([]model.ItemComponentUnit, []uint) {
	var units []model.ItemComponentUnit
	symbols := []string{"kg", "g", "mg", "l", "ml", "m", "cm", "mm", "pcs"}
	var ids []uint

	for i := uint(1); i <= rows; i++ {
		units = append(units, model.ItemComponentUnit{
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

func (seed *ItemSeeder) getCategoryData(rows uint) ([]model.ItemComponentCategory, []uint) {
	var categories []model.ItemComponentCategory
	var ids []uint

	for i := uint(1); i <= rows; i++ {
		categories = append(categories, model.ItemComponentCategory{
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
