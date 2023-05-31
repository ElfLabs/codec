package json

import (
	"encoding/json"
)

type Codec struct{}

func (Codec) Encode(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (Codec) Decode(b []byte, v any) error {
	return json.Unmarshal(b, v)
}
