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

type ItemComponentTypeRepository interface {
	core.TransactionRepository

	Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemComponentType, interface{}, error)
	FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemComponentType
	Create(form form.SettingForm) model.ItemComponentType
	Update(itemType model.ItemComponentType, form form.SettingForm) model.ItemComponentType
	Delete(itemType model.ItemComponentType)
}

func NewItemComponentTypeRepository(args ...*gorm.DB) ItemComponentTypeRepository {
	repository := itemComponentTypeRepository{}

	if len(args) > 0 {
		repository.transaction = args[0]
	}

	return &repository
}

type itemComponentTypeRepository struct {
	transaction *gorm.DB
}

func (repo *itemComponentTypeRepository) SetTransaction(tx *gorm.DB) {
	repo.transaction = tx
}

func (repo *itemComponentTypeRepository) Find(parameter url.Values, args ...func(query *gorm.DB) *gorm.DB) ([]model.ItemComponentType, interface{}, error) {
	var types []model.ItemComponentType

	fromDate, toDate := core.SetDateRange(parameter)
	query := config.PgSQL.Where(`"createdAt" BETWEEN ? AND ?`, fromDate, toDate)

	if search := parameter.Get("search"); len(search) > 3 {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query = query.Order("id DESC")

	types, pagination, err := xtrememodel.Paginate(query, parameter, model.ItemComponentType{})
	if err != nil {
		return nil, nil, err
	}

	return types, pagination, nil
}

func (repo *itemComponentTypeRepository) FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemComponentType {
	var itemType model.ItemComponentType

	query := config.PgSQL
	if len(args) > 0 {
		query = args[0](query)
	}

	err := query.First(&itemType, "id = ?", id).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemTypeGet(err.Error())
	}

	return itemType
}

func (repo *itemComponentTypeRepository) Create(form form.SettingForm) model.ItemComponentType {
	itemType := model.ItemComponentType{
		Name: form.Name,
	}

	err := repo.transaction.Create(&itemType).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemTypeCreate(err.Error())
	}

	return itemType
}

func (repo *itemComponentTypeRepository) Update(itemType model.ItemComponentType, form form.SettingForm) model.ItemComponentType {
	itemType.Name = form.Name

	err := repo.transaction.Save(&itemType).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemTypeUpdate(err.Error())
	}

	return itemType
}

func (repo *itemComponentTypeRepository) Delete(itemType model.ItemComponentType) {
	err := repo.transaction.Delete(&itemType).Error
	if err != nil {
		gxErr.ErrXtremeSettingItemTypeDelete(err.Error())
	}
}
