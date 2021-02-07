package work

import "SpikeSystem/demo/rabbitmq/simple"

func Receive2() {
	rabbitmq := simple.NewRabbitMQSimple("" + "rookie")
	rabbitmq.ConsumeSimple()
}
