package port

import (
	"gorm.io/gorm"
	"net/url"
	"service/internal/pkg/model"
)

/** --- ACTIVITY --- */

type ActivityRepository interface {
	Find(parameters url.Values) ([]model.Activity, interface{}, error)
}

/** --- SETTING --- */

type SettingItemBrandRepository interface {
	FirstById(id uint, args ...func(query *gorm.DB) *gorm.DB) model.ItemComponentBrand
}
