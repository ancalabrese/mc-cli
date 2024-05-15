package data

type BaseDevice struct {
	Kind                   DeviceType             `json:"Kind" prettyPrint:"Kind"`
	CompliancePolicyStatus CompliancePolicyStatus `json:"CompliancePolicyStatusType" prettyPrint:"Compliance policy status"`
	ComplianceStatus       bool                   `json:"ComplianceStatus" prettyPrint:"Compliance status"`
	ComplianceItems        []ComplianceItem       `json:"ComplianceItems" prettyPrint:"Compliance Items"`
	DeviceId               string                 `json:"DeviceId" prettyPrint:"Device ID"`
	DeviceName             string                 `json:"DeviceName" prettyPrint:"Device name"`
	EnrollmentType         DeviceEnrollmentType   `json:"EnrollmentType" prettyPrint:"Enrollment type"`
	EnrollmentTime         string                 `json:"EnrollmentTime" prettyPrint:"Enrollment time"`
	Family                 DeviceFamilyType       `json:"Family" prettyPrint:"Family"`
	HostName               string                 `json:"HostName" prettyPrint:"Host name"`
	IsAgentOnline          bool                   `json:"IsAgentOnline" prettyPrint:"Is online"`
	MacAddress             string                 `json:"MacAddress" prettyPrint:"MAC"`
	BluetoothMAC           string                 `json:"BluetoothMAC" prettyPrint:"Bluetooth MAC"`
	WifiMAC                string                 `json:"WifiMAC" prettyPrint:"Wifi MAC"`
	Mode                   DeviceMode             `json:"Mode" prettyPrint:"Mode"`
	Model                  string                 `json:"Model" prettyPrint:"Model"`
	OsVersion              string                 `json:"OsVersion" prettyPrint:"Os version"`
	Path                   string                 `json:"Path" prettyPrint:"Path"`
	ServerName             string                 `json:"ServerName" prettyPrint:"Server name"`
	Platform               PlatformType           `json:"Platform" prettyPrint:"Platform"`
	Manufacturer           string                 `json:"Manufacturer" prettyPrint:"Manufacturer"`
}
