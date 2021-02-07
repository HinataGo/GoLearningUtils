package simple

import "fmt"

func Publish() {
	rabbitmq := NewRabbitMQSimple("" + "rookie")
	rabbitmq.PublishSimple("hello ")
	fmt.Println("send success")
}
