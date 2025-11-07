package model

import xtrememodel "github.com/globalxtreme/go-core/v2/model"

type SettingWarehouse struct {
	xtrememodel.BaseModel
	BranchOfficeId uint    `gorm:"column:branchOfficeId;type:bigint;not null"`
	Name           string  `gorm:"column:name;type:varchar(255);not null"`
	Address        *string `gorm:"column:address;type:text"`
	ParentId       *uint   `gorm:"column:parentId"`

	Parent *SettingWarehouse `gorm:"foreignKey:parentId"`
}

func (SettingWarehouse) TableName() string {
	return "warehouses"
}

func (md SettingWarehouse) SetReference() uint {
	return md.ID
}
