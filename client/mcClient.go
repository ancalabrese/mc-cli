package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ancalabrese/mc-cli/auth"
	"github.com/ancalabrese/mc-cli/config"
	"github.com/ancalabrese/mc-cli/storage"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/oauth2"
)

type McClient struct {
	Host            string
	DisplayVersion  string
	Version         int64
	HttpClient      *http.Client
	ApiBaseAddress  *url.URL
	DevicesEndpoint *url.URL
	l               hclog.Logger
}

const mcPath = "Mobicontrol"
const apiPath = "api"
const MC_V14 = 14
const DEVICES_API_PATH = "devices"
const DEVICES_SEARCH_PATH = "search"

func NewMcClient(ctx context.Context, c *config.Config, l hclog.Logger) (*McClient, error) {
	if !isValidHostName(c.Api.HostName) {
		return nil, fmt.Errorf("Couldn't create a valid client for MC: Invalid host name %s", c.Api.HostName)
	}

	baseUrl := ConstructMcApiAddress(c.Api.HostName)
	oauth2Config, err := auth.GetOauth2Config(c)
	if err != nil {
		return nil, err
	}

	logger := l.Named("McClient")
	ss := storage.NewApiSecretStore(l)

	refreshToken, err := ss.GetRefreshAccessToken(c.Api.ClientId)
	if err != nil {
		return nil, err
	}

	//TODO: oauth2 library refreshes the token only if we set expire. For now always use refresh
	// token to get a new one for every new execution
	/* accessToken, err := ss.GetAccessToken(c.Api.ClientId)
	if err != nil {
		l.Error("Failed to retrieve access token. Using refresh token instead", "err:", err)
	} */

	t := &oauth2.Token{
		RefreshToken: refreshToken,
		// AccessToken:  accessToken,
	}

	return &McClient{
		Host:            c.Api.HostName,
		HttpClient:      oauth2Config.Client(ctx, t),
		ApiBaseAddress:  baseUrl,
		DevicesEndpoint: getDevicesApiEndpoint(baseUrl),
		l:               logger,
	}, nil
}

func ConstructMcApiAddress(host string) *url.URL {
	u, _ := url.Parse(host)
	if u.Scheme == "" {
		u.Scheme = "https"
	}
	return u.JoinPath(mcPath, apiPath)
}

func getDevicesApiEndpoint(baseUrl *url.URL) *url.URL {
	return baseUrl.JoinPath(DEVICES_API_PATH)
}

func isValidHostName(h string) bool {
	_, err := url.Parse(h)
	return err == nil
}
