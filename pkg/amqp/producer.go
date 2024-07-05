package amqp

import (
	"context"
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer interface {
	Init() error
	Close() error
	Publish(ctx context.Context, queue Queue, ch chan []byte, stop chan struct{}) error
}

type producer struct {
	config  *Config
	client  *amqp.Connection
	channel *amqp.Channel
	stop    chan struct{}
}

var _ Producer = (*producer)(nil)

func NewProducer(cfg *Config) (Producer, chan struct{}) {
	ch := make(chan struct{})
	return &producer{
		config: cfg,
		stop:   ch,
	}, ch
}

func (t *producer) Init() error {
	conn, err := amqp.Dial(t.config.DSN)
	if err != nil {
		return errors.Wrap(err, "amqp.Dial")
	}

	t.client = conn

	channel, err := t.client.Channel()
	if err != nil {
		return errors.Wrap(err, "amqp.Channel")
	}

	t.channel = channel

	return nil
}

func (t *producer) Close() error {
	return errors.Wrap(t.client.Close(), "client.Close")
}

func (t *producer) Publish(ctx context.Context, queue Queue,
	ch chan []byte,
	stop chan struct{}) error {
	q, err := t.channel.QueueDeclare(
		queue.String(),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "channel.QueueDeclare")
	}

	go func() {
		for {
			select {
			case data := <-ch:
				err = t.channel.PublishWithContext(
					ctx,
					"",
					q.Name,
					false,
					false,
					amqp.Publishing{
						Body: data,
					},
				)
			case <-stop:
				t.stop <- struct{}{}
				return
			}
		}
	}()

	return nil
}
