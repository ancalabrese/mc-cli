package request

import (
	"fmt"
	"net/http"
)

type GetBaseDeviceReq struct {
	request  *http.Request
	deviceId string
}

type GetBaseDeviceRequestOpt func(*GetBaseDeviceReq)

func NewGetBaseDeviceRequest(deviceId string, options ...RequestOption) (*GetBaseDeviceReq, error) {
	httpRequest, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	req := &GetBaseDeviceReq{
		request:  httpRequest,
		deviceId: deviceId,
	}

	for _, opt := range options {
		err := opt(req.request)
		if err != nil {
			return req, fmt.Errorf("could not create request: %w", err)
		}
	}
	return req, nil
}
