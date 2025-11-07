package form

import (
	"net/http"
	"service/internal/pkg/core"

	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"
)

type ItemComponentCategoryForm struct {
	Name      string `json:"name"`
	IsForSale bool   `json:"isForSale"`
}

func (rule *ItemComponentCategoryForm) Validate() {
	va := xtrememdw.Validator{}
	va.Make(rule)
}

func (rule *ItemComponentCategoryForm) APIParse(r *http.Request) {
	core.BaseForm{}.APIParse(r, &rule)
}
