package setting

import (
	"fmt"
	"net/http"
	"service/internal/pkg/activity"
	"service/internal/pkg/config"
	"service/internal/pkg/constant"
	SettingParser "service/internal/pkg/parser/Setting"
	"service/internal/pkg/port"
	"service/internal/setting/repository"
	"strconv"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type HasSettingForm interface {
	Validate()
	APIParse(r *http.Request)
}

type HasSettingService[T repository.SettingModel, F HasSettingForm] interface {
	ParseForm(form F) T
	TableName() string
	SetReference() uint
	GetArrayFields() map[string]interface{}
}

type HasSettingModel[T repository.SettingModel, F HasSettingForm] interface {
	repository.SettingModel
	HasSettingService[T, F]
}

type SettingService[T HasSettingModel[T, F], F HasSettingForm] interface {
	SetTransaction(tx *gorm.DB)
	SetActivityRepository(repo port.ActivityRepository)

	Create(w http.ResponseWriter, r *http.Request, form F)
	Update(w http.ResponseWriter, r *http.Request, form F)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewSettingService[T HasSettingModel[T, F], F HasSettingForm](r *http.Request) SettingService[T, F] {
	var model T
	if r != nil {
		if id := mux.Vars(r)["id"]; id != "" {
			repo := repository.NewSettingRepository[T]()

			parsedId, _ := strconv.ParseUint(id, 10, 0)
			model = repo.FirstById(uint(parsedId))
		}
	}

	return &settingService[T, F]{
		model: model,
	}
}

type settingService[T HasSettingModel[T, F], F HasSettingForm] struct {
	tx *gorm.DB

	repository   repository.SettingRepository[T]
	activityRepo port.ActivityRepository

	model T
}

func (srv *settingService[T, F]) SetTransaction(tx *gorm.DB) {
	srv.tx = tx
}

func (srv *settingService[T, F]) SetActivityRepository(repo port.ActivityRepository) {
	srv.activityRepo = repo
}

func (srv *settingService[T, F]) Create(w http.ResponseWriter, r *http.Request, form F) {
	var model T

	form.APIParse(r)
	form.Validate()

	var parser SettingParser.SettingParser[T]

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewSettingRepository[T](tx)

		model = srv.repository.Create(model.ParseForm(form))

		parser.Object = model

		activity.UseActivity{}.SetReference(model).SetParser(&parser).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Created new %s", model.FeatureName()))

		return nil
	})

	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (srv *settingService[T, F]) Update(w http.ResponseWriter, r *http.Request, form F) {

	form.APIParse(r)
	form.Validate()

	parser := SettingParser.SettingParser[T]{Object: srv.model}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		updateActivity := activity.UseActivity{}.SetReference(srv.model).SetParser(&parser).SetOldProperty(constant.ACTION_UPDATE)
		srv.repository = repository.NewSettingRepository[T](tx)

		srv.model = srv.repository.Update(srv.model.ParseForm(form))

		parser.Object = srv.model

		updateActivity.SetReference(srv.model).SetParser(&parser).SetNewProperty(constant.ACTION_UPDATE).
			Save(fmt.Sprintf("Updated %s", srv.model.FeatureName()))

		return nil
	})

	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (srv *settingService[T, F]) Delete(w http.ResponseWriter, r *http.Request) {
	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewSettingRepository[T](tx)

		srv.repository.Delete(srv.model)

		activity.UseActivity{}.SetReference(srv.model).Save(fmt.Sprintf("Delete %s", srv.model.FeatureName()))

		return nil
	})

	res := xtremeres.Response{}
	res.Success(w)
}
