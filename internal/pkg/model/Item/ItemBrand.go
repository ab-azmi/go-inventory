package Item

import (
	SettingForm "service/internal/pkg/form/setting"

	xtrememodel "github.com/globalxtreme/go-core/v2/model"
)

type ItemBrand struct {
	xtrememodel.BaseModel
	Name string `gorm:"column:name;type:varchar(255);not null"`
}

func (ItemBrand) TableName() string {
	return "item_brands"
}

func (md ItemBrand) SetReference() uint {
	return md.ID
}

func (ItemBrand) FeatureName() string {
	return "Item Brand"
}

func (ib ItemBrand) GetArrayFields() map[string]interface{} {
	return map[string]interface{}{
		"id":   ib.ID,
		"name": ib.Name,
	}
}

func (ItemBrand) ParseForm(form *SettingForm.ItemBrandForm) ItemBrand {
	return ItemBrand{
		Name: form.Name,
	}
}
