package model

import xtrememodel "github.com/globalxtreme/go-core/v2/model"

type ItemWarehouseStock struct {
	xtrememodel.BaseModel
	ItemWarehouseId   uint    `gorm:"column:itemWarehouseId;type:bigint;not null"`
	Physical          float64 `gorm:"column:physical;type:decimal(15,2);default:0.00"`
	PhysicalAllocated float64 `gorm:"column:physicalAllocated;decimal(15,2);default:0.00"`

	ItemWarehouse ItemWarehouse `gorm:"foreignKey:itemWarehouseId;references:ID"`
}

func (ItemWarehouseStock) TableName() string {
	return "item_warehouse_stocks"
}

func (md ItemWarehouseStock) SetReference() uint {
	return md.ID
}
