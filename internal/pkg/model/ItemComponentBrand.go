package model

import (
	xtrememodel "github.com/globalxtreme/go-core/v2/model"
)

type ItemComponentBrand struct {
	xtrememodel.BaseModel
	Name string `gorm:"column:name;type:varchar(255);not null"`
}

func (ItemComponentBrand) TableName() string {
	return "item_component_brands"
}

func (md ItemComponentBrand) SetReference() uint {
	return md.ID
}
