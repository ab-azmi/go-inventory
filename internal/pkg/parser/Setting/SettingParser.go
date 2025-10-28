package parser

type hasSettingFields interface {
	GetArrayFields() map[string]interface{}
}

type SettingParser[T hasSettingFields] struct {
	Array  []T
	Object T
}

func (parser *SettingParser[T]) Get() []interface{} {
	var result []interface{}

	for _, obj := range parser.Array {
		firstParser := SettingParser[T]{Object: obj}
		result = append(result, firstParser.First())
	}

	return result
}

func (parser *SettingParser[T]) First() interface{} {
	object := parser.Object

	return object.GetArrayFields()
}

func (parser *SettingParser[T]) CreateActivity(action string) interface{} {
	return nil
}

func (parser *SettingParser[T]) UpdateActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *SettingParser[T]) DeleteActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser *SettingParser[T]) GeneralActivity(action string) interface{} {
	return parser.CreateActivity(action)
}
