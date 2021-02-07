package work

import (
	"GoLearningUtils/rabbitmq/RabbitMQ"
)

func Receive1() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "rookie")
	rabbitmq.ConsumeSimple()
}
