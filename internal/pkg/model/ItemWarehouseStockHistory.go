package model

import (
	xtrememodel "github.com/globalxtreme/go-core/v2/model"
)

type ItemWarehouseStockHistory struct {
	xtrememodel.BaseModel
	ItemWarehouseId uint                           `gorm:"column:itemWarehouseId;type:bigint;not null"`
	Type            string                         `gorm:"column:type;type:varchar(100);not null"`
	Adjusted        float64                        `gorm:"column:adjusted;type:decimal(15,2);default:0.00"`
	NewQuantity     float64                        `gorm:"column:newQuantity;type:decimal(15,2);default:0.00"`
	Description     *string                        `gorm:"column:description;type:text"`
	SerialNumbers   *xtrememodel.ArrayStringColumn `gorm:"column:serialNumbers;type:json"`

	ItemWarehouse ItemWarehouse `gorm:"foreignKey:itemWarehouseId"`
}

func (ItemWarehouseStockHistory) TableName() string {
	return "item_warehouse_stock_histories"
}

func (md ItemWarehouseStockHistory) SetReference() uint {
	return md.ID
}
