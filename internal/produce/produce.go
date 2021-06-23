package main

import (
	"rabbitMq/internal/pgk"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	pgk.HandleError(err, "error while setting up server")

	defer conn.Close()

	channel, err := conn.Channel()
	pgk.HandleError(err, "error while getting channel")

	defer channel.Close()

	q := pgk.NewQueueInstance("test01", channel)
	q.CreateQueue()
	q.Publish("hello world!!")

}
