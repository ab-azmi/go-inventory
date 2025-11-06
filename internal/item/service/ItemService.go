package service

import (
	"service/internal/pkg/form"
	"service/internal/pkg/model"
	"service/internal/pkg/port"
)

type ItemService interface {
	SetSettingItemBrandRepository(repo port.SettingItemBrandRepository)
	SetSettingItemBrandService(service port.SettingItemBrandService)

	Create(form form.SettingForm) model.Item
}

func NewItemService() ItemService {
	return &itemService{}
}

type itemService struct {
	settingItemBrandRepository port.SettingItemBrandRepository
	settingItemBrandService    port.SettingItemBrandService
}

func (srv *itemService) SetSettingItemBrandRepository(repo port.SettingItemBrandRepository) {
	srv.settingItemBrandRepository = repo
}

func (srv *itemService) SetSettingItemBrandService(service port.SettingItemBrandService) {
	srv.settingItemBrandService = service
}

func (srv *itemService) Create(form form.SettingForm) model.Item {
	serviceSecond := NewItemWarehouseService()
	serviceSecond.SetSettingItemBrandService(srv.settingItemBrandService)

	return model.Item{}
}
