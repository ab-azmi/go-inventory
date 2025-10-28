package setting

import (
	"fmt"
	"service/internal/pkg/activity"
	"service/internal/pkg/config"
	"service/internal/pkg/constant"
	"service/internal/pkg/port"
	"service/internal/setting/repository"

	"gorm.io/gorm"
)

type HasSettingService[T repository.SettingModel, F any] interface {
	ParseForm(form F) T
	TableName() string
	SetReference() uint
}

type ParsableSettingModel[T repository.SettingModel, F any] interface {
	repository.SettingModel
	HasSettingService[T, F]
}

type SettingService[T ParsableSettingModel[T, F], F any] interface {
	SetTransaction(tx *gorm.DB)
	SetActivityRepository(repo port.ActivityRepository)

	Create(form F) T
}

func NewSettingService[T ParsableSettingModel[T, F], F any]() SettingService[T, F] {
	return &settingService[T, F]{}
}

type settingService[T ParsableSettingModel[T, F], F any] struct {
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

func (srv *settingService[T, F]) Create(form F) T {
	var model T

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewSettingRepository[T](tx)

		model = srv.repository.Create(model.ParseForm(form))

		activity.UseActivity{}.SetReference(model).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Created new %s", model.FeatureName()))

		return nil
	})

	return model
}
