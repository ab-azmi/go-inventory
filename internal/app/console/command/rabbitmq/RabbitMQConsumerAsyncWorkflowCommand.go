package rabbitmq

import (
	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
	xtremerabbitmq "github.com/globalxtreme/go-core/v2/rabbitmq"
	"github.com/spf13/cobra"
	"service/internal/app/rabbitmq"
	"service/internal/pkg/config"
)

type RabbitMQConsumerAsyncWorkflowCommand struct{}

func (class *RabbitMQConsumerAsyncWorkflowCommand) Command(cobraCmd *cobra.Command) {
	addCommand := cobra.Command{
		Use:  "rabbitmq:consumer-async-workflow",
		Long: "RabbitMQ Consumer Async Workflow",
		Run: func(cmd *cobra.Command, args []string) {
			xtremepkg.InitDevMode()
			xtremepkg.InitRedisPool()

			DBClose := config.InitDB()
			defer DBClose()

			rabbitmqConn := config.InitRabbitMQ()
			defer rabbitmqConn()

			dialRabbitMQConnClose := config.InitRabbitMQConnection()
			defer dialRabbitMQConnClose()

			logCleanup := xtremepkg.InitLogRPC()
			defer logCleanup()

			class.Handle()
		},
	}

	cobraCmd.AddCommand(&addCommand)
}

func (class *RabbitMQConsumerAsyncWorkflowCommand) Handle() {
	xtremerabbitmq.ConsumeWorkflow([]xtremerabbitmq.AsyncWorkflowConsumeOpt{
		{
			Queue:    "service.customer.convert.async-workflow-1", // TODO: Hanya contoh. nanti langsung hapus saja
			Consumer: &rabbitmq.TestingAsyncWorkflowExecutor{},
		},
	})
}
