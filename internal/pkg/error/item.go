package error

import (
	"net/http"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
)

func ErrXtremeItemNumber(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to generate Number", internalMsg, false, nil)
}

/** --- Components --- */

/** --- Type --- */

func ErrXtremeItemComponentTypeGet(internalMsg string) {
	xtremeres.Error(http.StatusNotFound, "Type not found", internalMsg, false, nil)
}
func ErrXtremeItemComponentTypeCreate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to create type", internalMsg, false, nil)
}
func ErrXtremeItemComponentTypeUpdate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to update type", internalMsg, false, nil)
}
func ErrXtremeItemComponentTypeDelete(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to delete type", internalMsg, false, nil)
}

/** --- Brand --- */

func ErrXtremeItemComponentBrandGet(internalMsg string) {
	xtremeres.Error(http.StatusNotFound, "Brand not found", internalMsg, false, nil)
}
func ErrXtremeItemComponentBrandCreate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to create brand", internalMsg, false, nil)
}
func ErrXtremeItemComponentBrandUpdate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to update brand", internalMsg, false, nil)
}
func ErrXtremeItemComponentBrandDelete(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to delete brand", internalMsg, false, nil)
}

/** --- Category --- */

func ErrXtremeItemComponentCategoryGet(internalMsg string) {
	xtremeres.Error(http.StatusNotFound, "Category not found", internalMsg, false, nil)
}
func ErrXtremeItemComponentCategoryCreate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to create category", internalMsg, false, nil)
}
func ErrXtremeItemComponentCategoryUpdate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to update category", internalMsg, false, nil)
}
func ErrXtremeItemComponentCategoryDelete(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to delete category", internalMsg, false, nil)
}

/** --- Unit --- */

func ErrXtremeItemComponentUnitGet(internalMsg string) {
	xtremeres.Error(http.StatusNotFound, "Unit not found", internalMsg, false, nil)
}
func ErrXtremeItemComponentUnitCreate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to create unit", internalMsg, false, nil)
}
func ErrXtremeItemComponentUnitUpdate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to update unit", internalMsg, false, nil)
}
func ErrXtremeItemComponentUnitDelete(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to delete unit", internalMsg, false, nil)
}
