package actions

type BaseDevice struct {
	Kind                   DeviceType                 `json:"Kind"`
	CompliancePolicyStatus CompliancePolicyStatusType `json:"CompliancePolicyStatusType"`
	ComplianceStatus       bool                       `json:"ComplianceStatus"`
	ComplianceItems        []ComplianceItem           `json:"ComplianceItems"`
	DeviceId               string                     `json:"DeviceId"`
	DeviceName             string                     `json:"DeviceName"`
	EnrollmentType         DeviceEnrollmentType       `json:"EnrollmentType"`
	EnrollmentTime         string                     `json:"EnrollmentTime"`
	Family                 DeviceFamilyType           `json:"Family"`
	HostName               string                     `json:"HostName"`
	IsAgentOnline          bool                       `json:"IsAgentOnline"`
	MacAddress             string                     `json:"MacAddress"`
	BluetoothMAC           string                     `json:"BluetoothMAC"`
	WifiMAC                string                     `json:"WifiMAC"`
	Mode                   DeviceMode                 `json:"Mode"`
	Model                  string                     `json:"Model"`
	OsVersion              string                     `json:"OsVersion"`
	Path                   string                     `json:"Path"`
	ServerName             string                     `json:"ServerName"`
	Platform               PlatformType               `json:"Platform"`
	Manufacturer           string                     `json:"Manufacturer"`
}

type ComplianceItem struct {
	ComplianceType  ComplianceItemType `json:"ComplianceType"`
	ComplianceValue bool               `json:"ComplianceValue"`
}

type DeviceCustomAttribute struct {
	Name  string   `json:"Name"`
	Value string   `json:"Value"`
	Type  DataType `json:"Type"`
}

type DeviceMode int

const (
	Unknown DeviceMode = iota
	Disabled
	UnenrollPendingUser
	UnenrollPendingAdmin
	UnenrolledByUser
	UnenrolledByAdmin
)

// PlatformType

type DeviceType struct {
	//TODO:ENUM
	// AndroidForWork, AndroidElm, AndroidKnox, AndroidPlus, AndroidGeneric, Ios, Mac, WindowsCE, WindowsDesktop, WindowsDesktopLegacy, WindowsPhone, WindowsRuntime, ZebraPrinter, Linux, WindowsHoloLens, WindowsXtHub, Unknown
}

type DataType struct {
	//TODO:ENUM
	// Boolean, Numeric, Text, DateTime, Enumerator, Date
}

type CompliancePolicyStatusType struct {
	//TODO:ENUM
	// Unknown, Pending, NonCompliant, Compliant
}

type ComplianceItemType struct {
	//TODO:ENUM
	// None, IsSecured, IsDeviceAdmin, NotWiped, IsEnabled, IsEnrolled
}

type DeviceEnrollmentType struct {
	//TODO:ENUM
	// NotApplicable, Device, User
}

type DeviceFamilyType struct {
	//TODO:ENUM
	// Unknown, WindowsCE, Apple, WindowsDesktop, AndroidPlus, Scanner, WindowsPhone, Blackberry, Printer, WindowsRuntime, Linux
}
