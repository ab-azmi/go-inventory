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

type ItemComponentBrandHandler struct{}

func (hlr *ItemComponentBrandHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := SettingRepo.NewItemComponentBrandRepository()
	brands, pagination, _ := repo.Find(r.URL.Query())

	parser := parser2.SettingItemBrandParser{Array: brands}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}

func (hlr *ItemComponentBrandHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := form2.SettingForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewItemComponentBrandService()

	brand := srv.Create(form)

	parser := parser2.SettingItemBrandParser{Object: brand}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (hlr *ItemComponentBrandHandler) Update(w http.ResponseWriter, r *http.Request) {
	form := form2.SettingForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewItemComponentBrandService()

	brand := srv.Update(mux.Vars(r)["id"], form)

	parser := parser2.SettingItemBrandParser{Object: brand}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (hlr *ItemComponentBrandHandler) Delete(w http.ResponseWriter, r *http.Request) {
	srv := service.NewItemComponentBrandService()

	srv.Delete(mux.Vars(r)["id"])

	res := xtremeres.Response{}
	res.Success(w)
}
