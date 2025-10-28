package setting

import (
	"net/http"
	"service/internal/pkg/core"

	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"
)

type ItemTypeForm struct {
	Name string `json:"name"`
}

func (rule *ItemTypeForm) Validate() {
	va := xtrememdw.Validator{}
	va.Make(rule)
}

func (rule *ItemTypeForm) APIParse(r *http.Request) {
	core.BaseForm{}.APIParse(r, &rule)
}
