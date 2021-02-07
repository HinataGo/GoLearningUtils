package simple

import "GoLearningUtils/rabbitmq/RabbitMQ"

func Receive() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "rookie")
	rabbitmq.ConsumeSimple()
}
