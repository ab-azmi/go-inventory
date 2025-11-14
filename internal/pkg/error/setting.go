package error

import (
	"net/http"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
)

/** --- Warehouse --- */

func ErrXtremeSettingWarehouseGet(internalMsg string) {
	xtremeres.Error(http.StatusNotFound, "Warehouse not found", internalMsg, false, nil)
}
func ErrXtremeSettingWarehouseCreate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to create warehouse", internalMsg, false, nil)
}
func ErrXtremeSettingWarehouseUpdate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to update warehouse", internalMsg, false, nil)
}
func ErrXtremeSettingWarehouseDelete(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to delete warehouse", internalMsg, false, nil)
}
