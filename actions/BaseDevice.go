package actions

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ancalabrese/mc-cli/mc/client"
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

const endpointPath = "devices"

type GetDevicesRequestOptions func() url.Values

func Take(v int) GetDevicesRequestOptions {
	return func() url.Values {
		return url.Values{
			"take": []string{strconv.Itoa(v)},
		}
	}
}

func Skip(v int) GetDevicesRequestOptions {
	return func() url.Values {
		return url.Values{
			"skip": []string{strconv.Itoa(v)},
		}
	}
}

func Path(p string) GetDevicesRequestOptions {
	return func() url.Values {
		return url.Values{
			"path": []string{url.QueryEscape(url.QueryEscape(p))},
		}
	}
}

func GetDevices(ctx context.Context,
	client client.McClient,
	opts ...GetDevicesRequestOptions) ([]*BaseDevice, error) {

	queryParams := url.Values{}

	for _, opt := range opts {
		for k, values := range opt() {
			queryParams[k] = append(queryParams[k], values...)
		}
	}

	endpoint := *client.DevicesEndpoint

	endpoint.RawQuery = queryParams.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
	utils.Check(err)

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	b := make([]byte, resp.ContentLength)
	_, err = resp.Body.Read(b)
	if err != nil {
		return nil, err
	}

	devices := make([]*BaseDevice, 0)
	json.Unmarshal(b, &devices)

	return devices, nil
}

func GetDeviceById(ctx context.Context, client client.McClient, deviceId string) (*BaseDevice, error) {
	devicesEndpoint := *client.DevicesEndpoint
	endpoint := devicesEndpoint.JoinPath(deviceId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
	utils.Check(err)

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	b := make([]byte, resp.ContentLength)
	_, err = resp.Body.Read(b)
	if err != nil {
		return nil, err
	}

	device := &BaseDevice{}
	json.Unmarshal(b, &device)

	return device, nil
}

func DeleteDevice(ctx context.Context, client client.McClient, deviceId string) error {
	devicesEndpoint := *client.DevicesEndpoint
	endpoint := devicesEndpoint.JoinPath(deviceId)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpoint.String(), nil)
	utils.Check(err)

	_, err = client.HttpClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
