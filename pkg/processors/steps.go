package processors

import (
	"encoding/json"
	"encoding/xml"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"test-model/pkg/helpers"
	"test-model/pkg/proto/event"
	"time"
)

type Processor interface {
	Step1(msg []byte) (*event.EventStep1, error)
	Step2(e *event.EventStep1) *event.EventStep2
	Step3(e *event.EventStep2) *event.EventStep3
}

type processor struct {
	rand helpers.RandGenerator
	cpu  *helpers.CPUBound
}

func NewProcessor() Processor {
	rand := helpers.NewRandGenerator()
	cp := helpers.NewCPUBound(rand)
	return &processor{
		rand: rand,
		cpu:  cp,
	}
}

func (p *processor) Step1(msg []byte) (*event.EventStep1, error) {
	v := event.RawEvent{}
	t := time.Now().UnixMilli()
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

	p.cpu.CPUBoundTask()

	return &event.EventStep1{
		Event:     e,
		Timestamp: t,
		Meta1:     p.rand.GetString(50),
	}, nil
}

func (p *processor) Step2(v *event.EventStep1) *event.EventStep2 {
	e := event.EventStep2{
		Event: v,
		Meta2: p.rand.GetString(30),
		Meta3: p.rand.GetString(50),
	}

	p.cpu.CPUBoundTask()

	return &e
}

func (p *processor) Step3(v *event.EventStep2) *event.EventStep3 {
	e := event.EventStep3{
		Event: v,
		Meta4: p.rand.GetString(30),
		Meta5: p.rand.GetString(50),
	}

	p.cpu.CPUBoundTask()

	return &e
}
