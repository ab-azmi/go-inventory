package Item

import xtrememodel "github.com/globalxtreme/go-core/v2/model"

type ItemCategory struct {
	xtrememodel.BaseModel
	Name      string `gorm:"column:name;type:varchar(255);not null"`
	IsForSale bool   `gorm:"column:isForSale;default:false;not null"`
}

func (ItemCategory) TableName() string {
	return "item_categories"
}

func (md *ItemCategory) SetReference() uint {
	return md.ID
}
