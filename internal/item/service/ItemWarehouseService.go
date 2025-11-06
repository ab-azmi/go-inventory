package service

import (
	"service/internal/pkg/form"
	"service/internal/pkg/model"
	"service/internal/pkg/port"
)

type ItemWarehouseService interface {
	SetSettingItemBrandRepository(repo port.SettingItemBrandRepository)
	SetSettingItemBrandService(service port.SettingItemBrandService)

	Create(form form.SettingForm) model.Item
}

func NewItemWarehouseService() ItemWarehouseService {
	return &itemWarehouseService{}
}

type itemWarehouseService struct {
	settingItemBrandRepository port.SettingItemBrandRepository
	settingItemBrandService    port.SettingItemBrandService
}

func (srv *itemWarehouseService) SetSettingItemBrandRepository(repo port.SettingItemBrandRepository) {
	srv.settingItemBrandRepository = repo
}

func (srv *itemWarehouseService) SetSettingItemBrandService(service port.SettingItemBrandService) {
	srv.settingItemBrandService = service
}

func (srv *itemWarehouseService) Create(form form.SettingForm) model.Item {
	return model.Item{}
}
