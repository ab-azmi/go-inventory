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

type ItemUnitHandler struct{}

func (hlr *ItemUnitHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := SettingRepo.NewSettingRepository[model.ItemUnit]()
	types, pagination, _ := repo.Find(r.URL.Query())

	parser := parser.SettingParser[model.ItemUnit]{Array: types}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}

func (hlr *ItemUnitHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := &form.SettingItemUnitForm{}

	srv := SettingService.NewSettingService[model.ItemUnit](r)
	srv.Create(w, r, form)
}

func (hlr *ItemUnitHandler) Update(w http.ResponseWriter, r *http.Request) {
	form := &form.SettingItemUnitForm{}

	srv := SettingService.NewSettingService[model.ItemUnit](r)
	srv.Update(w, r, form)
}

func (hlr *ItemUnitHandler) Delete(w http.ResponseWriter, r *http.Request) {
	srv := SettingService.NewSettingService[model.ItemUnit](r)
	srv.Delete(w, r)
}
