package screen

import (
	"errors"
	"fmt"
	"io"
	"maps"
	"reflect"
)

const (
	PRINTER_TAG_NAME = "prettyPrint"
	IGNORE_TAG       = "-"
)

type Printer struct {
	stdout io.Writer
}

func NewPrinter(stdout io.Writer) *Printer {
	return &Printer{
		stdout: stdout,
	}
}

func (p *Printer) PrettyPrint(v interface{}) error {
	valueOfV := reflect.ValueOf(v)
	labelsToValues, err := p.getValues(valueOfV)
	if err != nil {
		return err
	}

	for i, k := range labelsToValues {
		fmt.Fprintf(p.stdout, "%s:\t%s\n", i, k)
	}
	return nil
}

func (p *Printer) getValues(v reflect.Value) (map[string]interface{}, error) {
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, errors.New("pretty print: cannot parse data, invalid type")
	}

	t := v.Type()
	labelsToValues := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)
		if fieldValue.Kind() == reflect.Struct {
			nested, err := p.getValues(v)
			if err != nil {
				// do something - add error to struct label
			}
			maps.Copy(labelsToValues, nested)
		}

		t := p.getTagValue(fieldType)
		labelsToValues[t] = fmt.Sprint(fieldValue.Interface())
	}
	return labelsToValues, nil
}

func (p *Printer) getTagValue(t reflect.StructField) string {
	value := t.Tag.Get(PRINTER_TAG_NAME)
	if value == IGNORE_TAG || value == "" {
		return ""
	}

	return value
}
