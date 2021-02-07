package simple

import "GoLearningUtils/demo/rabbitmq/RabbitMQ"

func Receive() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "rookie")
	rabbitmq.ConsumeSimple()
}
