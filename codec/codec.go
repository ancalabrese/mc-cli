package codec

import (
	"encoding/json"
	"io"

	"gopkg.in/yaml.v3"
)

type Codec struct{}
type CodecType int

const (
	JSON CodecType = iota
	YAML
)

// Encode writes the objcet v in the provided io.Writer in Json format
func (c Codec) Encode(w io.Writer, v any, codec CodecType) error {
	if codec == JSON {
		return c.encodeJson(w, v)
	} else {
		return c.encodeYaml(w, v)
	}
}

// Decode reads from the reader and decode the next JSON value storing it in v
func (c Codec) Decode(r io.Reader, v any, codec CodecType) error {
	if codec == JSON {
		return c.decodeJson(r, v)
	} else {
		return c.decodeYaml(r, v)
	}
}

func (c Codec) encodeJson(w io.Writer, v any) error {
	enc := json.NewEncoder(w)
	err := enc.Encode(v)

	return err
}

func (c Codec) encodeYaml(w io.Writer, v any) error {
	enc := yaml.NewEncoder(w)
	err := enc.Encode(v)
	return err
}

func (c Codec) decodeJson(r io.Reader, v any) error {
	enc := json.NewDecoder(r)
	err := enc.Decode(v)

	return err
}

func (c Codec) decodeYaml(r io.Reader, v any) error {
	enc := yaml.NewDecoder(r)
	err := enc.Decode(v)
	return err
}
