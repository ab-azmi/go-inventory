package port

import (
	"service/internal/pkg/form"
	"service/internal/pkg/model"
)

type SettingItemBrandService interface {
	Create(form form.SettingForm) model.ItemComponentBrand
}
