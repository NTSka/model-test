package svc

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
	"path"
	"sync"
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
	clickhouse storage.Clickhouse
	counter    int
	sumTime    int64
	startTime  time.Time
	insertTime time.Duration
	rows       []*event.EventStep3
	msgs       []amqp.Message
	pool       *helpers.WorkersPull
	finished   chan struct{}
	*sync.Mutex
}

func NewSvc(config *Config, consumer amqp.Consumer, clickhouse storage.Clickhouse) Service {
	return &svc{
		config:     config,
		consumer:   consumer,
		clickhouse: clickhouse,
		rows:       make([]*event.EventStep3, 0, config.Total),
		msgs:       make([]amqp.Message, 0, config.Total),
		Mutex:      &sync.Mutex{},
		pool:       helpers.NewWorkersPull(config.Workers),
		finished:   make(chan struct{}),
	}
}

var _ Service = (*svc)(nil)

func (t *svc) Run(ctx context.Context) error {
	msgChan, errChan := t.consumer.Subscribe(ctx, amqp.QueueStep3)

	for {
		select {
		case msg := <-msgChan:
			if t.counter == 0 {
				t.startTime = time.Now()
			}
			t.pool.Add(func() {
				if err := t.process(ctx, msg); err != nil {
					log.Println("process err:", err)
				}
			})
		case <-t.finished:
			return nil
		case err := <-errChan:
			return err
		}
	}
}

func (t *svc) Close() error {
	text := fmt.Sprintf("total duration: %s; per event: %s; insert: %s",
		time.Now().Sub(t.startTime).String(),
		(time.Millisecond * time.Duration(t.sumTime/int64(t.counter))).String(),
		t.insertTime.String(),
	)

	err := os.WriteFile(path.Join(t.config.ReportDir, "report4.txt"), []byte(text), 0666)
	if err != nil {
		return errors.Wrap(err, "os.WriteFile")
	}

	return nil
}

func (t *svc) process(ctx context.Context, msg amqp.Message) error {
	v := event.EventStep3{}
	if err := proto.Unmarshal(msg.Data(), &v); err != nil {
		msg.Nack()
		return errors.Wrap(err, "proto.Unmarshal")
	}

	t.Mutex.Lock()
	t.counter++
	t.rows = append(t.rows, &v)
	t.msgs = append(t.msgs, msg)
	t.Mutex.Unlock()

	t.sumTime += time.Now().UnixMilli() - v.Event.Event.Timestamp
	if t.counter == t.config.Total {
		d := time.Now()
		log.Println("Start insert rows")
		if err := t.Insert(ctx, t.rows); err != nil {
			t.NackAll()
			return errors.Wrap(err, "t.Insert")
		}

		t.insertTime = time.Now().Sub(d)
		fmt.Printf("Inserted: %s", t.insertTime.String())

		t.AckAll()
		t.finished <- struct{}{}
		return nil
	}

	return nil
}

func (t *svc) AckAll() {
	for _, row := range t.msgs {
		row.Ack()
	}
}

func (t *svc) NackAll() {
	for _, row := range t.msgs {
		row.Nack()
	}
}
