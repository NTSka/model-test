package svc

import (
	"context"
	"encoding/json"
	"encoding/xml"
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
	msgChan, errChan := t.consumer.Subscribe(ctx, amqp.QueueEnter)
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
	v := event.RawEvent{}
	if err := proto.Unmarshal(msg, &v); err != nil {
		return errors.Wrap(err, "proto.Unmarshal")
	}

	e := &event.Event{}
	switch v.Format {
	case event.Format_XML:
		if err := xml.Unmarshal(v.Data, e); err != nil {
			return errors.Wrap(err, "xml.Unmarshal")
		}
	case event.Format_JSON:
		if err := json.Unmarshal(v.Data, e); err != nil {
			return errors.Wrap(err, "json.Unmarshal")
		}
	}

	t.counter++

	s1 := event.EventStep1{
		Event:     e,
		Timestamp: time.Now().Unix(),
		Meta1:     helpers.GenerateString(50),
	}

	raw, err := proto.Marshal(&s1)
	if err != nil {
		return errors.Wrap(err, "proto.Marshal")
	}

	return errors.Wrap(t.producer.Publish(ctx, amqp.QueueStep1, raw), "amqp.Publish")
}

func (t *svc) Close() error {
	text := fmt.Sprintf("total duration: %s;",
		time.Now().Sub(t.startTime).String(),
	)

	err := os.WriteFile(path.Join(t.config.ReportDir, "report1.txt"), []byte(text), 0666)
	if err != nil {
		return errors.Wrap(err, "os.WriteFile")
	}

	return nil
}
