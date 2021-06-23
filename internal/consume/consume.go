package main

import (
	"log"
	"rabbitMq/internal/pgk"

	"github.com/streadway/amqp"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	pgk.HandleError(err, "failed while connecting to the server")
	defer conn.Close()

	channel, err := conn.Channel()
	pgk.HandleError(err, "failed to create consume channel")

	defer conn.Close()

	q := pgk.NewQueueInstance("test01", channel)
	q.CreateQueue()

	msgs := q.Consume()
	wait := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("message consumed %s\n", d.Body)
		}
	}()

	log.Print("waiting...")

	<-wait
}
