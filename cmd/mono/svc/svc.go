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
	"test-model/pkg/storage"
	"time"
)

type Service interface {
	Run(ctx context.Context) error
	Close() error
}

type svc struct {
	config     *Config
	consumer   amqp.Consumer
	producer   amqp.Producer
	clickhouse storage.Clickhouse
	counter    int
	sumTime    int64
	startTime  time.Time
}

func NewSvc(config *Config, consumer amqp.Consumer, clickhouse storage.Clickhouse) Service {
	return &svc{
		config:     config,
		consumer:   consumer,
		clickhouse: clickhouse,
	}
}

func (t *svc) Close() error {
	text := fmt.Sprintf("total duration: %s; per event: %s;",
		time.Now().Sub(t.startTime).String(),
		time.Duration(t.sumTime/int64(t.counter)).String(),
	)

	err := os.WriteFile(path.Join(t.config.ReportDir, "mono.txt"), []byte(text), 0666)
	if err != nil {
		return errors.Wrap(err, "os.WriteFile")
	}

	return nil
}

func (t *svc) Run(ctx context.Context) error {
	msgChan, errChan := t.consumer.Subscribe(ctx, amqp.QueueEnter)
	rows := make([]*event.EventStep3, 0, t.config.XmlCount+t.config.JSONCount)

	for {
		select {
		case msg := <-msgChan:
			if t.counter == 0 {
				t.startTime = time.Now()
			}
			e1, err := t.processStep1(ctx, msg)
			if err != nil {
				return errors.Wrap(err, "t.process")
			}

			e2 := t.processStep2(ctx, e1)

			e3 := t.processStep3(ctx, e2)

			rows = append(rows, e3)

			if t.counter == t.config.XmlCount+t.config.JSONCount {
				if err = t.Insert(ctx, rows); err != nil {
					return errors.Wrap(err, "t.Insert")
				}
				return nil
			}
		case err := <-errChan:
			return err
		}
	}
}

func (t *svc) processStep1(ctx context.Context, msg []byte) (*event.EventStep1, error) {
	v := event.RawEvent{}
	if err := proto.Unmarshal(msg, &v); err != nil {
		return nil, errors.Wrap(err, "proto.Unmarshal")
	}

	e := &event.Event{}
	switch v.Format {
	case event.Format_XML:
		if err := xml.Unmarshal(v.Data, e); err != nil {
			return nil, errors.Wrap(err, "xml.Unmarshal")
		}
	case event.Format_JSON:
		if err := json.Unmarshal(v.Data, e); err != nil {
			return nil, errors.Wrap(err, "json.Unmarshal")
		}
	}

	t.counter++

	return &event.EventStep1{
		Event:     e,
		Timestamp: time.Now().Unix(),
		Meta1:     helpers.GenerateString(50),
	}, nil
}

func (t *svc) processStep2(ctx context.Context, v *event.EventStep1) *event.EventStep2 {
	e := event.EventStep2{
		Event: v,
		Meta2: helpers.GenerateString(30),
		Meta3: helpers.GenerateString(50),
	}

	return &e
}

func (t *svc) processStep3(ctx context.Context, v *event.EventStep2) *event.EventStep3 {
	e := event.EventStep3{
		Event: v,
		Meta4: helpers.GenerateString(30),
		Meta5: helpers.GenerateString(50),
	}

	return &e
}
