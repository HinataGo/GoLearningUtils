package work

import (
	"SpikeSystem/demo/rabbitmq/RabbitMQ"
)

func Receive2() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "rookie")
	rabbitmq.ConsumeSimple()
}
