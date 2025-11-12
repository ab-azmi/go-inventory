package parser

import "service/internal/pkg/model"

type ItemComponentBrandParser struct {
	Array  []model.ItemComponentBrand
	Object model.ItemComponentBrand
}

func (parser *ItemComponentBrandParser) Get() []interface{} {
	var brands []interface{}

	for _, brand := range parser.Array {
		parser.Object = brand
		brands = append(brands, parser.First())
	}

	return brands
}

func (parser *ItemComponentBrandParser) First() interface{} {
	object := parser.Object

	return map[string]interface{}{
		"id":        object.ID,
		"name":      object.Name,
		"createdAt": object.CreatedAt.Format("02/01/2006 15:04"),
		"updatedAt": object.UpdatedAt.Format("02/01/2006 15:04"),
	}
}

/** --- Activity --- */

func (parser *ItemComponentBrandParser) CreateActivity(action string) interface{} {
	object := parser.Object

	return map[string]interface{}{
		"id":   object.ID,
		"name": object.Name,
	}
}

func (parser *ItemComponentBrandParser) UpdateActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *ItemComponentBrandParser) DeleteActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *ItemComponentBrandParser) GeneralActivity(action string) interface{} {
	return parser.CreateActivity(action)
}
