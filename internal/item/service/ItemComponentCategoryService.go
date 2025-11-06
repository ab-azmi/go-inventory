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

type ItemComponentCategoryService interface {
	SetTransaction(tx *gorm.DB)

	Create(form form.SettingItemCategoryForm) model.ItemComponentCategory
	Update(id string, form form.SettingItemCategoryForm) model.ItemComponentCategory
	Delete(id string)
}

func NewItemComponentCategoryService() ItemComponentCategoryService {
	return &itemComponentCategoryService{}
}

type itemComponentCategoryService struct {
	tx *gorm.DB

	repository repository.ItemComponentCategoryRepository
}

func (srv *itemComponentCategoryService) SetTransaction(tx *gorm.DB) {
	srv.tx = tx
}

func (srv *itemComponentCategoryService) Create(form form.SettingItemCategoryForm) model.ItemComponentCategory {
	var itemCategory model.ItemComponentCategory

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewItemComponentCategoryRepository(tx)

		itemCategory = srv.repository.Create(form)

		itemCategoryParser := parser.SettingItemCategoryParser{Object: itemCategory}

		activity.UseActivity{}.SetReference(itemCategory).SetParser(&itemCategoryParser).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Create new Category: %s", itemCategory.Name))

		return nil
	})

	return itemCategory

}

func (srv *itemComponentCategoryService) Update(id string, form form.SettingItemCategoryForm) model.ItemComponentCategory {
	itemCategory := srv.prepare(id)

	parser := parser.SettingItemCategoryParser{Object: itemCategory}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		updateActivity := activity.UseActivity{}.SetReference(itemCategory).SetParser(&parser).SetOldProperty(constant.ACTION_UPDATE)

		srv.repository = repository.NewItemComponentCategoryRepository(tx)

		itemCategory = srv.repository.Update(itemCategory, form)

		parser.Object = itemCategory

		updateActivity.SetReference(itemCategory).SetParser(&parser).SetNewProperty(constant.ACTION_UPDATE).
			Save(fmt.Sprintf("Updated Category: %s", itemCategory.Name))

		return nil
	})

	return itemCategory
}

func (srv *itemComponentCategoryService) Delete(id string) {
	itemCategory := srv.prepare(id)

	parser := parser.SettingItemCategoryParser{Object: itemCategory}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewItemComponentCategoryRepository(tx)

		srv.repository.Delete(itemCategory)

		activity.UseActivity{}.SetReference(itemCategory).SetParser(&parser).SetOldProperty(constant.ACTION_DELETE).
			Save(fmt.Sprintf("Delete Category: %s", itemCategory.Name))

		return nil
	})
}

/** --- FUNCTIONS --- */

func (srv *itemComponentCategoryService) prepare(id string) model.ItemComponentCategory {
	srv.repository = repository.NewItemComponentCategoryRepository(config.PgSQL)

	uintId, _ := strconv.ParseUint(id, 10, 0)
	itemCategory := srv.repository.FirstById(uint(uintId))

	return itemCategory
}
