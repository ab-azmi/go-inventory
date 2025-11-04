package model

import (
	"service/internal/pkg/form"

	xtrememodel "github.com/globalxtreme/go-core/v2/model"
)

type ItemCategory struct {
	xtrememodel.BaseModel
	Name      string `gorm:"column:name;type:varchar(255);not null"`
	IsForSale bool   `gorm:"column:isForSale;default:false;not null"`
}

func (ItemCategory) TableName() string {
	return "item_categories"
}

func (md ItemCategory) SetReference() uint {
	return md.ID
}

func (ItemCategory) FeatureName() string {
	return "Item Category"
}

func (ic ItemCategory) GetArrayFields() map[string]interface{} {
	return map[string]interface{}{
		"id":        ic.ID,
		"name":      ic.Name,
		"isForSale": ic.IsForSale,
	}
}

func (ic ItemCategory) ParseForm(form *form.SettingItemCategoryForm) ItemCategory {
	return ItemCategory{
		BaseModel: xtrememodel.BaseModel{
			ID: ic.ID,
		},
		Name:      form.Name,
		IsForSale: form.IsForSale,
	}
}
