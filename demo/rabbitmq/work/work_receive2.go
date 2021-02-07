package work

import (
	"GoLearningUtils/demo/rabbitmq/RabbitMQ"
)

func Receive2() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "rookie")
	rabbitmq.ConsumeSimple()
}
