package topic

import "GoLearningUtils/demo/rabbitmq/RabbitMQ"

func TopicReceive() {
	TopicRec2 := RabbitMQ.NewRabbitMQTopic("exRookieTopic", "rookie.*.two")
	TopicRec2.RecieveTopic()
}
