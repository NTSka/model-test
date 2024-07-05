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

func Step1(msg []byte) (*event.EventStep1, error) {
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

	helpers.CPUBoundTask()

	return &event.EventStep1{
		Event:     e,
		Timestamp: t,
		Meta1:     helpers.GenerateString(50),
	}, nil
}

func Step2(v *event.EventStep1) *event.EventStep2 {
	e := event.EventStep2{
		Event: v,
		Meta2: helpers.GenerateString(30),
		Meta3: helpers.GenerateString(50),
	}

	helpers.CPUBoundTask()

	return &e
}

func Step3(v *event.EventStep2) *event.EventStep3 {
	e := event.EventStep3{
		Event: v,
		Meta4: helpers.GenerateString(30),
		Meta5: helpers.GenerateString(50),
	}

	helpers.CPUBoundTask()

	return &e
}
