package handler

import (
	"net/http"
	SettingForm "service/internal/pkg/form/setting"
	"service/internal/pkg/model"
	"service/internal/pkg/parser"
	SettingRepo "service/internal/setting/repository"
	SettingService "service/internal/setting/service"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
)

type ItemCategoryHandler struct{}

func (hlr *ItemCategoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := SettingRepo.NewSettingRepository[model.ItemCategory]()
	types, pagination, _ := repo.Find(r.URL.Query())

	parser := parser.SettingParser[model.ItemCategory]{Array: types}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}

func (hlr *ItemCategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := &SettingForm.ItemCategoryForm{}

	srv := SettingService.NewSettingService[model.ItemCategory](r)
	srv.Create(w, r, form)
}

func (hlr *ItemCategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	form := &SettingForm.ItemCategoryForm{}

	srv := SettingService.NewSettingService[model.ItemCategory](r)
	srv.Update(w, r, form)
}

func (hlr *ItemCategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	srv := SettingService.NewSettingService[model.ItemCategory](r)
	srv.Delete(w, r)
}
