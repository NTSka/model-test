package helpers

import (
	"crypto/sha256"
)

type CPUBound struct {
	rand RandGenerator
}

func NewCPUBound(rand RandGenerator) *CPUBound {
	return &CPUBound{
		rand: rand,
	}
}

func (t *CPUBound) CPUBoundTask() {
	st := ""
	for i := 0; i < 50; i++ {
		st += t.rand.GetString(10)
	}

	v := sha256.New()
	v.Write([]byte(st))
}
