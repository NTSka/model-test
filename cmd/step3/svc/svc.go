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
	"test-model/pkg/processors"
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
	pool      *helpers.WorkersPull
	finished  chan struct{}
	stop      chan struct{}
	ch        chan []byte
	*sync.Mutex
}

func NewSvc(config *Config, consumer amqp.Consumer, producer amqp.Producer) Service {
	return &svc{
		config:   config,
		consumer: consumer,
		producer: producer,
		finished: make(chan struct{}),
		pool:     helpers.NewWorkersPull(config.Workers),
		stop:     make(chan struct{}),
		ch:       make(chan []byte),
		Mutex:    &sync.Mutex{},
	}
}

var _ Service = (*svc)(nil)

func (t *svc) Run(ctx context.Context) error {
	msgChan, errChan := t.consumer.Subscribe(ctx, amqp.QueueStep2)

	err := t.producer.Publish(ctx, amqp.QueueStep3, t.ch, t.stop)
	if err != nil {
		return err
	}

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

func (t *svc) process(ctx context.Context, msg amqp.Message) error {
	v := event.EventStep2{}
	if err := proto.Unmarshal(msg.Data(), &v); err != nil {
		msg.Nack()
		return errors.Wrap(err, "proto.Unmarshal")
	}

	e := processors.Step3(&v)

	raw, err := proto.Marshal(e)
	if err != nil {
		msg.Nack()
		return errors.Wrap(err, "proto.Marshal")
	}

	t.ch <- raw

	msg.Ack()

	t.Lock()
	t.counter++

	if t.counter == t.config.Total {
		t.stop <- struct{}{}
		t.finished <- struct{}{}
	}

	t.Unlock()

	return nil
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
