package pubsub

import "SpikeSystem/demo/rabbitmq/RabbitMQ"

func Sub2() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("" +
		"newProduct")
	rabbitmq.RecieveSub()
}
