package Item

import xtrememodel "github.com/globalxtreme/go-core/v2/model"

type ItemUnit struct {
	xtrememodel.BaseModel
	Name          string  `gorm:"column:name;type:varchar(255);not null"`
	Abbreviation  string  `gorm:"column:abbreviation;type:varchar(100);not null"`
	Type          string  `gorm:"column:type;type:varchar(100);not null"`
	IsBaseUnit    bool    `gorm:"column:isBaseUnit;default:false;not null"`
	Conversion    string  `gorm:"column:conversion;type:varchar(100);default:0;not null"`
	CreatedBy     *string `gorm:"column:createdBy;type:varchar(45);index"`
	CreatedByName *string `gorm:"column:createdByName;type:varchar(255)"`
}

func (ItemUnit) TableName() string {
	return "item_units"
}

func (md *ItemUnit) SetReference() uint {
	return md.ID
}
