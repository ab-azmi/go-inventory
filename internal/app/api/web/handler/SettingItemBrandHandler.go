package handler

import (
	"net/http"
	"service/internal/pkg/form"
	"service/internal/pkg/model"
	"service/internal/pkg/parser"
	SettingRepo "service/internal/setting/repository"
	SettingService "service/internal/setting/service"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
)

type ItemBrandHandler struct{}

func (hlr *ItemBrandHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := SettingRepo.NewSettingRepository[model.ItemBrand]()
	types, pagination, _ := repo.Find(r.URL.Query())

	parser := parser.SettingParser[model.ItemBrand]{Array: types}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}

func (hlr *ItemBrandHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := &form.SettingForm{}

	srv := SettingService.NewSettingService[model.ItemBrand](nil)
	srv.Create(w, r, form)
}

func (hlr *ItemBrandHandler) Update(w http.ResponseWriter, r *http.Request) {
	form := &form.SettingForm{}

	srv := SettingService.NewSettingService[model.ItemBrand](r)
	srv.Update(w, r, form)
}

func (hlr *ItemBrandHandler) Delete(w http.ResponseWriter, r *http.Request) {
	srv := SettingService.NewSettingService[model.ItemBrand](r)
	srv.Delete(w, r)
}
