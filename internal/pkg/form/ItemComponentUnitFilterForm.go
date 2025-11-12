package form

type ItemComponentUnitFilterForm struct {
	IDs        []uint   `json:"ids"`
	ID         uint     `json:"id"`
	Search     string   `json:"search"`
	Types      []string `json:"types"`
	IsBaseUnit bool     `json:"isBaseUnit"`
}
