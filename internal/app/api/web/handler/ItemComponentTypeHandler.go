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

type ItemComponentTypeHandler struct{}

func (hlr *ItemComponentTypeHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := SettingRepo.NewItemComponentTypeRepository()
	types, pagination, _ := repo.Find(r.URL.Query())

	parser := parser2.ItemComponentTypeParser{Array: types}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}

func (hlr *ItemComponentTypeHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := form2.SettingForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewItemComponentTypeService()

	itemType := srv.Create(form)

	parser := parser2.ItemComponentTypeParser{Object: itemType}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (hlr *ItemComponentTypeHandler) Update(w http.ResponseWriter, r *http.Request) {
	form := form2.SettingForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewItemComponentTypeService()

	itemType := srv.Update(mux.Vars(r)["id"], form)

	parser := parser2.ItemComponentTypeParser{Object: itemType}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (hlr *ItemComponentTypeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	srv := service.NewItemComponentTypeService()

	srv.Delete(mux.Vars(r)["id"])

	res := xtremeres.Response{}
	res.Success(w)
}
