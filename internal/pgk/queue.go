package pgk

import (
	"time"

	"github.com/streadway/amqp"
)

type Queue struct {
	name    string
	ch      *amqp.Channel
	durable bool
}

func NewQueueInstance(ch *amqp.Channel, name string, durable bool) *Queue {
	return &Queue{name: name, ch: ch, durable: durable}
}

func (q *Queue) CreateQueue() (amqp.Queue, error) {
	return q.ch.QueueDeclare(q.name, q.durable, false, false, false, nil)
}

func (q *Queue) Consume(autoACk bool) (<-chan amqp.Delivery, error) {
	return q.ch.Consume(q.name, "", autoACk, false, false, false, nil)
}

func (q *Queue) QosConfig() error {
	return q.ch.Qos(1, 0, false)
}

func (q *Queue) Publish(msg string) error {
	return q.ch.Publish("", q.name, false, false, amqp.Publishing{
		Headers:         nil,
		ContentType:     "text/plain",
		ContentEncoding: "",
		DeliveryMode:    amqp.Persistent,
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
}
