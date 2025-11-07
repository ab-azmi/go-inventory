package handler

import (
	"net/http"
	ItemRepo "service/internal/item/repository"
	parser2 "service/internal/pkg/parser"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
)

type ItemHandler struct{}

func (hlr ItemHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := ItemRepo.NewItemRepository()
	items, pagination, _ := repo.Paginate(r.URL.Query())

	parser := parser2.ItemParser{Array: items}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}
