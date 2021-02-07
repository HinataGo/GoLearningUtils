package pubsub

import "GoLearningUtils/rabbitmq/RabbitMQ"

func Sub2() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("" +
		"newProduct")
	rabbitmq.RecieveSub()
}
