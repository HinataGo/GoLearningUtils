package RabbitMQ

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// url 格式 amqp:.../vhost
const mqUrl = "amqp://rookie:rookie@127.0.0.1:5672/sell"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// queue name
	QueueName string
	// switch 交换机
	ExChange string
	// key
	Key string
	// link msg
	MqUrl string
}

// NewRabbitMQSimple: 新建一个简单队列,  简单模式 工作模式也用这个(详细区别md文档分析)
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	rabbitMQ := &RabbitMQ{
		QueueName: queueName,
		ExChange:  exchange,
		Key:       key,
		MqUrl:     mqUrl,
	}
	var err error
	rabbitMQ.conn, err = amqp.Dial(rabbitMQ.MqUrl)
	rabbitMQ.checkErr(err, "创建连接错误")
	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	rabbitMQ.checkErr(err, "获取channel 失败")
	return rabbitMQ
}

// 断开channel 和 connection
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理函数
func (r *RabbitMQ) checkErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		// panic(fmt.Sprintf("%s:%s",msg,err))
	}
}

func (r *RabbitMQ) PublishSimple(msg string) {
	// 先申请队列
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// 消息持久化
		false,
		// 是否自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		log.Fatalf("%s", err)
	}
	// 发送消息队列中
	r.channel.Publish(
		r.ExChange,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentEncoding: "text/plain",
			Body:            []byte(msg),
		})
}

// simple 模式下消费者
func (r *RabbitMQ) ConsumeSimple() {
	// 1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞处理
		false,
		// 额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	// 接收消息
	msg, err := r.channel.Consume(
		q.Name, // queue
		// 用来区分多个消费者
		"", // consumer
		// 是否自动应答
		true, // auto-ack
		// 是否独有
		false, // exclusive
		// 设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		// 列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	// 启用协程处理消息
	go func() {
		for d := range msg {
			// 消息逻辑处理，可以自行设计逻辑
			log.Printf("Received a message: %s", d.Body)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
