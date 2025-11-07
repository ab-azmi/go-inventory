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

type ItemComponentTypeService interface {
	SetTransaction(tx *gorm.DB)

	Create(form form.SettingForm) model.ItemComponentType
	Update(id string, form form.SettingForm) model.ItemComponentType
	Delete(id string)
}

func NewItemComponentTypeService() ItemComponentTypeService {
	return &itemComponentTypeService{}
}

type itemComponentTypeService struct {
	tx *gorm.DB

	repository repository.ItemComponentTypeRepository
}

func (srv *itemComponentTypeService) SetTransaction(tx *gorm.DB) {
	srv.tx = tx
}

func (srv *itemComponentTypeService) Create(form form.SettingForm) model.ItemComponentType {
	var itemType model.ItemComponentType

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewItemComponentTypeRepository(tx)

		itemType = srv.repository.Create(form)

		itemTypeParser := parser.ItemComponentTypeParser{Object: itemType}

		activity.UseActivity{}.SetReference(itemType).SetParser(&itemTypeParser).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Create new Type: %s", itemType.Name))

		return nil
	})

	return itemType

}

func (srv *itemComponentTypeService) Update(id string, form form.SettingForm) model.ItemComponentType {
	itemType := srv.prepare(id)

	parser := parser.ItemComponentTypeParser{Object: itemType}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		updateActivity := activity.UseActivity{}.SetReference(itemType).SetParser(&parser).SetOldProperty(constant.ACTION_UPDATE)

		srv.repository = repository.NewItemComponentTypeRepository(tx)

		itemType = srv.repository.Update(itemType, form)

		parser.Object = itemType

		updateActivity.SetReference(itemType).SetParser(&parser).SetNewProperty(constant.ACTION_UPDATE).
			Save(fmt.Sprintf("Updated Type: %s", itemType.Name))

		return nil
	})

	return itemType
}

func (srv *itemComponentTypeService) Delete(id string) {
	itemType := srv.prepare(id)

	parser := parser.ItemComponentTypeParser{Object: itemType}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewItemComponentTypeRepository(tx)

		srv.repository.Delete(itemType)

		activity.UseActivity{}.SetReference(itemType).SetParser(&parser).SetOldProperty(constant.ACTION_DELETE).
			Save(fmt.Sprintf("Delete Type: %s", itemType.Name))

		return nil
	})
}

/** --- FUNCTIONS --- */

func (srv *itemComponentTypeService) prepare(id string) model.ItemComponentType {
	srv.repository = repository.NewItemComponentTypeRepository(config.PgSQL)

	uintId, _ := strconv.ParseUint(id, 10, 0)
	itemType := srv.repository.FirstByForm(form.IdNameFilterForm{
		IDs: []uint{uint(uintId)},
	})

	return itemType
}
