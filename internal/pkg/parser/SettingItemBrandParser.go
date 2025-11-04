package parser

import "service/internal/pkg/model"

type SettingItemBrandParser struct {
	Array  []model.ItemBrand
	Object model.ItemBrand
}

func (parser *SettingItemBrandParser) Get() []interface{} {
	var brands []interface{}

	for _, brand := range parser.Array {
		parser.Object = brand
		brands = append(brands, parser.First())
	}

	return brands
}

func (parser *SettingItemBrandParser) First() interface{} {
	object := parser.Object

	return map[string]interface{}{
		"id":        object.ID,
		"name":      object.Name,
		"createdAt": object.CreatedAt,
		"updatedAt": object.UpdatedAt,
	}
}

func (parser *SettingItemBrandParser) CreateActivity(action string) interface{} {
	return parser.First()
}

func (parser *SettingItemBrandParser) UpdateActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *SettingItemBrandParser) DeleteActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *SettingItemBrandParser) GeneralActivity(action string) interface{} {
	return parser.CreateActivity(action)
}
