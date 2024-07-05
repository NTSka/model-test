package helpers

import (
	"sync"
)

type WorkerFunc func()

type WorkersPull struct {
	ch    chan struct{}
	mutex *sync.Mutex
}

func (t *WorkersPull) Add(workerFunc WorkerFunc) {
	go func() {
		<-t.ch
		workerFunc()
		t.ch <- struct{}{}
	}()
}

func NewWorkersPull(count int) *WorkersPull {
	ch := make(chan struct{}, count)

	for i := 0; i < count; i++ {
		ch <- struct{}{}
	}
	return &WorkersPull{
		ch:    ch,
		mutex: &sync.Mutex{},
	}
}
