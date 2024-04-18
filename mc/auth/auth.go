package auth

import (
	"fmt"

	"github.com/ancalabrese/mc-cli/mc/config"
	"golang.org/x/oauth2"
)

type AuthClient struct {
	endpoint oauth2.Endpoint
}

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

	oauthConfig := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  callbackUrl,
		Endpoint: oauth2.Endpoint{
			//TODO: use real addresses
			AuthURL:   addr,
			TokenURL:  addr,
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
	fmt.Printf("Login at:\n%s\n", oauthConfig.AuthCodeURL(oauth2.GenerateVerifier(), oauth2.AccessTypeOffline))
	return nil
}

func getAccessToken() {}
