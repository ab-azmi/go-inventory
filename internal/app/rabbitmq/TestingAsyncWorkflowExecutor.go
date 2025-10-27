package rabbitmq

import (
	"encoding/json"
	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
	xtremerabbitmq "github.com/globalxtreme/go-core/v2/rabbitmq"
	"service/internal/pkg/core"
	form2 "service/internal/pkg/form"
	"sync"
)

// TODO: Hanya contoh. nanti langsung hapus saja
type TestingAsyncWorkflowExecutor struct {
	xtremerabbitmq.AsyncWorkflowConsumerBase

	mutex sync.Mutex
	form  *form2.TestingForm
}

func (c *TestingAsyncWorkflowExecutor) Consume(payload interface{}) (interface{}, error, []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return core.RabbitMQErrorHandler(func() (interface{}, error) {
		c.form = &form2.TestingForm{}
		err := c.form.AsyncWorkflowParse(payload)
		if err != nil {
			return nil, err
		}

		xtremepkg.LogInfo("Step 1")
		xtremepkg.LogInfo(c.GetReferenceId())
		xtremepkg.LogInfo(c.GetReferenceType())
		formJson, _ := json.Marshal(c.form)
		xtremepkg.LogInfo(string(formJson))

		return c.Response(payload), nil
	})
}

func (c *TestingAsyncWorkflowExecutor) Response(payload interface{}, data ...interface{}) interface{} {
	if c.form == nil {
		c.form = &form2.TestingForm{}
		err := c.form.AsyncWorkflowParse(payload)
		if err != nil {
			return nil
		}
	}

	return map[string]interface{}{
		"name": c.form.Name,
		"subs": c.form.Subs,
	}
}

func (c *TestingAsyncWorkflowExecutor) ForwardPayload() []xtremerabbitmq.AsyncWorkflowForwardPayloadResult {
	return []xtremerabbitmq.AsyncWorkflowForwardPayloadResult{
		{
			Queue: "service.customer.convert.async-workflow-4",
			Payload: map[string]interface{}{
				"action": "testing forward message for step 4",
				"status": map[string]interface{}{
					"id":   1,
					"name": "testing status",
				},
			},
		},
		{
			Queue: "service.customer.convert.async-workflow-3",
			Payload: map[string]interface{}{
				"action": "testing forward message for step 3",
			},
		},
	}
}
