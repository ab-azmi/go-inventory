package setting

import (
	"net/http"
	"service/internal/pkg/core"

	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"
)

type ItemCategoryForm struct {
	Name      string `json:"name"`
	IsForSale bool   `json:"isForSale"`
}

func (rule *ItemCategoryForm) Validate() {
	va := xtrememdw.Validator{}
	va.Make(rule)
}

func (rule *ItemCategoryForm) APIParse(r *http.Request) {
	core.BaseForm{}.APIParse(r, &rule)
}
