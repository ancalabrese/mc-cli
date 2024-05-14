package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ancalabrese/mc-cli/data"
)

type GetDevicesRequestOptions func() url.Values

func (mcc *McClient) Take(v int) GetDevicesRequestOptions {
	return func() url.Values {
		return url.Values{
			"take": []string{strconv.Itoa(v)},
		}
	}
}

func (mcc *McClient) Skip(v int) GetDevicesRequestOptions {
	return func() url.Values {
		return url.Values{
			"skip": []string{strconv.Itoa(v)},
		}
	}
}

func (mcc *McClient) Path(p string) GetDevicesRequestOptions {
	return func() url.Values {
		if p == "" {
			return url.Values{}
		}

		return url.Values{
			"path": []string{url.QueryEscape(url.QueryEscape(p))},
		}
	}
}

func (mcc *McClient) GetDevices(ctx context.Context,
	opts ...GetDevicesRequestOptions) ([]*data.BaseDevice, error) {
	queryParams := url.Values{}

	for _, opt := range opts {
		for k, values := range opt() {
			queryParams[k] = append(queryParams[k], values...)
		}
	}

	endpoint := *mcc.DevicesEndpoint
	endpoint.RawQuery = queryParams.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create device request: %w", err)
	}

	mcc.l.Debug("GetDevices", "executing", req.URL.String())
	resp, err := mcc.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP error while requesting devices: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// TODO: handle different type of responses
		mcc.l.Error("GetDevices", "server returned", resp.StatusCode)
		return nil, fmt.Errorf("Server returned non OK code: %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	devices := make([]*data.BaseDevice, 0)
	err = json.Unmarshal(b, &devices)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (mcc *McClient) GetDeviceById(ctx context.Context, deviceId string) (*data.BaseDevice, error) {
	devicesEndpoint := *mcc.DevicesEndpoint
	endpoint := devicesEndpoint.JoinPath(deviceId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create device request: %w", err)
	}

	mcc.l.Debug("GetDeviceById", "executing", req.URL.String())
	resp, err := mcc.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP error while requesting device info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// TODO: handle different type of responses
		mcc.l.Error("GetDeviceById", "server returned", resp.StatusCode)
		return nil, fmt.Errorf("Server returned non OK code: %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	device := &data.BaseDevice{}
	err = json.Unmarshal(b, &device)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (mcc *McClient) DeleteDevice(ctx context.Context, deviceId string) error {
	devicesEndpoint := *mcc.DevicesEndpoint
	endpoint := devicesEndpoint.JoinPath(deviceId)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpoint.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create device request: %w", err)
	}

	mcc.l.Debug("DeleteDevice", "executing", req.URL.String())
	resp, err := mcc.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		// TODO: handle different type of responses
		mcc.l.Error("DeleteDevice", "server returned", resp.StatusCode)
		return fmt.Errorf("Server returned non OK code: %d", resp.StatusCode)
	}

	return nil
}
