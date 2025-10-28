package Item

import xtrememodel "github.com/globalxtreme/go-core/v2/model"

type Item struct {
	xtrememodel.BaseModelUUID
	TypeId        uint    `gorm:"column:typeId;type:bigint;not null"`
	CategoryId    uint    `gorm:"column:categoryId;type:bigint;not null"`
	UnitId        uint    `gorm:"column:unitId;type:bigint;not null"`
	BrandId       *uint   `gorm:"column:brandId;type:bigint"`
	SKU           string  `gorm:"column:\"SKU\";type:varchar(20);not null"`
	Name          string  `gorm:"column:name;type:varchar(255);not null"`
	IsForSale     bool    `gorm:"column:isForSale;default:false;not null"`
	IsQualified   bool    `gorm:"column:isQualified;default:false;not null"`
	PurchasedCost float64 `gorm:"column:purchasedCost;type:decimal(20,2);not null;default:0.00"`
	Photo         *string `gorm:"column:photo;type:varchar(255)"`
	CreatedBy     *string `gorm:"column:createdBy;type:varchar(45);index"`
	CreatedByName *string `gorm:"column:createdByName;type:varchar(255)"`

	Type     ItemType     `gorm:"foreignKey:typeId;references:ID"`
	Category ItemCategory `gorm:"foreignKey:categoryId;references:ID"`
	Brand    *ItemBrand   `gorm:"foreignKey:brandId;references:ID"`
	Unit     ItemUnit     `gorm:"foreignKey:unitId;references:ID"`
}

func (Item) TableName() string {
	return "items"
}

func (md *Item) SetReference() uint {
	return md.ID
}
