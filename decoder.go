package codec

import (
	"fmt"
	"sync"
)

const (
	ErrNotFoundDecoder = codecError("not found Decoder")
)

type Decoder interface {
	Decode([]byte, any) error
}

type DecoderRegistry struct {
	decoders map[string]Decoder
	lock     sync.RWMutex
}

func NewDecoderRegistry() *DecoderRegistry {
	return &DecoderRegistry{
		decoders: make(map[string]Decoder),
	}
}

func (e *DecoderRegistry) RegisterDecoder(format string, enc Decoder) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.decoders[format] = enc
	return nil
}

func (e *DecoderRegistry) Decode(format string, b []byte, v any) error {
	e.lock.RLock()
	decoder, ok := e.decoders[format]
	e.lock.RUnlock()

	if !ok {
		return fmt.Errorf("%w format: %s", ErrNotFoundDecoder, format)
	}

	return decoder.Decode(b, v)
}
