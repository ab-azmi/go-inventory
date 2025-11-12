package parser

import "service/internal/pkg/model"

type ItemComponentTypeParser struct {
	Array  []model.ItemComponentType
	Object model.ItemComponentType
}

func (parser *ItemComponentTypeParser) Get() []interface{} {
	var types []interface{}

	for _, brand := range parser.Array {
		parser.Object = brand
		types = append(types, parser.First())
	}

	return types
}

func (parser *ItemComponentTypeParser) First() interface{} {
	object := parser.Object

	return map[string]interface{}{
		"id":        object.ID,
		"name":      object.Name,
		"createdAt": object.CreatedAt.Format("02/01/2006 15:04"),
		"updatedAt": object.UpdatedAt.Format("02/01/2006 15:04"),
	}
}

/** --- Activity --- */

func (parser *ItemComponentTypeParser) CreateActivity(action string) interface{} {
	object := parser.Object

	return map[string]interface{}{
		"id":   object.ID,
		"name": object.Name,
	}
}

func (parser *ItemComponentTypeParser) UpdateActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *ItemComponentTypeParser) DeleteActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *ItemComponentTypeParser) GeneralActivity(action string) interface{} {
	return parser.CreateActivity(action)
}
