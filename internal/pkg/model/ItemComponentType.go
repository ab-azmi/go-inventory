package model

import (
	xtrememodel "github.com/globalxtreme/go-core/v2/model"
)

type ItemComponentType struct {
	xtrememodel.BaseModel
	Name string `gorm:"column:name;type:varchar(255);not null"`
}

func (ItemComponentType) TableName() string {
	return "item_types"
}

func (md ItemComponentType) SetReference() uint {
	return md.ID
}
