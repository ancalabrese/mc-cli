package auth

import (
	"encoding/base64"

	"github.com/ancalabrese/mc-cli/mc"
	"golang.org/x/oauth2"
)

type AuthClient struct {
	oauth2.Config
}

func NewAuthClient(c *mc.Config) *AuthClient {
	config := &AuthClient{}
	endpoint := oauth2.Endpoint{}

	endpoint := oauth2.Endpoint{}
	config.Endpoint.DeviceAuthURL = c.DetectHost()
	config.ClientID = c.DetectClientId()
	config.ClientSecret = c.DetectClientSecret()
	config.RedirectURL = c.DetectCallbackURI()
	return config
}

func getAuthorizationCodeRequestHeader(clientId string, secret string) string {
	return base64.RawStdEncoding.EncodeToString([]byte(clientId + ":" + secret))
}
