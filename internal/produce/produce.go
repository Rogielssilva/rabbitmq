package main

import (
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

	defer channel.Close()

	q := pgk.NewQueueInstance(channel, "test01", false)
	_, err = q.CreateQueue()
	internal.HandleError(err, "error to create a queue")
	q.Publish("hello world!!")

}
