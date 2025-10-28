package Item

import xtrememodel "github.com/globalxtreme/go-core/v2/model"

type ItemType struct {
	xtrememodel.BaseModel
	Name string `gorm:"column:name;type:varchar(255);not null"`
}

func (ItemType) TableName() string {
	return "item_types"
}

func (md *ItemType) SetReference() uint {
	return md.ID
}

func (ItemType) ErrorName() string {
	return "Item Type"
}

func (it ItemType) GetArrayFields() map[string]interface{} {
	return map[string]interface{}{
		"id":   it.ID,
		"name": it.Name,
	}
}
