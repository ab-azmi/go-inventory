package form

type ItemComponentUnitFilterForm struct {
	IDs        []uint   `json:"ids"`
	Search     string   `json:"search"`
	Types      []string `json:"types"`
	IsBaseUnit bool     `json:"isBaseUnit"`
}
