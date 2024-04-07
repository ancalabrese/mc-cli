package codec

import (
	"encoding/json"
	"fmt"
	"io"
)

type Codec struct{}

// Encode writes the objcet v in the provided io.Writer in Json format
func (Codec) Encode(w io.Writer, v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("Json encoding error: %w", err)
	}
	_, err = w.Write(b)
	return err
}

// Decode reads from the reader and decode the next JSON value storing it in v
func (Codec) Decode(r io.Reader, v any) error {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		err = fmt.Errorf("Json decoding error: %w", err)
	}
	return err
}
