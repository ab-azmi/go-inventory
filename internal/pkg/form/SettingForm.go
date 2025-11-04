package form

import (
	"net/http"
	"service/internal/pkg/core"

	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"
)

type SettingForm struct {
	Name string `json:"name"`
}

func (rule *SettingForm) Validate() {
	va := xtrememdw.Validator{}
	va.Make(rule)
}

func (rule *SettingForm) APIParse(r *http.Request) {
	core.BaseForm{}.APIParse(r, &rule)
}
