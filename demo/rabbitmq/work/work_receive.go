package work

import (
	"SpikeSystem/demo/rabbitmq/RabbitMQ"
)

func Receive1() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "rookie")
	rabbitmq.ConsumeSimple()
}
