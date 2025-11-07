package form

type ItemComponentBrandFilterForm struct {
	IDs    []uint `json:"ids"`
	Search string `json:"search"`
	Page   uint   `json:"page"`
	Limit  uint   `json:"limit"`
}
