package setting

import (
	"net/http"
	"service/internal/pkg/core"

	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"
)

type ItemBrandForm struct {
	Name string `json:"name"`
}

func (rule *ItemBrandForm) Validate() {
	va := xtrememdw.Validator{}
	va.Make(rule)
}

func (rule *ItemBrandForm) APIParse(r *http.Request) {
	core.BaseForm{}.APIParse(r, &rule)
}
