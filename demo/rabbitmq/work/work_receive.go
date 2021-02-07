package work

import (
	"GoLearningUtils/demo/rabbitmq/RabbitMQ"
)

func Receive1() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "rookie")
	rabbitmq.ConsumeSimple()
}
