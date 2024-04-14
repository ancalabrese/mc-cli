package auth

import (
	"encoding/base64"
	"net/url"

	"github.com/ancalabrese/mc-cli/mc"
	"github.com/ancalabrese/mc-cli/utils"
	"golang.org/x/oauth2"
)

const TOKEN_URL_PATH = "token"

type AuthClient struct {
	oauth2.Config
}

func NewAuthClient(c *mc.Config) *AuthClient {
	config := &AuthClient{}

	endpoint := oauth2.Endpoint{
		AuthURL: createAuthURL(c.DetectHost()),
	}

	config.Endpoint = endpoint

	return config
}

func createAuthURL(hostName string) string {
	url, err := url.JoinPath(hostName, TOKEN_URL_PATH)
	utils.Check(err)

	return url
}

func createTokenURL() string {
	return ""
}

func getAuthorizationCodeRequestHeader(clientId string, secret string) string {
	return base64.RawStdEncoding.EncodeToString([]byte(clientId + ":" + secret))
}
