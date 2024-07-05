package amqp

import (
	"context"
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer interface {
	Init() error
	Close() error
	Subscribe(context.Context, Queue) (chan []byte, chan error)
}

type consumer struct {
	config  *Config
	client  *amqp.Connection
	channel *amqp.Channel

	errChan chan error
	msgChan chan []byte
}

func NewConsumer(config *Config) Consumer {
	return &consumer{
		config:  config,
		errChan: make(chan error),
		msgChan: make(chan []byte),
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

func (t *consumer) Subscribe(ctx context.Context, queue Queue) (chan []byte, chan error) {
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
		true,
		false,
		false,
		true,
		nil)
	if err != nil {
		t.errChan <- err
		return
	}

	for {
		select {
		case m := <-msg:
			t.msgChan <- m.Body
		case <-ctx.Done():
			if err = ctx.Err(); err != nil {
				t.errChan <- err
			}
			return
		}
	}
}
