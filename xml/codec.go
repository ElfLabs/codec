package json

import (
	"encoding/xml"
)

type Codec struct{}

func (Codec) Encode(v any) ([]byte, error) {
	return xml.Marshal(v)
}

func (Codec) Decode(b []byte, v any) error {
	return xml.Unmarshal(b, v)
}
