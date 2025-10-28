package setting

import (
	"net/http"
	ActivityRepository "service/internal/activity/repository"
	SettingForm "service/internal/pkg/form/setting"
	ItemModel "service/internal/pkg/model/Item"
	SettingParser "service/internal/pkg/parser/Setting"
	SettingRepo "service/internal/setting/repository"
	SettingService "service/internal/setting/service"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
)

type ItemCategoryHandler struct{}

func (hlr *ItemCategoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := SettingRepo.NewSettingRepository[ItemModel.ItemCategory]()
	types, pagination, _ := repo.Find(r.URL.Query())

	parser := SettingParser.SettingParser[ItemModel.ItemCategory]{Array: types}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}

func (hlr *ItemCategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := SettingForm.ItemCategoryForm{}
	form.APIParse(r)
	form.Validate()

	srv := SettingService.NewSettingService[ItemModel.ItemCategory]()
	srv.SetActivityRepository(ActivityRepository.NewActivityRepository())

	ItemCategory := srv.Create(form)

	parser := SettingParser.SettingParser[ItemModel.ItemCategory]{Object: ItemCategory}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}
