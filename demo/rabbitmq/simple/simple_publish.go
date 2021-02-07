package simple

import (
	"fmt"

	"SpikeSystem/demo/rabbitmq/RabbitMQ"
)

func Publish() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "rookie")
	rabbitmq.PublishSimple("hello ")
	fmt.Println("send success")
}
