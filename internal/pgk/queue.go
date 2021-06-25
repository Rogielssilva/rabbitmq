package pgk

import (
	"time"

	"github.com/streadway/amqp"
)

type Queue struct {
	name      string
	ch        *amqp.Channel
	durable   bool
	exclusive bool
	exchange  string
}

func NewQueueInstance(ch *amqp.Channel, name, exchange string, durable, exclusive bool) *Queue {
	return &Queue{ch: ch, name: name, exchange: exchange, durable: durable, exclusive: exclusive}
}

func (q *Queue) CreateQueue() (amqp.Queue, error) {
	return q.ch.QueueDeclare(q.name, q.durable, false, q.exclusive, false, nil)
}

func (q *Queue) Consume(autoACk bool) (<-chan amqp.Delivery, error) {
	return q.ch.Consume(q.name, "", autoACk, false, false, false, nil)
}

func (q *Queue) QosConfig() error {
	return q.ch.Qos(1, 0, false)
}

func (q *Queue) Publish(msg string) error {
	return q.ch.Publish(q.exchange, q.name, false, false, amqp.Publishing{
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

func (q *Queue) QueueBind() error {
	return q.ch.QueueBind(q.name, "", q.exchange, false, nil)
}
