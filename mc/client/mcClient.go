package client

import (
	"fmt"
	"net/http"
	"net/url"
)

type McClient struct {
	Host               string
	DisplayVersion     string
	Version            int64
	HttpClient         *http.Client
	ApiBaseAddress     string
	DevicesApiEndpoint func() string
}

const mcPath = "Mobicontrol"
const apiPath = "api"
const MC_V14 = 14
const DEVICES_API_PATH = "devices"
const DEVICES_SEARCH_PATH = "search"

func NewMcClient(host string, c *http.Client) (*McClient, error) {
	if !isValidHostName(host) {
		return nil, fmt.Errorf("Couldn't create a valid client for MC: Invalid host name %s", host)
	}

	return &McClient{
		Host:           host,
		HttpClient:     c,
		ApiBaseAddress: getApiAddress(host),
	}, nil
}

func (mcc *McClient) getDevicesApiEndpoint() string {
	u, _ := url.Parse(mcc.ApiBaseAddress)

	if mcc.Version == 0 || mcc.Version < MC_V14 {
		return u.JoinPath(DEVICES_API_PATH).String()
	}

	return u.JoinPath(DEVICES_SEARCH_PATH).String()
}

func getApiAddress(host string) string {
	u, _ := url.Parse("https://" + host)
	return u.JoinPath(mcPath, apiPath).String()
}

func isValidHostName(h string) bool {
	u, err := url.Parse("https://" + h)
	if err != nil {
		return false
	}

	return u.Hostname() == h
}
