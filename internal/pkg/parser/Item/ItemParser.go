package parser

import (
	"service/internal/pkg/model"
)

type ItemParser struct {
	Array  []model.Item
	Object model.Item
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

	var brand map[string]interface{}
	if object.Brand != nil {
		brand = map[string]interface{}{
			"id":   object.Brand.ID,
			"name": object.Brand.Name,
		}
	}

	return map[string]interface{}{
		"id":            object.ID,
		"name":          object.Name,
		"SKU":           object.SKU,
		"purchasedCost": object.PurchasedCost,
		"isForSale":     object.IsForSale,
		"isQualified":   object.IsQualified,
		"brand":         brand,
		"type": map[string]interface{}{
			"id":   object.Type.ID,
			"name": object.Type.Name,
		},
		"category": map[string]interface{}{
			"id":   object.Category.ID,
			"name": object.Category.Name,
		},
		"unit": map[string]interface{}{
			"id":           object.Unit.ID,
			"name":         object.Unit.Name,
			"abbreviation": object.Unit.Abbreviation,
		},
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
