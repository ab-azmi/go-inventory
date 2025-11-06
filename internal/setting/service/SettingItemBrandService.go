package service

import (
	"fmt"
	"service/internal/pkg/activity"
	"service/internal/pkg/config"
	"service/internal/pkg/constant"
	"service/internal/pkg/form"
	"service/internal/pkg/model"
	parser2 "service/internal/pkg/parser"
	"service/internal/setting/repository"
	"strconv"

	"gorm.io/gorm"
)

type SettingItemBrandService interface {
	SetTransaction(tx *gorm.DB)

	Create(form form.SettingForm) model.ItemComponentBrand
	Update(id string, form form.SettingForm) model.ItemComponentBrand
	Delete(id string)
}

func NewSettingItemBrandService() SettingItemBrandService {
	return &settingItemBrandService{}
}

type settingItemBrandService struct {
	tx *gorm.DB

	repository repository.SettingItemBrandRepository
}

func (srv *settingItemBrandService) SetTransaction(tx *gorm.DB) {
	srv.tx = tx
}

func (srv *settingItemBrandService) Create(form form.SettingForm) model.ItemComponentBrand {
	brand := srv.prepare(nil)

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository.SetTransaction(tx)

		brand = srv.repository.Create(form)

		brandParser := parser2.SettingItemBrandParser{Object: brand}

		activity.UseActivity{}.SetReference(brand).SetParser(&brandParser).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Create new Brand: %s", brand.Name))

		return nil
	})

	return brand

}

func (srv *settingItemBrandService) Update(id string, form form.SettingForm) model.ItemComponentBrand {
	brand := srv.prepare(&id)
	parser := parser2.SettingItemBrandParser{Object: brand}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository.SetTransaction(tx)

		updateActivity := activity.UseActivity{}.SetReference(brand).SetParser(&parser).SetOldProperty(constant.ACTION_UPDATE)

		brand = srv.repository.Update(brand, form)

		parser.Object = brand
		updateActivity.SetReference(brand).SetParser(&parser).SetNewProperty(constant.ACTION_UPDATE).
			Save(fmt.Sprintf("Updated brand: %s", brand.Name))

		return nil
	})

	return brand
}

func (srv *settingItemBrandService) Delete(id string) {
	brand := srv.prepare(&id)

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository.SetTransaction(tx)

		srv.repository.Delete(brand)

		activity.UseActivity{}.SetReference(brand).
			Save(fmt.Sprintf("Delete brand: %s", brand.Name))

		return nil
	})
}

/** --- FUNCTIONS --- */

func (srv *settingItemBrandService) prepare(id *string) model.ItemComponentBrand {
	srv.repository = repository.NewSettingItemBrandRepository()

	var brand model.ItemComponentBrand
	if id != nil || *id != "" {
		uintId, _ := strconv.ParseUint(*id, 10, 0)
		brand = srv.repository.FirstById(uint(uintId))
	}

	return brand
}
