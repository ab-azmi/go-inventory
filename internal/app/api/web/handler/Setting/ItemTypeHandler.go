package setting

import (
	"net/http"
	SettingForm "service/internal/pkg/form/setting"
	ItemModel "service/internal/pkg/model/Item"
	SettingParser "service/internal/pkg/parser/Setting"
	SettingRepo "service/internal/setting/repository"
	SettingService "service/internal/setting/service"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
)

type ItemTypeHandler struct{}

func (hlr *ItemTypeHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := SettingRepo.NewSettingRepository[ItemModel.ItemType]()
	types, pagination, _ := repo.Find(r.URL.Query())

	parser := SettingParser.SettingParser[ItemModel.ItemType]{Array: types}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}

func (hlr *ItemTypeHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := &SettingForm.ItemTypeForm{}

	srv := SettingService.NewSettingService[ItemModel.ItemType](nil)
	srv.Create(w, r, form)

}

func (hlr *ItemTypeHandler) Update(w http.ResponseWriter, r *http.Request) {
	form := &SettingForm.ItemTypeForm{}

	srv := SettingService.NewSettingService[ItemModel.ItemType](r)
	srv.Update(w, r, form)
}

func (hlr *ItemTypeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	srv := SettingService.NewSettingService[ItemModel.ItemType](r)
	srv.Delete(w, r)
}
