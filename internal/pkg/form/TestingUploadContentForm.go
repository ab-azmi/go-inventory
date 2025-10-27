package form

import (
	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"
	"net/http"
	"service/internal/pkg/core"
)

// TODO: Hanya contoh. nanti langsung hapus saja
type TestingUploadContentForm struct {
	Content string `json:"content" validate:"required"`
}

func (rule *TestingUploadContentForm) Validate() {
	va := xtrememdw.Validator{}
	va.Make(rule)
}

func (rule *TestingUploadContentForm) APIParse(r *http.Request) {
	core.BaseForm{}.APIParse(r, &rule)
}
