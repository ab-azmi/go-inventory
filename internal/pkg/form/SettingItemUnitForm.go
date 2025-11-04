package form

import (
	"net/http"
	"service/internal/pkg/core"

	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"
)

type SettingItemUnitForm struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Type         string `json:"type"`
	IsBaseUnit   bool   `json:"isBaseUnit"`
	Conversion   string `json:"conversion"`
}

func (rule *SettingItemUnitForm) Validate() {
	va := xtrememdw.Validator{}
	va.Make(rule)
}

func (rule *SettingItemUnitForm) APIParse(r *http.Request) {
	core.BaseForm{}.APIParse(r, &rule)
}
