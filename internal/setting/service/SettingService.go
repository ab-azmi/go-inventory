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

	xtremeres "github.com/globalxtreme/go-core/v2/response"
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
}

func NewSettingService[T HasSettingModel[T, F], F HasSettingForm]() SettingService[T, F] {
	return &settingService[T, F]{}
}

type settingService[T HasSettingModel[T, F], F HasSettingForm] struct {
	tx *gorm.DB

	repository   repository.SettingRepository[T]
	activityRepo port.ActivityRepository
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

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewSettingRepository[T](tx)

		model = srv.repository.Create(model.ParseForm(form))

		activity.UseActivity{}.SetReference(model).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Created new %s", model.FeatureName()))

		return nil
	})

	parser := SettingParser.SettingParser[T]{Object: model}
	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}
