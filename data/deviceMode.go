package data

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
	deviceModeToLabelMap = map[DeviceMode]string{
		DeviceModeUnknown:              "Unknown",
		DeviceModeActive:               "Active",
		DeviceModeUnenrollPendingAdmin: "UnenrollPendingAdmin",
		DeviceModeUnenrolledByUser:     "UnenrolledByUser",
		DeviceModeUnenrolledByAdmin:    "UnenrolledByAdmin",
	}
)

func (dm *DeviceMode) UnmarshalJSON(data []byte) error {
	return unmarshallCustomType(data, dm, labelToDeviceModeMap, DeviceModeUnknown)
}

func (dm DeviceMode) String() string {
	return deviceModeToLabelMap[dm]
}
