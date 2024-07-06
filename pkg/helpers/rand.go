package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"math"
	"sync"
)

type generator64 struct {
	n uint64
}

type generator32 struct {
	n uint
}

type generatorString struct {
	s string
}

func makeSeedString() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func makeSeed64() uint64 {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return binary.BigEndian.Uint64(b)
}

func makeSeed32() uint {
	b := make([]byte, 4)
	_, _ = rand.Read(b)
	return uint(binary.BigEndian.Uint32(b))
}

//go:generate minimock -s _mock.go -o ../../mocks -i RandGenerator
type RandGenerator interface {
	GetUint64() uint64
	GetInt64() int64
	GetInt() int
	GetIntN(n int) int
	GetUint() uint
	GetFloatRand1() float64
	GetFloatRange(a, b float64) float64
	GetString(length int) string
}

type randGenerator struct {
	p64    *sync.Pool
	p32    *sync.Pool
	string *sync.Pool
}

func NewRandGenerator() RandGenerator {
	return &randGenerator{
		p64: &sync.Pool{
			New: func() interface{} {
				return &generator64{
					n: makeSeed64(),
				}
			},
		},
		p32: &sync.Pool{
			New: func() interface{} {
				return &generator32{n: makeSeed32()}
			},
		},
		string: &sync.Pool{
			New: func() interface{} {
				return &generatorString{
					s: makeSeedString(),
				}
			},
		},
	}
}

func (r *randGenerator) GetUint64() uint64 {
	g := r.p64.Get().(*generator64)

	g.n ^= g.n << 13
	g.n ^= g.n >> 7
	g.n ^= g.n << 17

	res := g.n

	r.p64.Put(g)

	return res
}

func (r *randGenerator) GetInt64() int64 {
	return int64(r.GetUint64())
}

func (r *randGenerator) GetInt() int {
	return int(r.GetUint())
}

func (r *randGenerator) GetUint() uint {
	g := r.p32.Get().(*generator32)

	g.n ^= g.n << 13
	g.n ^= g.n >> 17
	g.n ^= g.n << 5

	res := g.n

	r.p32.Put(g)

	return res
}

const maxUint64Float = float64(math.MaxUint64)

func (r *randGenerator) GetFloatRand1() float64 {
	v := float64(r.GetUint64())
	return v / maxUint64Float
}

func (r *randGenerator) GetFloatRange(a, b float64) float64 {
	max := math.Max(a, b)
	min := math.Min(a, b)
	diff := max - min
	return r.GetFloatRand1()*diff + min
}

func (r *randGenerator) GetIntN(n int) int {
	g := r.p64.Get().(*generator64)
	v := float64(g.n) / float64(math.MaxUint64)
	return int(math.Round(v * float64(n-1)))
}

func (r *randGenerator) GetString(length int) string {
	current := 0
	res := make([]byte, 0)
	for current < length {
		v := r.string.Get().(*generatorString).s
		diff := length - current
		if diff < len(v) {
			res = append(res, v[:diff]...)
		} else {
			res = append(res, v...)
		}

		current += len(v)
	}

	return string(res)
}
