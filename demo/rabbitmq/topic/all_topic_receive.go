package topic

import "SpikeSystem/demo/rabbitmq/RabbitMQ"

func AllTopicReceive() {
	TopicRec1 := RabbitMQ.NewRabbitMQTopic("exRookieTopic", "#")
	TopicRec1.RecieveTopic()
}
