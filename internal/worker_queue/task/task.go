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

	q := pgk.NewQueueInstance(channel, "test02", false)

	b := bodyFrom(os.Args)

	err = q.Publish(b)
	internal.HandleError(err, "error to consume message")

	log.Printf("send %s", b)

}
