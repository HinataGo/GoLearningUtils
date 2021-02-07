package simple

import (
	"fmt"

	"GoLearningUtils/rabbitmq/RabbitMQ"
)

func Publish() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "rookie")
	rabbitmq.PublishSimple("hello ")
	fmt.Println("send success")
}
