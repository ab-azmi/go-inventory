package service

import (
	"fmt"
	"service/internal/pkg/activity"
	"service/internal/pkg/config"
	"service/internal/pkg/constant"
	"service/internal/pkg/form"
	"service/internal/pkg/model"
	"service/internal/pkg/parser"
	"service/internal/pkg/port"
	"service/internal/setting/repository"
	"strconv"

	"gorm.io/gorm"
)

type SettingItemBrandService interface {
	SetTransaction(tx *gorm.DB)
	SetActivityRepository(repo port.ActivityRepository)

	Create(form form.SettingForm) model.ItemBrand
	Update(id string, form form.SettingForm) model.ItemBrand
	Delete(id string)
}

func NewSettingItemBrandService() SettingItemBrandService {
	return &settingItemBrandService{}
}

type settingItemBrandService struct {
	tx *gorm.DB

	repository         repository.SettingItemBrandRepository
	activityRepository port.ActivityRepository
}

func (srv *settingItemBrandService) SetTransaction(tx *gorm.DB) {
	srv.tx = tx
}

func (srv *settingItemBrandService) SetActivityRepository(repo port.ActivityRepository) {
	srv.activityRepository = repo
}

func (srv *settingItemBrandService) Create(form form.SettingForm) model.ItemBrand {
	var brand model.ItemBrand

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewSettingItemBrandRepository(tx)

		brand = srv.repository.Create(form)

		brandParser := parser.SettingItemBrandParser{Object: brand}

		activity.UseActivity{}.SetReference(brand).SetParser(&brandParser).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Create new Brand: %s", brand.Name))

		return nil
	})

	return brand

}

func (srv *settingItemBrandService) Update(id string, form form.SettingForm) model.ItemBrand {
	brand := srv.prepare(id)

	parser := parser.SettingItemBrandParser{Object: brand}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		updateActivity := activity.UseActivity{}.SetReference(brand).SetParser(&parser).SetOldProperty(constant.ACTION_UPDATE)

		srv.repository = repository.NewSettingItemBrandRepository(tx)

		brand = srv.repository.Update(brand, form)

		parser.Object = brand

		updateActivity.SetReference(brand).SetParser(&parser).SetNewProperty(constant.ACTION_UPDATE).
			Save(fmt.Sprintf("Updated brand: %s", brand.Name))

		return nil
	})

	return brand
}

func (srv *settingItemBrandService) Delete(id string) {
	brand := srv.prepare(id)

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewSettingItemBrandRepository(tx)

		srv.repository.Delete(brand)

		parser := parser.SettingItemBrandParser{Object: brand}

		activity.UseActivity{}.SetReference(brand).SetParser(&parser).
			Save(fmt.Sprintf("Delete brand: %s", brand.Name))

		return nil
	})
}

/** --- FUNCTIONS --- */

func (srv *settingItemBrandService) prepare(id string) model.ItemBrand {
	srv.repository = repository.NewSettingItemBrandRepository(config.PgSQL)

	uintId, _ := strconv.ParseUint(id, 10, 0)
	brand := srv.repository.FirstById(uint(uintId))

	return brand
}
