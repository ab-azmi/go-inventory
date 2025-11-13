package form

type ComponentFilterForm struct {
	IDs     []uint `json:"ids"`
	ID      uint   `json:"id"`
	Search  string `json:"search"`
	Page    uint   `json:"page"`
	OrderBy string `json:"orderBy"` //name:asc
	Limit   uint   `json:"limit"`
}
