package amqp

import (
	"context"
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Message struct {
	msg *amqp.Delivery
}

func (t *Message) Ack() error {
	return t.msg.Ack(false)
}
func (t *Message) Nack() error {
	return t.msg.Nack(false, false)
}

func (t *Message) Data() []byte {
	return t.msg.Body
}

type Consumer interface {
	Init() error
	Close() error
	Subscribe(context.Context, Queue) (chan Message, chan error)
}

type consumer struct {
	config  *Config
	client  *amqp.Connection
	channel *amqp.Channel

	errChan chan error
	msgChan chan Message
}

func NewConsumer(config *Config) Consumer {
	return &consumer{
		config:  config,
		errChan: make(chan error),
		msgChan: make(chan Message),
	}
}

var _ Consumer = (*consumer)(nil)

func (t *consumer) Init() error {
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

func (t *consumer) Close() error {
	if err := t.channel.Close(); err != nil {
		return errors.Wrap(err, "amqp.Channel.Close")
	}
	return errors.Wrap(t.client.Close(), "amqp.Close")
}

func (t *consumer) Subscribe(ctx context.Context, queue Queue) (chan Message, chan error) {
	go t.subscribe(ctx, queue)

	return t.msgChan, t.errChan
}

func (t *consumer) subscribe(ctx context.Context, queue Queue) {
	q, err := t.channel.QueueDeclare(
		queue.String(),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		t.errChan <- err
		return
	}

	msg, err := t.channel.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		t.errChan <- err
		return
	}

	for {
		select {
		case m := <-msg:
			t.msgChan <- Message{
				&m,
			}
		case <-ctx.Done():
			if err = ctx.Err(); err != nil {
				t.errChan <- err
			}
			return
		}
	}
}
