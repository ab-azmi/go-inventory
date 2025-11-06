package parser

import "service/internal/pkg/model"

type SettingItemTypeParser struct {
	Array  []model.ItemComponentType
	Object model.ItemComponentType
}

func (parser *SettingItemTypeParser) Get() []interface{} {
	var types []interface{}

	for _, brand := range parser.Array {
		parser.Object = brand
		types = append(types, parser.First())
	}

	return types
}

func (parser *SettingItemTypeParser) First() interface{} {
	object := parser.Object

	return map[string]interface{}{
		"id":        object.ID,
		"name":      object.Name,
		"createdAt": object.CreatedAt.Format("02/01/2006 15:04"),
		"updatedAt": object.UpdatedAt.Format("02/01/2006 15:04"),
	}
}

func (parser *SettingItemTypeParser) CreateActivity(action string) interface{} {
	return parser.First()
}

func (parser *SettingItemTypeParser) UpdateActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *SettingItemTypeParser) DeleteActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *SettingItemTypeParser) GeneralActivity(action string) interface{} {
	return parser.CreateActivity(action)
}
