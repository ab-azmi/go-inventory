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

type ItemComponentUnitService interface {
	SetTransaction(tx *gorm.DB)

	Create(form form.SettingItemUnitForm) model.ItemComponentUnit
	Update(id string, form form.SettingItemUnitForm) model.ItemComponentUnit
	Delete(id string)
}

func NewItemComponentUnitService() ItemComponentUnitService {
	return &itemComponentUnitService{}
}

type itemComponentUnitService struct {
	tx *gorm.DB

	repository repository.ItemComponentUnitRepository
}

func (srv *itemComponentUnitService) SetTransaction(tx *gorm.DB) {
	srv.tx = tx
}

func (srv *itemComponentUnitService) Create(form form.SettingItemUnitForm) model.ItemComponentUnit {
	var itemUnit model.ItemComponentUnit

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewItemComponentUnitRepository(tx)

		itemUnit = srv.repository.Create(form)

		itemUnitParser := parser.SettingItemUnitParser{Object: itemUnit}

		activity.UseActivity{}.SetReference(itemUnit).SetParser(&itemUnitParser).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Create new Unit: %s", itemUnit.Name))

		return nil
	})

	return itemUnit

}

func (srv *itemComponentUnitService) Update(id string, form form.SettingItemUnitForm) model.ItemComponentUnit {
	itemUnit := srv.prepare(id)

	parser := parser.SettingItemUnitParser{Object: itemUnit}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		updateActivity := activity.UseActivity{}.SetReference(itemUnit).SetParser(&parser).SetOldProperty(constant.ACTION_UPDATE)

		srv.repository = repository.NewItemComponentUnitRepository(tx)

		itemUnit = srv.repository.Update(itemUnit, form)

		parser.Object = itemUnit

		updateActivity.SetReference(itemUnit).SetParser(&parser).SetNewProperty(constant.ACTION_UPDATE).
			Save(fmt.Sprintf("Updated Unit: %s", itemUnit.Name))

		return nil
	})

	return itemUnit
}

func (srv *itemComponentUnitService) Delete(id string) {
	itemUnit := srv.prepare(id)

	parser := parser.SettingItemUnitParser{Object: itemUnit}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewItemComponentUnitRepository(tx)

		srv.repository.Delete(itemUnit)

		activity.UseActivity{}.SetReference(itemUnit).SetParser(&parser).SetOldProperty(constant.ACTION_DELETE).
			Save(fmt.Sprintf("Delete Unit: %s", itemUnit.Name))

		return nil
	})
}

/** --- FUNCTIONS --- */

func (srv *itemComponentUnitService) prepare(id string) model.ItemComponentUnit {
	srv.repository = repository.NewItemComponentUnitRepository(config.PgSQL)

	uintId, _ := strconv.ParseUint(id, 10, 0)
	itemUnit := srv.repository.FirstById(uint(uintId))

	return itemUnit
}
