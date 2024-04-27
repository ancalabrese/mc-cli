package actions

type DeviceMode int8

const (
	DeviceModeUnknown DeviceMode = iota
	DeviceModeDisabled
	DeviceModeUnenrollPendingUser
	DeviceModeUnenrollPendingAdmin
	DeviceModeUnenrolledByUser
	DeviceModeUnenrolledByAdmin
)

var (
	labelToDeviceModeMap = map[string]DeviceMode{
		"Unknown":              DeviceModeUnknown,
		"Disabled":             DeviceModeDisabled,
		"UnenrollPendingUser":  DeviceModeUnenrollPendingUser,
		"UnenrollPendingAdmin": DeviceModeUnenrollPendingAdmin,
		"UnenrolledByUser":     DeviceModeUnenrolledByUser,
		"UnenrolledByAdmin":    DeviceModeUnenrolledByAdmin,
	}
)

func (dm *DeviceMode) UnmarshalJSON(data []byte) error {
	return unmarshallCustomType(data, dm, labelToDeviceEnrollmentTypeMap)
}
