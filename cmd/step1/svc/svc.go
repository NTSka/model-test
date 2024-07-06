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
	processor processors.Processor
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
		config:    config,
		consumer:  consumer,
		producer:  producer,
		pool:      helpers.NewWorkersPull(config.Workers),
		finished:  make(chan struct{}),
		stop:      make(chan struct{}),
		ch:        make(chan []byte),
		Mutex:     &sync.Mutex{},
		processor: processors.NewProcessor(),
	}
}

var _ Service = (*svc)(nil)

func (t *svc) Run(ctx context.Context) error {
	msgChan, errChan := t.consumer.Subscribe(ctx, amqp.QueueEnter)

	err := t.producer.Publish(ctx, amqp.QueueStep1, t.ch, t.stop)
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
	s1, err := t.processor.Step1(msg.Data())
	if err != nil {
		msg.Nack()
		return errors.Wrap(err, "processors.Step1")
	}

	raw, err := proto.Marshal(s1)
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

	err := os.WriteFile(path.Join(t.config.ReportDir, "report1.txt"), []byte(text), 0666)
	if err != nil {
		return errors.Wrap(err, "os.WriteFile")
	}

	return nil
}
