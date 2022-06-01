package main

import (
	"github.com/streadway/amqp"
	"log"
)

func FailOnError1(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// 消费者

func main(){
	// 连接rabbit的ip和port
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError1(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError1(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // 队列name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	FailOnError1(err, "Failed to declare a queue")

	// 定义一个消费者
	// msgs是<-chan Delivery类型的channel
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError1(err, "Failed to register a consumer")

	forever := make(chan bool)

	// 开启一个goroutine,遍历管道的值，并遍历
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}