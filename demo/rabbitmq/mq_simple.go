package rabbitmq

import (
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

func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	rabbitMQ := &RabbitMQ{
		QueueName: exchange,
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
func (r *RabbitMQ) Destory() {
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

func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
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
