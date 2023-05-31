package toml

import (
	"github.com/pelletier/go-toml/v2"
)

type Codec struct{}

func (Codec) Encode(v any) ([]byte, error) {
	return toml.Marshal(v)
}

func (Codec) Decode(b []byte, v any) error {
	return toml.Unmarshal(b, v)
}
