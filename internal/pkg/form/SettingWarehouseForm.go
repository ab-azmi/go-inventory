package form

import (
	"net/http"
	"service/internal/pkg/core"

	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"
)

type SettingWarehouseForm struct {
	BranchOfficeId uint   `json:"branchOfficeId"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	ParentId       uint   `json:"parentId"`
}

func (rule *SettingWarehouseForm) Validate() {
	va := xtrememdw.Validator{}
	va.Make(rule)
}

func (rule *SettingWarehouseForm) APIParse(r *http.Request) {
	core.BaseForm{}.APIParse(r, &rule)
}
