package Item

import (
	SettingForm "service/internal/pkg/form/setting"

	xtrememodel "github.com/globalxtreme/go-core/v2/model"
)

type ItemUnit struct {
	xtrememodel.BaseModel
	Name          string  `gorm:"column:name;type:varchar(255);not null"`
	Abbreviation  string  `gorm:"column:abbreviation;type:varchar(100);not null"`
	Type          string  `gorm:"column:type;type:varchar(100);not null"`
	IsBaseUnit    bool    `gorm:"column:isBaseUnit;default:false;not null"`
	Conversion    string  `gorm:"column:conversion;type:varchar(100);default:0;not null"`
	CreatedBy     *string `gorm:"column:createdBy;type:varchar(45);index"`
	CreatedByName *string `gorm:"column:createdByName;type:varchar(255)"`
}

func (ItemUnit) TableName() string {
	return "item_units"
}

func (md ItemUnit) SetReference() uint {
	return md.ID
}

func (ItemUnit) FeatureName() string {
	return "Item Unit"
}

func (unit ItemUnit) GetArrayFields() map[string]interface{} {
	return map[string]interface{}{
		"id":           unit.ID,
		"name":         unit.Name,
		"abbreviation": unit.Abbreviation,
	}
}

func (ItemUnit) ParseForm(form SettingForm.ItemUnitForm) ItemUnit {
	return ItemUnit{
		Name:         form.Name,
		Abbreviation: form.Abbreviation,
		Type:         form.Type,
		IsBaseUnit:   form.IsBaseUnit,
		Conversion:   form.Conversion,
	}
}
