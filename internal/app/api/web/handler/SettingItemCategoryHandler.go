package handler

import (
	"net/http"
	"service/internal/pkg/form"
	"service/internal/pkg/parser"
	SettingRepo "service/internal/setting/repository"
	"service/internal/setting/service"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"github.com/gorilla/mux"
)

type ItemCategoryHandler struct{}

func (hlr *ItemCategoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := SettingRepo.NewSettingItemCategoryRepository()
	categories, pagination, _ := repo.Find(r.URL.Query())

	parser := parser.SettingItemCategoryParser{Array: categories}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}

func (hlr *ItemCategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := form.SettingItemCategoryForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewSettingItemCategoryService()

	category := srv.Create(form)

	parser := parser.SettingItemCategoryParser{Object: category}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (hlr *ItemCategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	form := form.SettingItemCategoryForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewSettingItemCategoryService()

	category := srv.Update(mux.Vars(r)["id"], form)

	parser := parser.SettingItemCategoryParser{Object: category}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (hlr *ItemCategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	srv := service.NewSettingItemCategoryService()

	srv.Delete(mux.Vars(r)["id"])

	res := xtremeres.Response{}
	res.Success(w)
}
