package main

import (
	"github.com/gpmgo/gopm/modules/log"
	"fmt"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	log.Fatal("%s, %s", msg, err)
	panic(fmt.Sprintf("%s, %s", msg, err))
}

func main() {
	conn, err := amqp.Dial("amqp://xie:Fs9006@192.168.102.180:5672")
	failOnError(err, "failed connect to rabbitMQ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msg, err := ch.Consume(
		q.Name, "",
		true,
		false,
		false,
		false,
		nil)
	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)

	go func() {
		for d := range msg {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Print(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
