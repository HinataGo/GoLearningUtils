package routing

import (
	"fmt"
	"strconv"
	"time"

	"GoLearningUtils/demo/rabbitmq/RabbitMQ"
)

func routing() {
	route1 := RabbitMQ.NewRabbitMQRouting("rookie", "rookie_one")
	route2 := RabbitMQ.NewRabbitMQRouting("rookie", "rookie_two")
	for i := 0; i <= 10; i++ {
		route1.PublishRouting("Hello rookie one!" + strconv.Itoa(i))
		route2.PublishRouting("Hello rookie Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
