package actions

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ancalabrese/mc-cli/utils"
)

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

const endpoint = "devices"

func GetDevices(ctx context.Context, client http.Client) ([]BaseDevice, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "", nil)
	utils.Check(err)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	b := make([]byte, resp.ContentLength)
	_, err = resp.Body.Read(b)
	if err != nil {
		return nil, err
	}
	devices := make([]BaseDevice, 0)
	json.Unmarshal(b, &devices)

	return devices, nil
}

func GetDeviceById(client http.Client) ([]BaseDevice, error) {
	return nil, nil
}
