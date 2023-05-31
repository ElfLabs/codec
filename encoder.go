package codec

import (
	"fmt"
	"sync"
)

const (
	ErrNotFoundEncoder = codecError("not found Encoder")
)

type Encoder interface {
	Encode(v any) ([]byte, error)
}

type EncoderRegistry struct {
	encoders map[string]Encoder
	lock     sync.RWMutex
}

func NewEncoderRegistry() *EncoderRegistry {
	return &EncoderRegistry{
		encoders: make(map[string]Encoder),
	}
}

func (e *EncoderRegistry) RegisterEncoder(format string, enc Encoder) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.encoders[format] = enc
	return nil
}

func (e *EncoderRegistry) Encode(format string, v any) ([]byte, error) {
	e.lock.RLock()
	encoder, ok := e.encoders[format]
	e.lock.RUnlock()

	if !ok {
		return nil, fmt.Errorf("%w format: %s", ErrNotFoundEncoder, format)
	}

	return encoder.Encode(v)
}
