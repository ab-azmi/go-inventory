package rabbitmq

import (
	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
	xtremerabbitmq "github.com/globalxtreme/go-core/v2/rabbitmq"
	"github.com/spf13/cobra"
	"service/internal/app/rabbitmq"
	"service/internal/pkg/config"
	"service/internal/pkg/constant"
)

type RabbitMQConsumerGlobalCommand struct{}

func (class *RabbitMQConsumerGlobalCommand) Command(cobraCmd *cobra.Command) {
	addCommand := cobra.Command{
		Use:  "rabbitmq:consumer-global",
		Long: "RabbitMQ Consumer Global",
		Run: func(cmd *cobra.Command, args []string) {
			xtremepkg.InitDevMode()

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

func (class *RabbitMQConsumerGlobalCommand) Handle() {
	xtremerabbitmq.Consume(xtremerabbitmq.RABBITMQ_CONNECTION_GLOBAL, []xtremerabbitmq.RabbitMQConsumeOpt{
		//{
		//	Exchange: "service.domain.feature.action.exchange",
		//	Consumer: &rabbitmq.TestingConsumer{},
		//},
		{
			Queue:    constant.RABBITMQ_QUEUE_SERVICE_DOMAIN_FEATURE_ACTION, // TODO: Hanya contoh. nanti langsung hapus saja
			Consumer: &rabbitmq.TestingRabbitMQConsumer{},
		},
	})
}
