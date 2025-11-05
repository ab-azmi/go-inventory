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

type SettingItemTypeService interface {
	SetTransaction(tx *gorm.DB)
	SetActivityRepository(repo port.ActivityRepository)

	Create(form form.SettingForm) model.ItemType
	Update(id string, form form.SettingForm) model.ItemType
	Delete(id string)
}

func NewSettingItemTypeService() SettingItemTypeService {
	return &settingItemTypeService{}
}

type settingItemTypeService struct {
	tx *gorm.DB

	repository         repository.SettingItemTypeRepository
	activityRepository port.ActivityRepository
}

func (srv *settingItemTypeService) SetTransaction(tx *gorm.DB) {
	srv.tx = tx
}

func (srv *settingItemTypeService) SetActivityRepository(repo port.ActivityRepository) {
	srv.activityRepository = repo
}

func (srv *settingItemTypeService) Create(form form.SettingForm) model.ItemType {
	var itemType model.ItemType

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewSettingItemTypeRepository(tx)

		itemType = srv.repository.Create(form)

		itemTypeParser := parser.SettingItemTypeParser{Object: itemType}

		activity.UseActivity{}.SetReference(itemType).SetParser(&itemTypeParser).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Create new Type: %s", itemType.Name))

		return nil
	})

	return itemType

}

func (srv *settingItemTypeService) Update(id string, form form.SettingForm) model.ItemType {
	itemType := srv.prepare(id)

	parser := parser.SettingItemTypeParser{Object: itemType}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		updateActivity := activity.UseActivity{}.SetReference(itemType).SetParser(&parser).SetOldProperty(constant.ACTION_UPDATE)

		srv.repository = repository.NewSettingItemTypeRepository(tx)

		itemType = srv.repository.Update(itemType, form)

		parser.Object = itemType

		updateActivity.SetReference(itemType).SetParser(&parser).SetNewProperty(constant.ACTION_UPDATE).
			Save(fmt.Sprintf("Updated Type: %s", itemType.Name))

		return nil
	})

	return itemType
}

func (srv *settingItemTypeService) Delete(id string) {
	itemType := srv.prepare(id)

	parser := parser.SettingItemTypeParser{Object: itemType}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewSettingItemTypeRepository(tx)

		srv.repository.Delete(itemType)

		activity.UseActivity{}.SetReference(itemType).SetParser(&parser).SetOldProperty(constant.ACTION_DELETE).
			Save(fmt.Sprintf("Delete Type: %s", itemType.Name))

		return nil
	})
}

/** --- FUNCTIONS --- */

func (srv *settingItemTypeService) prepare(id string) model.ItemType {
	srv.repository = repository.NewSettingItemTypeRepository(config.PgSQL)

	uintId, _ := strconv.ParseUint(id, 10, 0)
	itemType := srv.repository.FirstById(uint(uintId))

	return itemType
}
