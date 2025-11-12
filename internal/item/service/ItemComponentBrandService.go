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

	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
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
	brand := srv.prepare(nil)

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository.SetTransaction(tx)

		brand = srv.repository.Create(form)

		brandParser := parser.ItemComponentBrandParser{Object: brand}

		activity.UseActivity{}.SetReference(brand).SetParser(&brandParser).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Create new Brand: %s", brand.Name))

		return nil
	})

	return brand
}

func (srv *itemComponentBrandService) Update(id string, form form.SettingForm) model.ItemComponentBrand {
	brand := srv.prepare(&id)

	brandParser := parser.ItemComponentBrandParser{Object: brand}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		updateActivity := activity.UseActivity{}.SetReference(brand).SetParser(&brandParser).SetOldProperty(constant.ACTION_UPDATE)

		srv.repository.SetTransaction(tx)

		brand = srv.repository.Update(brand, form)

		brandParser.Object = brand

		updateActivity.SetReference(brand).SetParser(&brandParser).SetNewProperty(constant.ACTION_UPDATE).
			Save(fmt.Sprintf("Updated brand: %s", brand.Name))

		return nil
	})

	return brand
}

func (srv *itemComponentBrandService) Delete(id string) {
	brand := srv.prepare(&id)

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository.SetTransaction(tx)

		srv.repository.Delete(brand)

		activity.UseActivity{}.SetReference(brand).Save(fmt.Sprintf("Delete brand: %s", brand.Name))

		return nil
	})
}

/** --- FUNCTIONS --- */

func (srv *itemComponentBrandService) prepare(id *string) model.ItemComponentBrand {
	srv.repository = repository.NewItemComponentBrandRepository(config.PgSQL)
	var brand model.ItemComponentBrand

	if id != nil {
		uintId := xtremepkg.ToInt(*id)
		brand = srv.repository.FirstByForm(form.IdNameFilterForm{
			ID: uint(uintId),
		})
	}

	return brand
}
