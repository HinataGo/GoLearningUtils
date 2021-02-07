package simple

func Receive() {
	rabbitmq := NewRabbitMQSimple("" + "rookie")
	rabbitmq.ConsumeSimple()
}
