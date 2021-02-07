package topic

import "GoLearningUtils/rabbitmq/RabbitMQ"

func TopicReceive() {
	TopicRec2 := RabbitMQ.NewRabbitMQTopic("exRookieTopic", "rookie.*.two")
	TopicRec2.RecieveTopic()
}
