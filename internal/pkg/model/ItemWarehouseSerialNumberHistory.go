package model

import xtrememodel "github.com/globalxtreme/go-core/v2/model"

type ItemWarehouseSerialNumberHistory struct {
	xtrememodel.BaseModel
	SerialNumberId uint    `gorm:"column:serialNumberId;type:bigint;not null"`
	Type           string  `gorm:"column:type;type:varchar(10);check:type IN ('in','out');default:'in'"`
	Action         string  `gorm:"column:action;type:varchar(100);not null"`
	Description    *string `gorm:"column:description;type:text"`
	Reference      *uint   `gorm:"column:reference;type:bigint"`
	ReferenceType  *string `gorm:"column:referenceType;type:varchar(100)"`

	SerialNumber ItemWarehouseSerialNumber `gorm:"foreignKey:serialNumberId"`
}

func (ItemWarehouseSerialNumberHistory) TableName() string {
	return "item_warehouse_serial_number_histories"
}

func (md *ItemWarehouseSerialNumberHistory) SetReference() uint {
	return md.ID
}
