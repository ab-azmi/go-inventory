package parser

import (
	ItemModel "service/internal/pkg/model/Item"
)

type ItemParser struct {
	Array  []ItemModel.Item
	Object ItemModel.Item
}

func (parser *ItemParser) Get() []interface{} {
	var result []interface{}
	for _, arr := range parser.Array {
		firstParser := ItemParser{Object: arr}
		result = append(result, firstParser.First())
	}

	return result
}

func (parser *ItemParser) First() interface{} {
	object := parser.Object

	return map[string]interface{}{
		"id":        object.ID,
		"createdAt": object.CreatedAt.Format("02/01/2006 15:04"),
	}
}

func (parser *ItemParser) CreateActivity(action string) interface{} {
	object := parser.Object

	return map[string]interface{}{
		"id":        object.ID,
		"createdAt": object.CreatedAt.Format("02/01/2006 15:04"),
	}
}

func (parser *ItemParser) UpdateActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *ItemParser) DeleteActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *ItemParser) GeneralActivity(action string) interface{} {
	return parser.CreateActivity(action)
}
