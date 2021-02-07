package topic

import "GoLearningUtils/demo/rabbitmq/RabbitMQ"

func AllTopicReceive() {
	TopicRec1 := RabbitMQ.NewRabbitMQTopic("exRookieTopic", "#")
	TopicRec1.RecieveTopic()
}
