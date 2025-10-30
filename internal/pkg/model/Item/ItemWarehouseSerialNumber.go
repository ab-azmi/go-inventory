package Item

import xtrememodel "github.com/globalxtreme/go-core/v2/model"

type ItemWarehouseSerialNumber struct {
	xtrememodel.BaseModel
	ItemWarehouseId uint   `gorm:"column:itemWarehouseId;type:bigint;not null"`
	Number          string `gorm:"column:number;type:varchar(100);not null"`
	Status          string `gorm:"column:status;type:varchar(10);check:status IN ('in','out');default:'in'"`

	ItemWarehouse ItemWarehouse `gorm:"foreignKey:itemWarehouseId;references:ID"`
}

func (ItemWarehouseSerialNumber) TableName() string {
	return "item_warehouse_serial_numbers"
}

func (md ItemWarehouseSerialNumber) SetReference() uint {
	return md.ID
}
