package RabbitMQ

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// 话题模式
// 创建RabbitMQ实例
func NewRabbitMQTopic(exchangeName string, routingKey string) *RabbitMQ {
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

// 话题模式发送消息
func (r *RabbitMQ) PublishTopic(message string) {
	// 1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.ExChange,
		// 要改成topic
		"topic",
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

// 话题模式接受消息
// 要注意key,规则
// 其中“*”用于匹配一个单词，“#”用于匹配多个单词（可以是零个）
// 匹配 imooc.* 表示匹配 imooc.hello, 但是imooc.hello.one需要用imooc.#才能匹配到
func (r *RabbitMQ) RecieveTopic() {
	// 1.试探性创建交换机
	err := r.channel.ExchangeDeclare(
		r.ExChange,
		// 交换机类型
		"topic",
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
	r.checkErr(err, "Failed to declare a queue")

	// 绑定队列到 exchange 中
	err = r.channel.QueueBind(
		q.Name,
		// 在pub/sub模式下，这里的key要为空
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
