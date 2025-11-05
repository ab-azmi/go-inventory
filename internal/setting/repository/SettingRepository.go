package repository

import (
	"net/url"
	"service/internal/pkg/config"
	"service/internal/pkg/core"

	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	"gorm.io/gorm"
)

type SettingModel interface {
	FeatureName() string
}

type SettingRepository[T SettingModel] interface {
	core.TransactionRepository

	Find(parameter url.Values) ([]T, interface{}, error)
	FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) T
	Delete(model T)
	Create(model T) T
	Update(model T) T
}

func NewSettingRepository[T SettingModel](args ...*gorm.DB) SettingRepository[T] {
	repository := settingRepository[T]{}
	if len(args) > 0 {
		repository.transaction = args[0]
	}

	return &repository
}

type settingRepository[T SettingModel] struct {
	transaction *gorm.DB
}

func (repo *settingRepository[T]) SetTransaction(tx *gorm.DB) {
	repo.transaction = tx
}

func (repo *settingRepository[T]) Find(parameter url.Values) ([]T, interface{}, error) {
	fromDate, toDate := core.SetDateRange(parameter)

	query := config.PgSQL.Where("\"createdAt\" BETWEEN ? AND ?", fromDate, toDate)

	if search := parameter.Get("search"); len(search) > 3 {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	var model T
	objects, pagination, err := xtrememodel.Paginate(query.Order("id DESC"), parameter, model)
	if err != nil {
		return nil, nil, err
	}

	return objects, pagination, nil
}

func (repo *settingRepository[T]) FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) T {
	var model T

	query := config.PgSQL
	if len(args) > 0 {
		query = args[0](query)
	}

	err := query.First(&model, "id = ?", id).Error
	if err != nil {
		// xtremeErr.ErrXtremeSettingGet(model.FeatureName(), err.Error())
	}

	return model
}

func (repo *settingRepository[T]) Create(model T) T {
	err := repo.transaction.Create(&model).Error
	if err != nil {
		// xtremeErr.ErrXtremeSettingCreate(model.FeatureName(), err.Error())
	}

	return model
}

func (repo *settingRepository[T]) Update(model T) T {
	err := repo.transaction.Model(&model).Updates(&model).Error
	if err != nil {
		// xtremeErr.ErrXtremeSettingUpdate(model.FeatureName(), err.Error())
	}

	return model
}

func (repo *settingRepository[T]) Delete(model T) {
	err := repo.transaction.Delete(&model).Error

	if err != nil {
		// xtremeErr.ErrXtremeSettingDelete(model.FeatureName(), err.Error())
	}
}
