package client

import (
	"fmt"
	"net/http"
	"net/url"
)

type McClient struct {
	Host           string
	DisplayVersion string
	Version        int64
	HttpClient     http.Client
}

const mcPath = "Mobicontrol"
const apiPath = "api"

func NewMcClient(host string) (*McClient, error) {
	if !isValidHostName(host) {
		return nil, fmt.Errorf("Couldn't create a valid client for MC: Invalid host name %s", host)
	}

	return &McClient{
		Host: host,
	}, nil
}

func isValidHostName(h string) bool {
	u, err := url.Parse("https://" + h)
	if err != nil {
		return false
	}

	return u.Hostname() == h
}
