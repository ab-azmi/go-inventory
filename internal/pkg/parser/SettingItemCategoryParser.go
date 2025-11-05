package parser

import "service/internal/pkg/model"

type SettingItemCategoryParser struct {
	Array  []model.ItemCategory
	Object model.ItemCategory
}

func (parser *SettingItemCategoryParser) Get() []interface{} {
	var categories []interface{}

	for _, brand := range parser.Array {
		parser.Object = brand
		categories = append(categories, parser.First())
	}

	return categories
}

func (parser *SettingItemCategoryParser) First() interface{} {
	object := parser.Object

	return map[string]interface{}{
		"id":        object.ID,
		"name":      object.Name,
		"isForSale": object.IsForSale,
		"createdAt": object.CreatedAt.Format("02/01/2006 15:04"),
		"updatedAt": object.UpdatedAt.Format("02/01/2006 15:04"),
	}
}

func (parser *SettingItemCategoryParser) CreateActivity(action string) interface{} {
	return parser.First()
}

func (parser *SettingItemCategoryParser) UpdateActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *SettingItemCategoryParser) DeleteActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *SettingItemCategoryParser) GeneralActivity(action string) interface{} {
	return parser.CreateActivity(action)
}
