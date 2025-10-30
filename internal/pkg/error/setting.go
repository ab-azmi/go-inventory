package error

import (
	"fmt"
	"net/http"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
)

func ErrXtremeSettingGet(name string, internalMsg string) {
	message := fmt.Sprintf("%s not found", name)
	xtremeres.Error(http.StatusNotFound, message, internalMsg, false, nil)
}

func ErrXtremeSettingCreate(name string, internalMsg string) {
	message := fmt.Sprintf("Unable to create %s", name)
	xtremeres.Error(http.StatusInternalServerError, message, internalMsg, false, nil)
}

func ErrXtremeSettingUpdate(name string, internalMsg string) {
	message := fmt.Sprintf("Unable to update %s", name)
	xtremeres.Error(http.StatusInternalServerError, message, internalMsg, false, nil)
}

func ErrXtremeSettingDelete(name string, internalMsg string) {
	message := fmt.Sprintf("Unable to delete %s", name)
	xtremeres.Error(http.StatusInternalServerError, message, internalMsg, false, nil)
}
