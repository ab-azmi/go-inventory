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

type ItemUnitHandler struct{}

func (hlr *ItemUnitHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := SettingRepo.NewSettingRepository[ItemModel.ItemUnit]()
	types, pagination, _ := repo.Find(r.URL.Query())

	parser := SettingParser.SettingParser[ItemModel.ItemUnit]{Array: types}

	response := xtremeres.Response{Array: parser.Get(), Pagination: &pagination}
	response.Success(w)
}

func (hlr *ItemUnitHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := SettingForm.ItemUnitForm{}
	form.APIParse(r)
	form.Validate()

	srv := SettingService.NewSettingService[ItemModel.ItemUnit]()
	srv.SetActivityRepository(ActivityRepository.NewActivityRepository())

	ItemUnit := srv.Create(form)

	parser := SettingParser.SettingParser[ItemModel.ItemUnit]{Object: ItemUnit}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}
