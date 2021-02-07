package work

import (
	"GoLearningUtils/rabbitmq/RabbitMQ"
)

func Receive2() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "rookie")
	rabbitmq.ConsumeSimple()
}
