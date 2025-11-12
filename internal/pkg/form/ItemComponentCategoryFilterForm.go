package form

type ItemComponentCategoryFilterForm struct {
	IDs       []uint `json:"ids"`
	ID        uint   `json:"id"`
	IsForSale bool   `json:"isForSale"`
	Search    string `json:"search"`
	Page      uint   `json:"page"`
	Limit     uint   `json:"limit"`
}
