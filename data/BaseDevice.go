package data

type BaseDevice struct {
	Kind                   DeviceType             `json:"Kind" prettyPrint:"Kind"`
	CompliancePolicyStatus CompliancePolicyStatus `json:"CompliancePolicyStatusType" prettyPrint:"CompliancePolicyStatus"`
	ComplianceStatus       bool                   `json:"ComplianceStatus" prettyPrint:"ComplianceStatus"`
	ComplianceItems        []ComplianceItem       `json:"ComplianceItems" prettyPrint:"ComplianceItems"`
	DeviceId               string                 `json:"DeviceId" prettyPrint:"DeviceId"`
	DeviceName             string                 `json:"DeviceName" prettyPrint:"DeviceName"`
	EnrollmentType         DeviceEnrollmentType   `json:"EnrollmentType" prettyPrint:"EnrollmentType"`
	EnrollmentTime         string                 `json:"EnrollmentTime" prettyPrint:"EnrollmentTime"`
	Family                 DeviceFamilyType       `json:"Family" prettyPrint:"Family"`
	HostName               string                 `json:"HostName" prettyPrint:"HostName"`
	IsAgentOnline          bool                   `json:"IsAgentOnline" prettyPrint:"IsAgentOnline"`
	MacAddress             string                 `json:"MacAddress" prettyPrint:"MacAddress"`
	BluetoothMAC           string                 `json:"BluetoothMAC" prettyPrint:"BluetoothMAC"`
	WifiMAC                string                 `json:"WifiMAC" prettyPrint:"WifiMAC"`
	Mode                   DeviceMode             `json:"Mode" prettyPrint:"Mode"`
	Model                  string                 `json:"Model" prettyPrint:"Model"`
	OsVersion              string                 `json:"OsVersion" prettyPrint:"OsVersion"`
	Path                   string                 `json:"Path" prettyPrint:"Path"`
	ServerName             string                 `json:"ServerName" prettyPrint:"ServerName"`
	Platform               PlatformType           `json:"Platform" prettyPrint:"Platform"`
	Manufacturer           string                 `json:"Manufacturer" prettyPrint:"Manufacturer"`
}
