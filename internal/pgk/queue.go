package pgk

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

type Queue struct {
	name string
	ch   *amqp.Channel
}

func HandleError(err error, message string) {
	if err != nil {
		log.Fatalf("messge: %v error: %v\n", message, err)
	}
}

func NewQueueInstance(name string, ch *amqp.Channel) *Queue {
	return &Queue{name: name, ch: ch}
}

func (q *Queue) CreateQueue() (amqp.Queue, error) {
	return q.ch.QueueDeclare(q.name, false, false, false, false, nil)
}

func (q *Queue) Consume() <-chan amqp.Delivery {
	msgs, err := q.ch.Consume(q.name, "", true, false, false, false, nil)
	HandleError(err, "error to consume message")
	return msgs
}

func (q *Queue) Publish(msg string) {
	err := q.ch.Publish("", q.name, false, false, amqp.Publishing{
		Headers:         nil,
		ContentType:     "text/plain",
		ContentEncoding: "",
		DeliveryMode:    0,
		Priority:        0,
		CorrelationId:   "",
		ReplyTo:         "",
		Expiration:      "",
		MessageId:       "",
		Timestamp:       time.Time{},
		Type:            "",
		UserId:          "",
		AppId:           "",
		Body:            []byte(msg),
	})

	HandleError(err, "failed while publishing a message")
}
