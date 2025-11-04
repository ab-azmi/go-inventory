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

type SettingItemCategoryService interface {
	SetTransaction(tx *gorm.DB)
	SetActivityRepository(repo port.ActivityRepository)

	Create(form form.SettingItemCategoryForm) model.ItemCategory
	Update(id string, form form.SettingItemCategoryForm) model.ItemCategory
	Delete(id string)
}

func NewSettingItemCategoryService() SettingItemCategoryService {
	return &settingItemCategoryService{}
}

type settingItemCategoryService struct {
	tx *gorm.DB

	repository         repository.SettingItemCategoryRepository
	activityRepository port.ActivityRepository
}

func (srv *settingItemCategoryService) SetTransaction(tx *gorm.DB) {
	srv.tx = tx
}

func (srv *settingItemCategoryService) SetActivityRepository(repo port.ActivityRepository) {
	srv.activityRepository = repo
}

func (srv *settingItemCategoryService) Create(form form.SettingItemCategoryForm) model.ItemCategory {
	var itemCategory model.ItemCategory

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewSettingItemCategoryRepository(tx)

		itemCategory = srv.repository.Create(form)

		itemCategoryParser := parser.SettingItemCategoryParser{Object: itemCategory}

		activity.UseActivity{}.SetReference(itemCategory).SetParser(&itemCategoryParser).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Create new Category: %s", itemCategory.Name))

		return nil
	})

	return itemCategory

}

func (srv *settingItemCategoryService) Update(id string, form form.SettingItemCategoryForm) model.ItemCategory {
	itemCategory := srv.prepare(id)

	parser := parser.SettingItemCategoryParser{Object: itemCategory}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		updateActivity := activity.UseActivity{}.SetReference(itemCategory).SetParser(&parser).SetOldProperty(constant.ACTION_UPDATE)

		srv.repository = repository.NewSettingItemCategoryRepository(tx)

		itemCategory = srv.repository.Update(itemCategory, form)

		parser.Object = itemCategory

		updateActivity.SetReference(itemCategory).SetParser(&parser).SetNewProperty(constant.ACTION_UPDATE).
			Save(fmt.Sprintf("Updated Category: %s", itemCategory.Name))

		return nil
	})

	return itemCategory
}

func (srv *settingItemCategoryService) Delete(id string) {
	itemCategory := srv.prepare(id)

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewSettingItemCategoryRepository(tx)

		srv.repository.Delete(itemCategory)

		parser := parser.SettingItemCategoryParser{Object: itemCategory}

		activity.UseActivity{}.SetReference(itemCategory).SetParser(&parser).
			Save(fmt.Sprintf("Delete Category: %s", itemCategory.Name))

		return nil
	})
}

/** --- FUNCTIONS --- */

func (srv *settingItemCategoryService) prepare(id string) model.ItemCategory {
	srv.repository = repository.NewSettingItemCategoryRepository(config.PgSQL)

	uintId, _ := strconv.ParseUint(id, 10, 0)
	itemCategory := srv.repository.FirstById(uint(uintId))

	return itemCategory
}
