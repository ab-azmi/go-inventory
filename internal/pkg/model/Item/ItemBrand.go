package Item

import xtrememodel "github.com/globalxtreme/go-core/v2/model"

type ItemBrand struct {
	xtrememodel.BaseModel
	Name string `gorm:"column:name;type:varchar(255);not null"`
}

func (ItemBrand) TableName() string {
	return "item_brands"
}

func (md *ItemBrand) SetReference() uint {
	return md.ID
}
