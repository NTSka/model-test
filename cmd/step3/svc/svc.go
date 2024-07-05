package svc

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"os"
	"path"
	"test-model/pkg/amqp"
	"test-model/pkg/helpers"
	"test-model/pkg/proto/event"
	"time"
)

type Service interface {
	Run(ctx context.Context) error
	Close() error
}

type svc struct {
	config    *Config
	consumer  amqp.Consumer
	producer  amqp.Producer
	counter   int
	startTime time.Time
}

func NewSvc(config *Config, consumer amqp.Consumer, producer amqp.Producer) Service {
	return &svc{
		config:   config,
		consumer: consumer,
		producer: producer,
	}
}

var _ Service = (*svc)(nil)

func (t *svc) Run(ctx context.Context) error {
	msgChan, errChan := t.consumer.Subscribe(ctx, amqp.QueueStep2)
	for {
		select {
		case msg := <-msgChan:
			if t.counter == 0 {
				t.startTime = time.Now()
			}
			if err := t.process(ctx, msg); err != nil {
				return errors.Wrap(err, "t.process")
			}
			if t.counter == t.config.XmlCount+t.config.JSONCount {
				return nil
			}
		case err := <-errChan:
			return err
		}
	}
}

func (t *svc) process(ctx context.Context, msg []byte) error {
	v := event.EventStep2{}
	if err := proto.Unmarshal(msg, &v); err != nil {
		return errors.Wrap(err, "proto.Unmarshal")
	}

	e := event.EventStep3{
		Event: &v,
		Meta4: helpers.GenerateString(30),
		Meta5: helpers.GenerateString(50),
	}

	t.counter++

	raw, err := proto.Marshal(&e)
	if err != nil {
		return errors.Wrap(err, "proto.Marshal")
	}

	return errors.Wrap(t.producer.Publish(ctx, amqp.QueueStep3, raw), "amqp.Publish")
}

func (t *svc) Close() error {
	text := fmt.Sprintf("total duration: %s;",
		time.Now().Sub(t.startTime).String(),
	)

	err := os.WriteFile(path.Join(t.config.ReportDir, "report3.txt"), []byte(text), 0666)
	if err != nil {
		return errors.Wrap(err, "os.WriteFile")
	}

	return nil
}
