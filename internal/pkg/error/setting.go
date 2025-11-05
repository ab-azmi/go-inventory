package error

import (
	"net/http"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
)

/** --- Type --- */

func ErrXtremeSettingItemTypeGet(internalMsg string) {
	xtremeres.Error(http.StatusNotFound, "Type not found", internalMsg, false, nil)
}
func ErrXtremeSettingItemTypeCreate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to create type", internalMsg, false, nil)
}
func ErrXtremeSettingItemTypeUpdate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to update type", internalMsg, false, nil)
}
func ErrXtremeSettingItemTypeDelete(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to delete type", internalMsg, false, nil)
}

/** --- Brand --- */

func ErrXtremeSettingItemBrandGet(internalMsg string) {
	xtremeres.Error(http.StatusNotFound, "Brand not found", internalMsg, false, nil)
}
func ErrXtremeSettingItemBrandCreate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to create brand", internalMsg, false, nil)
}
func ErrXtremeSettingItemBrandUpdate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to update brand", internalMsg, false, nil)
}
func ErrXtremeSettingItemBrandDelete(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to delete brand", internalMsg, false, nil)
}

/** --- Category --- */

func ErrXtremeSettingItemCategoryGet(internalMsg string) {
	xtremeres.Error(http.StatusNotFound, "Category not found", internalMsg, false, nil)
}
func ErrXtremeSettingItemCategoryCreate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to create category", internalMsg, false, nil)
}
func ErrXtremeSettingItemCategoryUpdate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to update category", internalMsg, false, nil)
}
func ErrXtremeSettingItemCategoryDelete(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to delete category", internalMsg, false, nil)
}

/** --- Unit --- */

func ErrXtremeSettingItemUnitGet(internalMsg string) {
	xtremeres.Error(http.StatusNotFound, "Unit not found", internalMsg, false, nil)
}
func ErrXtremeSettingItemUnitCreate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to create unit", internalMsg, false, nil)
}
func ErrXtremeSettingItemUnitUpdate(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to update unit", internalMsg, false, nil)
}
func ErrXtremeSettingItemUnitDelete(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to delete unit", internalMsg, false, nil)
}
