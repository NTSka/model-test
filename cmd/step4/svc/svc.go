package svc

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"os"
	"path"
	"test-model/pkg/amqp"
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

var _ Service = (*svc)(nil)

func (t *svc) Run(ctx context.Context) error {
	msgChan, errChan := t.consumer.Subscribe(ctx, amqp.QueueStep3)
	rows := make([]*event.EventStep3, 0, t.config.XmlCount+t.config.JSONCount)

	for {
		select {
		case msg := <-msgChan:
			if t.counter == 0 {
				t.startTime = time.Now()
			}
			v, err := t.process(ctx, msg)
			if err != nil {
				return errors.Wrap(err, "t.process")
			}
			rows = append(rows, v)
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

func (t *svc) Close() error {
	text := fmt.Sprintf("total duration: %s; per event: %s;",
		time.Now().Sub(t.startTime).String(),
		time.Duration(t.sumTime/int64(t.counter)).String(),
	)

	err := os.WriteFile(path.Join(t.config.ReportDir, "report4.txt"), []byte(text), 0666)
	if err != nil {
		return errors.Wrap(err, "os.WriteFile")
	}

	return nil
}

func (t *svc) process(ctx context.Context, msg []byte) (*event.EventStep3, error) {
	v := event.EventStep3{}
	if err := proto.Unmarshal(msg, &v); err != nil {
		return nil, errors.Wrap(err, "proto.Unmarshal")
	}

	t.counter++

	t.sumTime += time.Now().Unix() - v.Event.Event.Timestamp

	return &v, nil
}
