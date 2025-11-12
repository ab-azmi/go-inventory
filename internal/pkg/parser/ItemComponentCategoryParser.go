package parser

import "service/internal/pkg/model"

type ItemComponentCategoryParser struct {
	Array  []model.ItemComponentCategory
	Object model.ItemComponentCategory
}

func (parser *ItemComponentCategoryParser) Get() []interface{} {
	var categories []interface{}

	for _, brand := range parser.Array {
		parser.Object = brand
		categories = append(categories, parser.First())
	}

	return categories
}

func (parser *ItemComponentCategoryParser) First() interface{} {
	object := parser.Object

	return map[string]interface{}{
		"id":        object.ID,
		"name":      object.Name,
		"isForSale": object.IsForSale,
		"createdAt": object.CreatedAt.Format("02/01/2006 15:04"),
		"updatedAt": object.UpdatedAt.Format("02/01/2006 15:04"),
	}
}

/** --- Activity --- */

func (parser *ItemComponentCategoryParser) CreateActivity(action string) interface{} {
	object := parser.Object

	return map[string]interface{}{
		"id":        object.ID,
		"name":      object.Name,
		"isForSale": object.IsForSale,
	}
}

func (parser *ItemComponentCategoryParser) UpdateActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *ItemComponentCategoryParser) DeleteActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *ItemComponentCategoryParser) GeneralActivity(action string) interface{} {
	return parser.CreateActivity(action)
}
