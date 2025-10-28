package setting

import (
	"net/http"
	"service/internal/pkg/core"

	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"
)

type ItemUnitForm struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Type         string `json:"type"`
	IsBaseUnit   bool   `json:"isBaseUnit"`
	Conversion   string `json:"conversion"`
}

func (rule *ItemUnitForm) Validate() {
	va := xtrememdw.Validator{}
	va.Make(rule)
}

func (rule *ItemUnitForm) APIParse(r *http.Request) {
	core.BaseForm{}.APIParse(r, &rule)
}
