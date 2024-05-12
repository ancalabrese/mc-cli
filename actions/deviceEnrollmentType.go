package actions

type DeviceEnrollmentType int8

const (
	DeviceEnrollmentTypeNotApplicable DeviceEnrollmentType = iota
	DeviceEnrollmentTypeDevice
	DeviceEnrollmentTypeUser
	UnknownEnrollmentType
)

var (
	labelToDeviceEnrollmentTypeMap = map[string]DeviceEnrollmentType{
		"NotApplicable": DeviceEnrollmentTypeNotApplicable,
		"Device":        DeviceEnrollmentTypeDevice,
		"User":          DeviceEnrollmentTypeUser,
		"Unknown":       UnknownEnrollmentType,
	}
)

func (det *DeviceEnrollmentType) UnmarshalJSON(data []byte) error {
	return unmarshallCustomType(data, det, labelToDeviceEnrollmentTypeMap, UnknownEnrollmentType)
}
