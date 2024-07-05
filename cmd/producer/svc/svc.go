package svc

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"test-model/pkg/amqp"
	"test-model/pkg/helpers"
	"test-model/pkg/proto/event"
)

type Svc interface {
	Run(ctx context.Context) error
}

type svc struct {
	config   *Config
	producer amqp.Producer
}

func NewSvc(config *Config, producer amqp.Producer) Svc {
	return &svc{
		config:   config,
		producer: producer,
	}
}

var _ Svc = (*svc)(nil)

func (t *svc) Run(ctx context.Context) error {
	for i := 0; i < t.config.JSONCount; i++ {
		raw, err := json.Marshal(t.makeEvent())
		if err != nil {
			return errors.Wrap(err, "json.Marshal")
		}

		rawEvent, err := proto.Marshal(&event.RawEvent{
			Format: event.Format_JSON,
			Data:   raw,
		})
		if err != nil {
			return errors.Wrap(err, "json.Marshal")
		}

		if err = t.producer.Publish(ctx, amqp.QueueEnter, rawEvent); err != nil {
			return errors.Wrap(err, "amqp.Publish")
		}
	}

	for i := 0; i < t.config.XmlCount; i++ {
		e := t.makeEvent()
		raw, err := xml.Marshal(e)
		if err != nil {
			return errors.Wrap(err, "xml.Marshal")
		}

		rawEvent, err := proto.Marshal(&event.RawEvent{
			Format: event.Format_XML,
			Data:   raw,
		})
		if err != nil {
			return errors.Wrap(err, "xml.Marshal")
		}

		if err = t.producer.Publish(ctx, amqp.QueueEnter, rawEvent); err != nil {
			return errors.Wrap(err, "producer.Publish")
		}
	}

	return nil
}

func (t *svc) makeEvent() *event.Event {
	assets := make([]string, 0)
	for i := 0; i < rand.Intn(5); i++ {
		assets = append(assets, helpers.GenerateAsset())
	}

	return &event.Event{
		Assets: assets,
		EventSrc: &event.EventSrc{
			Host: helpers.GenerateString(20),
			Type: helpers.GenerateString(10),
		},
		Action:     helpers.GenerateString(10),
		Importance: event.Importance(rand.Intn(4)),
		Object:     helpers.GenerateString(10),
	}
}
