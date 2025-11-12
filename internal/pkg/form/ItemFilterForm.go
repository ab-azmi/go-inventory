package form

type ItemFilterForm struct {
	IDs         []uint   `json:"ids"`
	ID          uint     `json:"id"`
	Search      string   `json:"search"`
	TypeIds     []uint   `json:"typeIds"`
	CategoryIds []uint   `json:"categoryIds"`
	UnitIds     []uint   `json:"unitIds"`
	BrandIds    []uint   `json:"brandIds"`
	Preloads    []string `json:"preloads"`
}
