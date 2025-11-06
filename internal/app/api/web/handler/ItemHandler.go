package handler

import (
	"net/http"
	ItemRepo "service/internal/item/repository"
	"service/internal/item/service"
	form2 "service/internal/pkg/form"
	"service/internal/pkg/parser"
	"service/internal/setting/repository"
	service2 "service/internal/setting/service"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
)

type ItemHandler struct{}

func (ItemHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := ItemRepo.NewItemRepository()
	items, pagination, _ := repo.Find(r.URL.Query())

	parser := parser.ItemParser{Array: items}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}

func (ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := form2.SettingForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewItemService()
	srv.SetSettingItemBrandRepository(repository.NewSettingItemBrandRepository())
	srv.SetSettingItemBrandService(service2.NewSettingItemBrandService())

	item := srv.Create(form)

	parser := parser.ItemParser{Object: item}
	response := xtremeres.Response{Object: parser.First()}
	response.Success(w)
}
