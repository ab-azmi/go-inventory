package form

type IdNameFilterForm struct {
	IDs    []uint `json:"ids"`
	ID     uint   `json:"id"`
	Search string `json:"search"`
	Page   uint   `json:"page"`
	Limit  uint   `json:"limit"`
}
