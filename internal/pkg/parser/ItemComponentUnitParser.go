package parser

import "service/internal/pkg/model"

type SettingItemUnitParser struct {
	Array  []model.ItemComponentUnit
	Object model.ItemComponentUnit
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
		"createdAt":    object.CreatedAt.Format("02/01/2006 15:04"),
		"updatedAt":    object.UpdatedAt.Format("02/01/2006 15:04"),
	}
}

/** --- Activity --- */

func (parser *SettingItemUnitParser) CreateActivity(action string) interface{} {
	object := parser.Object

	return map[string]interface{}{
		"id":           object.ID,
		"name":         object.Name,
		"abbreviation": object.Abbreviation,
		"type":         object.Type,
		"isBaseUnit":   object.IsBaseUnit,
		"conversion":   object.Conversion,
	}
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
