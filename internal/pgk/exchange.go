package pgk

import (
	"github.com/streadway/amqp"
)

type Exchange struct {
	ex *amqp.Channel
}

func NewExchange(ch *amqp.Channel) *Exchange {
	return &Exchange{ch}
}

func (e *Exchange) CreateExchange(name, kind string, durable bool) error {
	return e.ex.ExchangeDeclare(name, kind, durable, false, false, false, nil)
}

func (e *Exchange) PublishExchange(name string, msg string) error {
	return e.ex.Publish(name, "", false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
}

func (e *Exchange) PublishExchangeRoute(name string, route, msg string) error {
	return e.ex.Publish(name, severityFrom(route), false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
}

func severityFrom(route string) string {
	if route == "" {
		return "anonymous.log"
	}

	return route
}
