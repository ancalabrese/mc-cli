package auth

import (
	"fmt"
	url "net/url"

	"github.com/ancalabrese/mc-cli/mc/config"
	"github.com/ancalabrese/mc-cli/utils"
	"golang.org/x/oauth2"
)

const (
	mcUrlAuthorizationPath = "oauth/authorize"
	mcTokenUrlPath         = "token"
	mcApiUrlPath           = "api"
)

type AuthorizationResponseType int

const (
	AuthResponseTypeCode AuthorizationResponseType = iota
	AuthResponseTypeToken
)

func RequestAuthCode(c *config.Config) error {
	addr, err := c.Authentication.Get(config.McHostKey)
	if err != nil {
		return fmt.Errorf("Failed to retrieve auth address: %w", err)
	}

	clientId, err := c.Authentication.Get(config.McClientIdKey)
	if err != nil {
		return fmt.Errorf("Failed to retrieve client id: %w", err)
	}

	clientSecret, err := c.Authentication.Get(config.McSecretKey)
	if err != nil {
		return fmt.Errorf("Failed to retrieve client secret : %w", err)
	}

	callbackUrl, err := c.Authentication.Get(config.McCallbackUriKey)
	if err != nil {
		return fmt.Errorf("Failed to retrieve auth callback url: %w", err)
	}

	mcAuthUrl, err := url.JoinPath(addr, mcUrlAuthorizationPath)
	mcTokenUrl, err := url.JoinPath(addr, mcApiUrlPath, mcTokenUrlPath)
	utils.Check(err)

	oauthConfig := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  callbackUrl,
		Endpoint: oauth2.Endpoint{
			AuthURL:   mcAuthUrl,
			TokenURL:  mcTokenUrl,
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}

	authCodeUrl := oauthConfig.AuthCodeURL(
		oauth2.GenerateVerifier(),
		oauth2.AccessTypeOffline,
		setAuthorizationType(AuthResponseTypeCode))

	// For some reasons redirect_url breaks Mobicontrol Authorization page
	parsedUrl, err := url.Parse(authCodeUrl)
	utils.Check(err)
	urlParams := parsedUrl.Query()
	urlParams.Del("redirect_uri")
	parsedUrl.RawQuery = urlParams.Encode()

	fmt.Printf("Login at:\n%s\n", parsedUrl.String())
	return nil
}

func getAccessToken() {}

func setAuthorizationType(responseType AuthorizationResponseType) oauth2.AuthCodeOption {
	if responseType == AuthResponseTypeCode {
		return oauth2.SetAuthURLParam("response_type", "code")
	} else {
		return oauth2.SetAuthURLParam("response_type", "token")
	}
}
