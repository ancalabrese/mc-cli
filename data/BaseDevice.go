package data

type BaseDevice struct {
	Kind                   DeviceType             `json:"Kind"`
	CompliancePolicyStatus CompliancePolicyStatus `json:"CompliancePolicyStatusType"`
	ComplianceStatus       bool                   `json:"ComplianceStatus"`
	ComplianceItems        []ComplianceItem       `json:"ComplianceItems"`
	DeviceId               string                 `json:"DeviceId"`
	DeviceName             string                 `json:"DeviceName"`
	EnrollmentType         DeviceEnrollmentType   `json:"EnrollmentType"`
	EnrollmentTime         string                 `json:"EnrollmentTime"`
	Family                 DeviceFamilyType       `json:"Family"`
	HostName               string                 `json:"HostName"`
	IsAgentOnline          bool                   `json:"IsAgentOnline"`
	MacAddress             string                 `json:"MacAddress"`
	BluetoothMAC           string                 `json:"BluetoothMAC"`
	WifiMAC                string                 `json:"WifiMAC"`
	Mode                   DeviceMode             `json:"Mode"`
	Model                  string                 `json:"Model"`
	OsVersion              string                 `json:"OsVersion"`
	Path                   string                 `json:"Path"`
	ServerName             string                 `json:"ServerName"`
	Platform               PlatformType           `json:"Platform"`
	Manufacturer           string                 `json:"Manufacturer"`
}
