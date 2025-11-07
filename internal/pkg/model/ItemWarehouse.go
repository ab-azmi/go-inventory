package model

import (
	"time"

	xtrememodel "github.com/globalxtreme/go-core/v2/model"
)

type ItemWarehouse struct {
	xtrememodel.BaseModelUUID
	ItemId       uint       `gorm:"column:itemId;type:bigint;not null"`
	WarehouseId  uint       `gorm:"column:warehouseId;type:bigint;not null"`
	OrderType    string     `gorm:"column:orderType;type:varchar(100);default:'available'"`
	SellingPrice float64    `gorm:"column:sellingPrice;type:decimal(20, 2);default:0.00"`
	IsIncludeTax bool       `gorm:"column:isIncludeTax;default:false"`
	Location     *string    `gorm:"column:location;type:text"`
	ActivatedAt  *time.Time `gorm:"column:activatedAt;"`

	Item      Item             `gorm:"foreignKey:itemId"`
	Warehouse SettingWarehouse `gorm:"foreignKey:warehouseId"`
}

func (ItemWarehouse) TableName() string {
	return "item_warehouses"
}

func (md ItemWarehouse) SetReference() uint {
	return md.ID
}
