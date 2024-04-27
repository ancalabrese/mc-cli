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

type DeviceCustomAttribute struct {
	Name  string   `json:"Name"`
	Value string   `json:"Value"`
	Type  DataType `json:"Type"`
}

// PlatformType

type DataType struct {
	//TODO:ENUM
	// Boolean, Numeric, Text, DateTime, Enumerator, Date
}

type DeviceFamilyType struct {
	//TODO:ENUM
	// Unknown, WindowsCE, Apple, WindowsDesktop, AndroidPlus, Scanner, WindowsPhone, Blackberry, Printer, WindowsRuntime, Linux
}
