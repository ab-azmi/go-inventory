package service

import (
	"service/internal/pkg/form"
	"service/internal/pkg/model"
	"service/internal/setting/repository"

	"gorm.io/gorm"
)

type SettingWarehouseService interface {
	SetTransaction(tx *gorm.DB)

	Create(form form.SettingWarehouseForm) model.SettingWarehouse
	Update(id string, form form.SettingWarehouseForm) model.SettingWarehouse
	Delete(id string)
}

func NewSettingWarehouseService() SettingWarehouseService {
	return &settingWarehouseService{}
}

type settingWarehouseService struct {
	tx *gorm.DB

	repository repository.SettingWarehouseRepository
}

func (srv *settingWarehouseService) SetTransaction(tx *gorm.DB) {
	srv.tx = tx
}

func (srv *settingWarehouseService) Create(form form.SettingWarehouseForm) model.SettingWarehouse {
	panic("implement me")
}

func (srv *settingWarehouseService) Update(id string, form form.SettingWarehouseForm) model.SettingWarehouse {
	//TODO implement me
	panic("implement me")
}

func (srv *settingWarehouseService) Delete(id string) {
	//TODO implement me
	panic("implement me")
}
