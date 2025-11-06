package form

import (
	"net/http"
)

type SettingItemBrandFilterForm struct {
	ID             uint   `schema:"id"`
	StatusId       int    `json:"statusId"`
	Search         string `json:"search"`
	UseTransaction bool   `json:"useTransaction"`
	Page           int    `json:"page"`
	Limit          int    `json:"limit"`
	Preloads       []string
	Order          map[string]string
}

func (f *SettingItemBrandFilterForm) APIFilterParse(r *http.Request) {
	parameter := r.URL.Query()

	f.Search = parameter.Get("search")
}
