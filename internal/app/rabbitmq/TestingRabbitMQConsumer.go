package rabbitmq

import (
	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	"service/internal/pkg/core"
	form2 "service/internal/pkg/form"
	"service/internal/testing/service"
	"sync"
)

// TODO: Hanya contoh. nanti langsung hapus saja
type TestingRabbitMQConsumer struct {
	mutex sync.Mutex
}

func (consume *TestingRabbitMQConsumer) Consume(message xtrememodel.RabbitMQMessage) (interface{}, error, []byte) {
	consume.mutex.Lock()
	defer consume.mutex.Unlock()

	return core.RabbitMQErrorHandler(func() (interface{}, error) {
		form := form2.TestingForm{}
		err := form.RabbitMQParse(message)
		if err != nil {
			return nil, err
		}

		srv := service.NewTestingService()
		testing := srv.CreateConsumer(form)

		// Taruh di parser
		result := map[string]interface{}{
			"name":     testing.Name,
			"totalSub": len(testing.Subs),
		}

		return result, nil
	})
}
