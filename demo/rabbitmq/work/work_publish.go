package work

import (
	"fmt"
	"strconv"
	"time"

	"SpikeSystem/demo/rabbitmq/simple"
)

func WorkPublish() {
	rabbitmq := simple.NewRabbitMQSimple("" + "rookie")
	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Hello rookie!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
