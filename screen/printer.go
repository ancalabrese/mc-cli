package screen

import (
	"errors"
	"fmt"
	"io"
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
	// TODO: here check for slices, maps etc
	if valueOfV.Kind() == reflect.Pointer {
		valueOfV = valueOfV.Elem()
	}

	if valueOfV.Kind() != reflect.Struct {
		return errors.New("pretty print: cannot parse data, invalid type")
	}

	labelsToValues := make(map[string]any)
	typeOf := valueOfV.Type()

	for i := 0; i < valueOfV.NumField(); i++ {
		fValue := valueOfV.Field(i)
		fType := typeOf.Field(i)
		t := fType.Tag.Get(PRINTER_TAG_NAME)
		if t == IGNORE_TAG || t == "" {
		}
		labelsToValues[t] = fValue.Interface()
	}

	for i, k := range labelsToValues {
		fmt.Fprintf(p.stdout, "%s:\t%s\n", i, k)
	}

	return nil
}
