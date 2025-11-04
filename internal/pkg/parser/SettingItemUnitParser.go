package parser

import "service/internal/pkg/model"

type SettingItemUnitParser struct {
	Array  []model.ItemUnit
	Object model.ItemUnit
}

func (parser *SettingItemUnitParser) Get() []interface{} {
	var categories []interface{}

	for _, brand := range parser.Array {
		parser.Object = brand
		categories = append(categories, parser.First())
	}

	return categories
}

func (parser *SettingItemUnitParser) First() interface{} {
	object := parser.Object

	return map[string]interface{}{
		"id":           object.ID,
		"name":         object.Name,
		"abbreviation": object.Abbreviation,
		"type":         object.Type,
		"isBaseUnit":   object.IsBaseUnit,
		"conversion":   object.Conversion,
		"createdAt":    object.CreatedAt,
		"updatedAt":    object.UpdatedAt,
	}
}

func (parser *SettingItemUnitParser) CreateActivity(action string) interface{} {
	return parser.First()
}

func (parser *SettingItemUnitParser) UpdateActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *SettingItemUnitParser) DeleteActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *SettingItemUnitParser) GeneralActivity(action string) interface{} {
	return parser.CreateActivity(action)
}
