package main

import (
	"log"
	"rabbitMq/internal"
	"rabbitMq/internal/pgk"

	"github.com/streadway/amqp"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	internal.HandleError(err, "error while setting up server")

	defer conn.Close()

	channel, err := conn.Channel()
	internal.HandleError(err, "error while getting channel")

	exName := "myexch"
	e := pgk.NewExchange(channel)
	err = e.CreateExchange(exName, "fanout", true)
	internal.HandleError(err, "error to create exchange")

	q := pgk.NewQueueInstance(channel, "", exName, true, false)
	_, err = q.CreateQueue()
	internal.HandleError(err, "error to create a queue")

	err = q.QueueBind()
	internal.HandleError(err, "error for biding")

	wait := make(chan bool)

	msgs, err := q.Consume(true)
	internal.HandleError(err, "error to consume messages")

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf("waiting....")
	<-wait

}
