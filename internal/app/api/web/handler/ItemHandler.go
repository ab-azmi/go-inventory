package handler

import (
	"net/http"
	ItemRepo "service/internal/item/repository"
	"service/internal/pkg/parser"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
)

type ItemHandler struct{}

func (hlr ItemHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := ItemRepo.NewItemRepository()
	items, pagination, _ := repo.Find(r.URL.Query())

	parser := parser.ItemParser{Array: items}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}
