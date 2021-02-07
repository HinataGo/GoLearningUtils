package simple

import "SpikeSystem/demo/rabbitmq/RabbitMQ"

func Receive() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "rookie")
	rabbitmq.ConsumeSimple()
}
