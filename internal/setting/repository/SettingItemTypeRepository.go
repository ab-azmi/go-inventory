package repository

import (
	"net/url"
	"service/internal/pkg/config"
	"service/internal/pkg/core"
	gxErr "service/internal/pkg/error"
	"service/internal/pkg/form"
	"service/internal/pkg/model"

	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	"gorm.io/gorm"
)

type SettingItemTypeRepository interface {
	core.TransactionRepository

	Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemType, interface{}, error)
	FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemType
	Create(form form.SettingForm) model.ItemType
	Update(itemType model.ItemType, form form.SettingForm) model.ItemType
	Delete(itemType model.ItemType)
}

func NewSettingItemTypeRepository(args ...*gorm.DB) SettingItemTypeRepository {
	repository := settingItemTypeRepository{}

	if len(args) > 0 {
		repository.transaction = args[0]
	}

	return &repository
}

type settingItemTypeRepository struct {
	transaction *gorm.DB
}

func (repo *settingItemTypeRepository) SetTransaction(tx *gorm.DB) {
	repo.transaction = tx
}

func (repo *settingItemTypeRepository) Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemType, interface{}, error) {
	var types []model.ItemType

	fromDate, toDate := core.SetDateRange(parameter)
	query := config.PgSQL.Where(`"createdAt" BETWEEN ? AND ?`, fromDate, toDate)

	if search := parameter.Get("search"); len(search) > 3 {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query = query.Order("id DESC")

	types, pagination, err := xtrememodel.Paginate(query, parameter, model.ItemType{})
	if err != nil {
		return nil, nil, err
	}

	return types, pagination, nil
}

func (repo *settingItemTypeRepository) FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemType {
	var itemType model.ItemType

	query := config.PgSQL
	if len(args) > 0 {
		query = args[0](query)
	}

	err := query.First(&itemType, "id = ?", id).Error
	if err != nil {
		gxErr.ErrXtremeSettingGet(itemType.FeatureName(), err.Error())
	}

	return itemType
}

func (repo *settingItemTypeRepository) Create(form form.SettingForm) model.ItemType {
	itemType := model.ItemType{
		Name: form.Name,
	}

	err := repo.transaction.Create(&itemType).Error
	if err != nil {
		gxErr.ErrXtremeSettingCreate(itemType.FeatureName(), err.Error())
	}

	return itemType
}

func (repo *settingItemTypeRepository) Update(itemType model.ItemType, form form.SettingForm) model.ItemType {
	itemType.Name = form.Name

	err := repo.transaction.Save(&itemType).Error
	if err != nil {
		gxErr.ErrXtremeSettingUpdate(itemType.FeatureName(), err.Error())
	}

	return itemType
}

func (repo *settingItemTypeRepository) Delete(itemType model.ItemType) {
	err := repo.transaction.Delete(&itemType).Error
	if err != nil {
		gxErr.ErrXtremeSettingDelete(itemType.FeatureName(), err.Error())
	}
}
