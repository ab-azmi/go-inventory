package form

import (
	"net/http"
)

// TODO: Hanya contoh. nanti langsung hapus saja
type TestingUploadForm struct {
	Request *http.Request
}

func (rule *TestingUploadForm) APIParse(r *http.Request) {
	rule.Request = r
}
