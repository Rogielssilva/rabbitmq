package main

import (
	"log"
	"os"
	"rabbitMq/internal"
	"rabbitMq/internal/pgk"
	"strings"

	"github.com/streadway/amqp"
)

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	internal.HandleError(err, "error while setting up server")

	defer conn.Close()

	channel, err := conn.Channel()
	internal.HandleError(err, "error while getting channel")

	exName := "logs_types"
	e := pgk.NewExchange(channel)
	err = e.CreateExchange(exName, "topic", true)
	internal.HandleError(err, "error to create exchange")

	// q := pgk.NewQueueInstance(channel, "test_exchange", exName, false, true)
	// _, err = q.CreateQueue()
	// internal.HandleError(err, "error to create a queue")

	// err = q.QueueBind()
	// internal.HandleError(err, "error for biding")

	b := bodyFrom(os.Args)
	err = e.PublishExchangeRoute(exName, os.Args[1], b)

	internal.HandleError(err, "error while publish message")

	log.Printf("send %s", b)

}
