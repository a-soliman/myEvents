package amqp

import (
	"encoding/json"
	"fmt"

	"github.com/a-soliman/projects/myEvents/src/lib/msgqueue"
	"github.com/streadway/amqp"
)

type amqpEventEmitter struct {
	connection *amqp.Connection
}

func (a *amqpEventEmitter) setup() error {
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	return channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)
}

func (a *amqpEventEmitter) Emit(event msgqueue.Event) error {
	fmt.Println("Emmit")
	jsonDoc, err := json.Marshal(event)
	if err != nil {
		return err
	}

	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	msg := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": event.EventName()},
		Body:        jsonDoc,
		ContentType: "application/json",
	}
	return channel.Publish(
		"events",
		event.EventName(),
		false,
		false,
		msg,
	)
}

// NewAMQPEventEmitter initialize and return a pointer to an emitter
func NewAMQPEventEmitter(conn *amqp.Connection) (msgqueue.EventEmitter, error) {
	emitter := &amqpEventEmitter{
		connection: conn,
	}
	err := emitter.setup()
	if err != nil {
		return nil, err
	}
	return emitter, nil
}
