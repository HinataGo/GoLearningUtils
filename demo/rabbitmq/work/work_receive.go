package work

import "SpikeSystem/demo/rabbitmq/simple"

func Receive1() {
	rabbitmq := simple.NewRabbitMQSimple("" + "rookie")
	rabbitmq.ConsumeSimple()
}
