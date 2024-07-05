package amqp

import "github.com/pkg/errors"

var (
	ErrUnknownQueue = errors.New("unknown queue")
)
