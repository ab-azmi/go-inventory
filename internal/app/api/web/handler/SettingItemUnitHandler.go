package handler

import (
	"net/http"
	"service/internal/activity/repository"
	"service/internal/pkg/form"
	"service/internal/pkg/parser"
	SettingRepo "service/internal/setting/repository"
	"service/internal/setting/service"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"github.com/gorilla/mux"
)

type ItemUnitHandler struct{}

func (hlr *ItemUnitHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := SettingRepo.NewSettingItemUnitRepository()
	units, pagination, _ := repo.Find(r.URL.Query())

	parser := parser.SettingItemUnitParser{Array: units}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}

func (hlr *ItemUnitHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := form.SettingItemUnitForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewSettingItemUnitService()
	srv.SetActivityRepository(repository.NewActivityRepository())

	unit := srv.Create(form)

	parser := parser.SettingItemUnitParser{Object: unit}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (hlr *ItemUnitHandler) Update(w http.ResponseWriter, r *http.Request) {
	form := form.SettingItemUnitForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewSettingItemUnitService()
	srv.SetActivityRepository(repository.NewActivityRepository())

	unit := srv.Update(mux.Vars(r)["id"], form)

	parser := parser.SettingItemUnitParser{Object: unit}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (hlr *ItemUnitHandler) Delete(w http.ResponseWriter, r *http.Request) {
	srv := service.NewSettingItemUnitService()
	srv.SetActivityRepository(repository.NewActivityRepository())

	srv.Delete(mux.Vars(r)["id"])

	res := xtremeres.Response{}
	res.Success(w)
}
