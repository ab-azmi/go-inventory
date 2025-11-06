package model

import (
	xtrememodel "github.com/globalxtreme/go-core/v2/model"
)

type ItemComponentCategory struct {
	xtrememodel.BaseModel
	Name      string `gorm:"column:name;type:varchar(255);not null"`
	IsForSale bool   `gorm:"column:isForSale;default:false;not null"`
}

func (ItemComponentCategory) TableName() string {
	return "item_categories"
}

func (md ItemComponentCategory) SetReference() uint {
	return md.ID
}
