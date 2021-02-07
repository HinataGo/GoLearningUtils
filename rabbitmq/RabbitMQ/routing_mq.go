package RabbitMQ

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// 路由模式
// 创建RabbitMQ实例
func NewRabbitMQRouting(exchangeName string, routingKey string) *RabbitMQ {
	// 创建RabbitMQ实例
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	var err error
	// 获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.checkErr(err, "failed to connect rabbitmq!")
	// 获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.checkErr(err, "failed to open a channel")
	return rabbitmq
}

// 路由模式发送消息
func (r *RabbitMQ) PublishRouting(message string) {
	// 1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.ExChange,
		// 要改成direct
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	r.checkErr(err, "Failed to declare an excha"+
		"nge")

	// 2.发送消息
	err = r.channel.Publish(
		r.ExChange,
		// 要设置
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// 路由模式接受消息
func (r *RabbitMQ) RecieveRouting() {
	// 1.试探性创建交换机
	err := r.channel.ExchangeDeclare(
		r.ExChange,
		// 交换机类型
		"direct",
		true,
		false,
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
		// 需要绑定key
		r.Key,
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
