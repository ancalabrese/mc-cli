package data

type DeviceFamilyType int8

const (
	DeviceFamilyTypeUnknown DeviceFamilyType = iota
	DeviceFamilyTypeWindowsCE
	DeviceFamilyTypeApple
	DeviceFamilyTypeWindowsDesktop
	DeviceFamilyTypeAndroidPlus
	DeviceFamilyTypeScanner
	DeviceFamilyTypeWindowsPhone
	DeviceFamilyTypeBlackberry
	DeviceFamilyTypePrinter
	DeviceFamilyTypeWindowsRuntime
	DeviceFamilyTypeLinux
)

var (
	labelToDeviceFamilyTypeMap = map[string]DeviceFamilyType{
		"Unknown":        DeviceFamilyTypeUnknown,
		"WindowsCE":      DeviceFamilyTypeWindowsCE,
		"Apple":          DeviceFamilyTypeApple,
		"WindowsDesktop": DeviceFamilyTypeWindowsDesktop,
		"AndroidPlus":    DeviceFamilyTypeAndroidPlus,
		"Scanner":        DeviceFamilyTypeScanner,
		"WindowsPhone":   DeviceFamilyTypeWindowsPhone,
		"Blackberry":     DeviceFamilyTypeBlackberry,
		"Printer":        DeviceFamilyTypePrinter,
		"WindowsRuntime": DeviceFamilyTypeWindowsRuntime,
		"Linux":          DeviceFamilyTypeLinux,
	}

	deviceFamilyTypeToLabelMap = map[DeviceFamilyType]string{
		DeviceFamilyTypeUnknown:        "Unknown",
		DeviceFamilyTypeWindowsCE:      "WindowsCE",
		DeviceFamilyTypeApple:          "Apple",
		DeviceFamilyTypeWindowsDesktop: "WindowsDesktop",
		DeviceFamilyTypeAndroidPlus:    "AndroidPlus",
		DeviceFamilyTypeScanner:        "Scanner",
		DeviceFamilyTypeWindowsPhone:   "WindowsPhone",
		DeviceFamilyTypeBlackberry:     "Blackberry",
		DeviceFamilyTypePrinter:        "Printer",
		DeviceFamilyTypeWindowsRuntime: "WindowsRuntime",
		DeviceFamilyTypeLinux:          "Linux",
	}
)

func (dft *DeviceFamilyType) UnmarshalJSON(data []byte) error {
	return unmarshallCustomType(data, dft, labelToDeviceFamilyTypeMap, DeviceFamilyTypeUnknown)
}

func (dft DeviceFamilyType) String() string {
	return deviceFamilyTypeToLabelMap[dft]
}
