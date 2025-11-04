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

type SettingItemUnitService interface {
	SetTransaction(tx *gorm.DB)
	SetActivityRepository(repo port.ActivityRepository)

	Create(form form.SettingItemUnitForm) model.ItemUnit
	Update(id string, form form.SettingItemUnitForm) model.ItemUnit
	Delete(id string)
}

func NewSettingItemUnitService() SettingItemUnitService {
	return &settingItemUnitService{}
}

type settingItemUnitService struct {
	tx *gorm.DB

	repository         repository.SettingItemUnitRepository
	activityRepository port.ActivityRepository
}

func (srv *settingItemUnitService) SetTransaction(tx *gorm.DB) {
	srv.tx = tx
}

func (srv *settingItemUnitService) SetActivityRepository(repo port.ActivityRepository) {
	srv.activityRepository = repo
}

func (srv *settingItemUnitService) Create(form form.SettingItemUnitForm) model.ItemUnit {
	var itemUnit model.ItemUnit

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewSettingItemUnitRepository(tx)

		itemUnit = srv.repository.Create(form)

		itemUnitParser := parser.SettingItemUnitParser{Object: itemUnit}

		activity.UseActivity{}.SetReference(itemUnit).SetParser(&itemUnitParser).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Create new Unit: %s", itemUnit.Name))

		return nil
	})

	return itemUnit

}

func (srv *settingItemUnitService) Update(id string, form form.SettingItemUnitForm) model.ItemUnit {
	itemUnit := srv.prepare(id)

	parser := parser.SettingItemUnitParser{Object: itemUnit}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		updateActivity := activity.UseActivity{}.SetReference(itemUnit).SetParser(&parser).SetOldProperty(constant.ACTION_UPDATE)

		srv.repository = repository.NewSettingItemUnitRepository(tx)

		itemUnit = srv.repository.Update(itemUnit, form)

		parser.Object = itemUnit

		updateActivity.SetReference(itemUnit).SetParser(&parser).SetNewProperty(constant.ACTION_UPDATE).
			Save(fmt.Sprintf("Updated Unit: %s", itemUnit.Name))

		return nil
	})

	return itemUnit
}

func (srv *settingItemUnitService) Delete(id string) {
	itemUnit := srv.prepare(id)

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewSettingItemUnitRepository(tx)

		srv.repository.Delete(itemUnit)

		parser := parser.SettingItemUnitParser{Object: itemUnit}

		activity.UseActivity{}.SetReference(itemUnit).SetParser(&parser).
			Save(fmt.Sprintf("Delete Unit: %s", itemUnit.Name))

		return nil
	})
}

/** --- FUNCTIONS --- */

func (srv *settingItemUnitService) prepare(id string) model.ItemUnit {
	srv.repository = repository.NewSettingItemUnitRepository(config.PgSQL)

	uintId, _ := strconv.ParseUint(id, 10, 0)
	itemUnit := srv.repository.FirstById(uint(uintId))

	return itemUnit
}
