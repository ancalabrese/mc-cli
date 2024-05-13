package actions

type DeviceMode int8

const (
	DeviceModeUnknown DeviceMode = iota
	DeviceModeActive
	DeviceModeDisabled
	DeviceModeUnenrollPendingUser
	DeviceModeUnenrollPendingAdmin
	DeviceModeUnenrolledByUser
	DeviceModeUnenrolledByAdmin
)

var (
	labelToDeviceModeMap = map[string]DeviceMode{
		"Unknown":              DeviceModeUnknown,
		"Active":               DeviceModeActive,
		"UnenrollPendingAdmin": DeviceModeUnenrollPendingAdmin,
		"UnenrolledByUser":     DeviceModeUnenrolledByUser,
		"UnenrolledByAdmin":    DeviceModeUnenrolledByAdmin,
	}
)

func (dm *DeviceMode) UnmarshalJSON(data []byte) error {
	return unmarshallCustomType(data, dm, labelToDeviceModeMap, DeviceModeUnknown)
}
