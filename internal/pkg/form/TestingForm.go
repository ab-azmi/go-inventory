package form

import (
	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"
	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	xtremerabbitmq "github.com/globalxtreme/go-core/v2/rabbitmq"
	"net/http"
	"service/internal/pkg/core"
)

// TODO: Hanya contoh. nanti langsung hapus saja
type TestingForm struct {
	Name   string   `json:"name"`
	Subs   []string `json:"subs" validate:"required"`
	Action string   `json:"action"`
	Status struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Subs []struct {
			Name string `json:"name"`
		} `json:"subs"`
		Types []string `json:"types"`
	} `json:"status"`
}

func (rule *TestingForm) Validate() {
	va := xtrememdw.Validator{}
	va.Make(rule)
}

func (rule *TestingForm) APIParse(r *http.Request) {
	core.BaseForm{}.APIParse(r, &rule)
}

func (rule *TestingForm) AsyncWorkflowParse(payload interface{}) error {
	return core.BaseForm{}.AsyncWorkflowParse(payload, &rule)
}

func (rule *TestingForm) RabbitMQParse(message xtrememodel.RabbitMQMessage) error {
	return core.BaseForm{}.RabbitMQParse(message, &rule)
}

func (rule *TestingForm) RabbitMQProcessedParse(message xtrememodel.RabbitMQMessage) (*xtremerabbitmq.RabbitMQDeliveryResponse, error) {
	return core.BaseForm{}.RabbitMQProcessedParse(message, &rule)
}
