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

type ItemComponentCategoryHandler struct{}

func (hlr *ItemComponentCategoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := SettingRepo.NewItemComponentCategoryRepository()
	categories, pagination, _ := repo.Find(r.URL.Query())

	parser := parser2.ItemComponentCategoryParser{Array: categories}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}

func (hlr *ItemComponentCategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := form2.ItemComponentCategoryForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewItemComponentCategoryService()

	category := srv.Create(form)

	parser := parser2.ItemComponentCategoryParser{Object: category}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (hlr *ItemComponentCategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	form := form2.ItemComponentCategoryForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewItemComponentCategoryService()

	category := srv.Update(mux.Vars(r)["id"], form)

	parser := parser2.ItemComponentCategoryParser{Object: category}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (hlr *ItemComponentCategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	srv := service.NewItemComponentCategoryService()

	srv.Delete(mux.Vars(r)["id"])

	res := xtremeres.Response{}
	res.Success(w)
}
