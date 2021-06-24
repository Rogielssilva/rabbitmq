package main

import (
	"bytes"
	"log"
	"rabbitMq/internal"
	"rabbitMq/internal/pgk"
	"time"

	"github.com/streadway/amqp"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	internal.HandleError(err, "error while setting up server")

	defer conn.Close()

	channel, err := conn.Channel()
	internal.HandleError(err, "error while getting channel")

	q := pgk.NewQueueInstance(channel, "test02", true)

	_, err = q.CreateQueue()
	internal.HandleError(err, "error to create queue")

	err = q.QosConfig()
	internal.HandleError(err, "error to config QoS")

	msgs, err := q.Consume(false)
	internal.HandleError(err, "error to consume message")

	wait := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	log.Print("waiting...")

	<-wait
}
