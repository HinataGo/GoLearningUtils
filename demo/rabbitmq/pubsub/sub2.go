package pubsub

import "GoLearningUtils/demo/rabbitmq/RabbitMQ"

func Sub2() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("" +
		"newProduct")
	rabbitmq.RecieveSub()
}
