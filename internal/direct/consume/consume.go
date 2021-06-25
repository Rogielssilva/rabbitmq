package main

import (
	"log"
	"rabbitMq/internal"
	"rabbitMq/internal/pgk"

	"github.com/streadway/amqp"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	internal.HandleError(err, "failed while connecting to the server")
	defer conn.Close()

	channel, err := conn.Channel()
	internal.HandleError(err, "failed to create consume channel")

	defer conn.Close()

	q := pgk.NewQueueInstance(channel, "test01", "", false, false)
	_, err = q.CreateQueue()
	internal.HandleError(err, "error to create a queue")

	msgs, err := q.Consume(true)
	internal.HandleError(err, "error to consume message")

	wait := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("message consumed %s\n", d.Body)
		}
	}()

	log.Print("waiting...")

	<-wait
}
