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
	"time"
)

type Svc interface {
	Run(ctx context.Context) error
}

type svc struct {
	config   *Config
	producer amqp.Producer
	total    int
	stop     chan struct{}
	ch       chan []byte
	rand     helpers.RandGenerator
}

func NewSvc(config *Config, producer amqp.Producer) Svc {
	return &svc{
		config:   config,
		producer: producer,
		stop:     make(chan struct{}),
		ch:       make(chan []byte),
		rand:     helpers.NewRandGenerator(),
	}
}

var _ Svc = (*svc)(nil)

func (t *svc) Run(ctx context.Context) error {
	ticker := time.NewTicker(time.Second)

	if err := t.producer.Publish(ctx, amqp.QueueEnter, t.ch, t.stop); err != nil {
		return errors.Wrap(err, "amqp.Publish")
	}

	for {
		select {
		case <-ticker.C:
			for i := 0; i < t.config.EPS; i++ {
				var rawEvent []byte
				format := rand.Intn(2)
				switch format {
				case 0:
					raw, err := json.Marshal(t.makeEvent())
					if err != nil {
						return errors.Wrap(err, "json.Marshal")
					}

					rawEvent, err = proto.Marshal(&event.RawEvent{
						Format: event.Format_JSON,
						Data:   raw,
					})
					if err != nil {
						return errors.Wrap(err, "json.Marshal")
					}
				case 1:
					raw, err := xml.Marshal(t.makeEvent())
					if err != nil {
						return errors.Wrap(err, "json.Marshal")
					}

					rawEvent, err = proto.Marshal(&event.RawEvent{
						Format: event.Format_XML,
						Data:   raw,
					})
					if err != nil {
						return errors.Wrap(err, "xml.Marshal")
					}
				}

				t.ch <- rawEvent

				t.total++
				if t.total >= t.config.Total {
					ticker.Stop()
					t.stop <- struct{}{}
					return nil
				}
			}
		}
	}
}

func (t *svc) makeEvent() *event.Event {
	assets := make([]string, 0)
	for i := 0; i < rand.Intn(5); i++ {
		assets = append(assets, helpers.GenerateAsset())
	}

	return &event.Event{
		Assets: assets,
		EventSrc: &event.EventSrc{
			Host: t.rand.GetString(20),
			Type: t.rand.GetString(10),
		},
		Action:     t.rand.GetString(10),
		Importance: event.Importance(rand.Intn(4)),
		Object:     t.rand.GetString(10),
	}
}
