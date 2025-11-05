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

type ItemTypeHandler struct{}

func (hlr *ItemTypeHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := SettingRepo.NewSettingItemTypeRepository()
	types, pagination, _ := repo.Find(r.URL.Query())

	parser := parser.SettingItemTypeParser{Array: types}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}

func (hlr *ItemTypeHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := form.SettingForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewSettingItemTypeService()

	itemType := srv.Create(form)

	parser := parser.SettingItemTypeParser{Object: itemType}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (hlr *ItemTypeHandler) Update(w http.ResponseWriter, r *http.Request) {
	form := form.SettingForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewSettingItemTypeService()

	itemType := srv.Update(mux.Vars(r)["id"], form)

	parser := parser.SettingItemTypeParser{Object: itemType}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (hlr *ItemTypeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	srv := service.NewSettingItemTypeService()

	srv.Delete(mux.Vars(r)["id"])

	res := xtremeres.Response{}
	res.Success(w)
}
