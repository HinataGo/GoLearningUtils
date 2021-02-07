package pubsub

import "GoLearningUtils/demo/rabbitmq/RabbitMQ"

func Sub1() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("" +
		"newProduct")
	rabbitmq.RecieveSub()
}
