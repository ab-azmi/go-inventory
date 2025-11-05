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

type ItemBrandHandler struct{}

func (hlr *ItemBrandHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := SettingRepo.NewSettingItemBrandRepository()
	brands, pagination, _ := repo.Find(r.URL.Query())

	parser := parser.SettingItemBrandParser{Array: brands}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}

func (hlr *ItemBrandHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := form.SettingForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewSettingItemBrandService()

	brand := srv.Create(form)

	parser := parser.SettingItemBrandParser{Object: brand}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (hlr *ItemBrandHandler) Update(w http.ResponseWriter, r *http.Request) {
	form := form.SettingForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewSettingItemBrandService()

	brand := srv.Update(mux.Vars(r)["id"], form)

	parser := parser.SettingItemBrandParser{Object: brand}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (hlr *ItemBrandHandler) Delete(w http.ResponseWriter, r *http.Request) {
	srv := service.NewSettingItemBrandService()

	srv.Delete(mux.Vars(r)["id"])

	res := xtremeres.Response{}
	res.Success(w)
}
