package Setting

import xtrememodel "github.com/globalxtreme/go-core/v2/model"

type Warehouse struct {
	xtrememodel.BaseModel
	BranchOfficeId uint    `gorm:"column:branchOfficeId;type:bigint;not null"`
	Name           string  `gorm:"column:name;type:varchar(255);not null"`
	Address        *string `gorm:"column:address;type:text"`
	ParentId       *uint   `gorm:"column:parentId"`

	Parent *Warehouse `gorm:"foreignKey:parentId;references:ID"`
}

func (Warehouse) TableName() string {
	return "warehouses"
}

func (md Warehouse) SetReference() uint {
	return md.ID
}
