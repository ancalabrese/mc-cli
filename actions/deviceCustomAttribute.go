package actions

type DeviceCustomAttribute struct {
	Name  string   `json:"Name"`
	Value string   `json:"Value"`
	Type  DataType `json:"Type"`
}

type DataType int8

const (
	DataTypeBoolean DataType = iota
	DataTypeNumeric
	DataTypeText
	DataTypeDateTime
	DataTypeEnumerator
	DataTypeDate
)

var (
	labelToDataTypeMap = map[string]DataType{
		"Boolean":    DataTypeBoolean,
		"Numeric":    DataTypeNumeric,
		"Text":       DataTypeText,
		"ateTime":    DataTypeDateTime,
		"Enumerator": DataTypeEnumerator,
		"Date":       DataTypeDate,
	}
)

func (dca *DeviceCustomAttribute) UnmarshalJSON(data []byte) error {
	return unmarshallCustomType(data, dca, labelToDataTypeMap)
}
