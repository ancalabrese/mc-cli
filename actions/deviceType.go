package actions

type DeviceType int8

const (
	DeviceTypeAndroidForWork DeviceType = iota
	DeviceTypeAndroidElm
	DeviceTypeAndroidKnox
	DeviceTypeAndroidPlus
	DeviceTypeAndroidGeneric
	DeviceTypeIos
	DeviceTypeMac
	DeviceTypeWindowsCE
	DeviceTypeWindowsDesktop
	DeviceTypeWindowsDesktopLegacy
	DeviceTypeWindowsPhone
	DeviceTypeWindowsRuntime
	DeviceTypeZebraPrinter
	DeviceTypeLinux
	DeviceTypeWindowsHoloLens
	DeviceTypeWindowsXtHub
	DeviceTypeUnknown
)

var (
	labelToDeviceType = map[string]DeviceType{
		"AndroidForWork":       DeviceTypeAndroidForWork,
		"AndroidElm":           DeviceTypeAndroidElm,
		"AndroidKnox":          DeviceTypeAndroidKnox,
		"AndroidPlus":          DeviceTypeAndroidPlus,
		"AndroidGeneric":       DeviceTypeAndroidGeneric,
		"Ios":                  DeviceTypeIos,
		"Mac":                  DeviceTypeMac,
		"WindowsCE":            DeviceTypeWindowsCE,
		"WindowsDesktop":       DeviceTypeWindowsDesktop,
		"WindowsDesktopLegacy": DeviceTypeWindowsDesktopLegacy,
		"WindowsPhone":         DeviceTypeWindowsPhone,
		"WindowsRuntime":       DeviceTypeWindowsRuntime,
		"ZebraPrinter":         DeviceTypeZebraPrinter,
		"Linux":                DeviceTypeLinux,
		"WindowsHoloLens":      DeviceTypeWindowsHoloLens,
		"WindowsXtHub":         DeviceTypeWindowsXtHub,
		"Unknown":              DeviceTypeUnknown,
	}
)

func (dt *DeviceType) UnmarshalJSON(data []byte) error {
	return unmarshallCustomType(data, dt, labelToDeviceType)
}
