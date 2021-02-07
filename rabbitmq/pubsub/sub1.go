package pubsub

import "GoLearningUtils/rabbitmq/RabbitMQ"

func Sub1() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("" +
		"newProduct")
	rabbitmq.RecieveSub()
}
