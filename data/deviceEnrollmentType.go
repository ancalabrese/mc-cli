package data

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

	deviceEnrollmentTypeToLabelMap = map[DeviceEnrollmentType]string{
		DeviceEnrollmentTypeNotApplicable: "NotApplicable",
		DeviceEnrollmentTypeDevice:        "Device",
		DeviceEnrollmentTypeUser:          "User",
		UnknownEnrollmentType:             "Unknown",
	}
)

func (det *DeviceEnrollmentType) UnmarshalJSON(data []byte) error {
	return unmarshallCustomType(data, det, labelToDeviceEnrollmentTypeMap, UnknownEnrollmentType)
}

func (det DeviceEnrollmentType) String() string {
	return deviceEnrollmentTypeToLabelMap[det]
}
