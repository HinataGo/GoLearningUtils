package topic

import (
	"fmt"
	"strconv"
	"time"

	"SpikeSystem/demo/rabbitmq/RabbitMQ"
)

func TopicPub() {
	Topic1 := RabbitMQ.NewRabbitMQTopic("exRookieTopic", "rookie.topic.one")
	Topic2 := RabbitMQ.NewRabbitMQTopic("exRookieTopic", "rookie.topic.two")
	for i := 0; i <= 10; i++ {
		Topic1.PublishTopic("Hello rookie topic one!" + strconv.Itoa(i))
		Topic2.PublishTopic("Hello rookie topic Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
