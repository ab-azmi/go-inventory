package handler

import (
	"net/http"
	SettingRepo "service/internal/item/repository"
	"service/internal/item/service"
	form2 "service/internal/pkg/form"
	parser2 "service/internal/pkg/parser"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"github.com/gorilla/mux"
)

type ItemComponentUnitHandler struct{}

func (hlr *ItemComponentUnitHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := SettingRepo.NewItemComponentUnitRepository()
	units, pagination, _ := repo.Find(r.URL.Query())

	parser := parser2.SettingItemUnitParser{Array: units}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}

func (hlr *ItemComponentUnitHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := form2.SettingItemUnitForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewItemComponentUnitService()

	unit := srv.Create(form)

	parser := parser2.SettingItemUnitParser{Object: unit}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (hlr *ItemComponentUnitHandler) Update(w http.ResponseWriter, r *http.Request) {
	form := form2.SettingItemUnitForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewItemComponentUnitService()

	unit := srv.Update(mux.Vars(r)["id"], form)

	parser := parser2.SettingItemUnitParser{Object: unit}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (hlr *ItemComponentUnitHandler) Delete(w http.ResponseWriter, r *http.Request) {
	srv := service.NewItemComponentUnitService()

	srv.Delete(mux.Vars(r)["id"])

	res := xtremeres.Response{}
	res.Success(w)
}
