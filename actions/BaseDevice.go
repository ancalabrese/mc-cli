package actions

type BaseDevice struct {
	Kind                   DeviceType
	CompliancePolicyStatus CompliancePolicyStatusType
	ComplianceStatus       bool
	ComplianceItems        []ComplianceItem
	DeviceId               string
	DeviceName             string
	EnrollmentType         DeviceEnrollmentType
	EnrollmentTime         string
	Family                 DeviceFamilyType
	HostName               string
	IsAgentOnline          bool
	MacAddress             string
	BluetoothMAC           string
	WifiMAC                string
	Mode                   DeviceMode
	Model                  string
	OsVersion              string
	Path                   string
	ServerName             string
	Platform               PlatformType
	Manufacturer           string
}

func (bd *BaseDevice) Get() (*BaseDevice, error) {

}

type ComplianceItem struct {
	ComplianceType  ComplianceItemType
	ComplianceValue bool
}

type DeviceCustomAttribute struct {
	Name  string
	Value string
	Type  DataType
}

type DeviceMode struct {
	//Unknown PlatformType
	//  Active, Disabled, UnenrollPendingUser, UnenrollPendingAdmin, UnenrolledByUser, UnenrolledByAdmin
}

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
