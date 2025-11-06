package service

import (
	"fmt"
	"service/internal/item/repository"
	"service/internal/pkg/activity"
	"service/internal/pkg/config"
	"service/internal/pkg/constant"
	"service/internal/pkg/form"
	"service/internal/pkg/model"
	"service/internal/pkg/parser"
	"strconv"

	"gorm.io/gorm"
)

type ItemComponentBrandService interface {
	SetTransaction(tx *gorm.DB)

	Create(form form.SettingForm) model.ItemComponentBrand
	Update(id string, form form.SettingForm) model.ItemComponentBrand
	Delete(id string)
}

func NewItemComponentBrandService() ItemComponentBrandService {
	return &itemComponentBrandService{}
}

type itemComponentBrandService struct {
	tx *gorm.DB

	repository repository.ItemComponentBrandRepository
}

func (srv *itemComponentBrandService) SetTransaction(tx *gorm.DB) {
	srv.tx = tx
}

func (srv *itemComponentBrandService) Create(form form.SettingForm) model.ItemComponentBrand {
	var brand model.ItemComponentBrand

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewItemComponentBrandRepository(tx)

		brand = srv.repository.Create(form)

		brandParser := parser.SettingItemBrandParser{Object: brand}

		activity.UseActivity{}.SetReference(brand).SetParser(&brandParser).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Create new Brand: %s", brand.Name))

		return nil
	})

	return brand

}

func (srv *itemComponentBrandService) Update(id string, form form.SettingForm) model.ItemComponentBrand {
	brand := srv.prepare(id)

	brandParser := parser.SettingItemBrandParser{Object: brand}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		updateActivity := activity.UseActivity{}.SetReference(brand).SetParser(&brandParser).SetOldProperty(constant.ACTION_UPDATE)

		srv.repository = repository.NewItemComponentBrandRepository(tx)

		brand = srv.repository.Update(brand, form)

		brandParser.Object = brand

		updateActivity.SetReference(brand).SetParser(&brandParser).SetNewProperty(constant.ACTION_UPDATE).
			Save(fmt.Sprintf("Updated brand: %s", brand.Name))

		return nil
	})

	return brand
}

func (srv *itemComponentBrandService) Delete(id string) {
	brand := srv.prepare(id)

	parser := parser.SettingItemBrandParser{Object: brand}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewItemComponentBrandRepository(tx)

		srv.repository.Delete(brand)

		activity.UseActivity{}.SetReference(brand).SetParser(&parser).SetOldProperty(constant.ACTION_DELETE).
			Save(fmt.Sprintf("Delete brand: %s", brand.Name))

		return nil
	})
}

/** --- FUNCTIONS --- */

func (srv *itemComponentBrandService) prepare(id string) model.ItemComponentBrand {
	srv.repository = repository.NewItemComponentBrandRepository(config.PgSQL)

	uintId, _ := strconv.ParseUint(id, 10, 0)
	brand := srv.repository.FirstById(uint(uintId))

	return brand
}
