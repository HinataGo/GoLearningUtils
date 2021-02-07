package RabbitMQ

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// 订阅模式创建RabbitMQ实例
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	// 创建RabbitMQ实例
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	var err error
	// 获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.checkErr(err, "failed to connect rabbitmq!")
	// 获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.checkErr(err, "failed to open a channel")
	return rabbitmq
}

// 订阅模式生产
func (r *RabbitMQ) PublishPub(message string) {
	// 1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.ExChange,
		"fanout",
		true,
		false,
		// true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		false,
		nil,
	)

	r.checkErr(err, "Failed to declare an excha"+
		"nge")

	// 2.发送消息
	err = r.channel.Publish(
		r.ExChange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// 订阅模式消费端代码
func (r *RabbitMQ) RecieveSub() {
	// 1.试探性创建交换机
	err := r.channel.ExchangeDeclare(
		r.ExChange,
		// 交换机类型
		"fanout",
		true,
		false,
		// YES表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	r.checkErr(err, "Failed to declare an exch"+
		"ange")
	// 2.试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		"", // 随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	r.checkErr(err, "Failed to declare a queue")

	// 绑定队列到 exchange 中
	err = r.channel.QueueBind(
		q.Name,
		// 在pub/sub模式下，这里的key要为空
		"",
		r.ExChange,
		false,
		nil)

	// 消费消息
	messages, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range messages {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	fmt.Println("退出请按 CTRL+C")
	<-forever
}
