package config

import (
	"bufio"
	"errors"
	"os"

	"github.com/ancalabrese/mc-cli/utils"
	"github.com/zalando/go-keyring"
)

const (
	McTokenKey       = "MC_TOKEN"
	McHostKey        = "MC_HOST"
	McClientIdKey    = "MC_CLIENT_ID"
	McSecretKey      = "MC_SECRET"
	McCallbackUriKey = "MC_CALLBACK_URI"
)

var AuthNotInitializedError = errors.New("AuthConfig not initialized")

type AuthConfig struct {
	host           string
	clientId       string
	clientSecret   string
	callbackURL    string
	keyringService string
	stdinScanner   *bufio.Scanner
}

func NewAuthConfig() *AuthConfig {
	return &AuthConfig{
		keyringService: "mcUtility",
	}

}

func (ac *AuthConfig) Write() error {
	if !ac.isInitialized() {
		return AuthNotInitializedError
	}

	err := os.Setenv(McHostKey, ac.host)
	err = os.Setenv(McClientIdKey, ac.clientId)
	err = os.Setenv(McSecretKey, ac.clientSecret)
	err = os.Setenv(McCallbackUriKey, ac.callbackURL)

	err = keyring.Set(ac.keyringService, McHostKey, ac.host)
	err = keyring.Set(ac.keyringService, McClientIdKey, ac.clientId)
	err = keyring.Set(ac.keyringService, McSecretKey, ac.clientSecret)
	err = keyring.Set(ac.keyringService, McCallbackUriKey, ac.callbackURL)

	return err
}

func (ac *AuthConfig) Get(k ConfigKey) (string, error) {
	value := os.Getenv(string(k))
	if value == "" {
		return keyring.Get(ac.keyringService, string(k))
	}
	return value, nil
}

func (ac *AuthConfig) Load() error {
	ac.host = ac.DetectOrPromptForHostname()
	ac.clientId = ac.DetectOrPromptForClientId()
	ac.clientSecret = ac.DetectOrPromptForClientSecret()
	ac.callbackURL = ac.DetectOrPromptForCallbackURL()

	return nil
}

func (ac *AuthConfig) DetectHost() (string, error) {
	host := os.Getenv(McHostKey)
	if host != "" {
		return host, nil
	}
	return keyring.Get(ac.keyringService, McHostKey)
}

func (ac *AuthConfig) DetectClientId() (string, error) {
	clientId := os.Getenv(McClientIdKey)
	if clientId != "" {
		return clientId, nil
	}

	return keyring.Get(ac.keyringService, McClientIdKey)
}

func (ac *AuthConfig) DetectClientSecret() (string, error) {
	secret := os.Getenv(McSecretKey)
	if secret != "" {
		return secret, nil
	}
	return keyring.Get(ac.keyringService, McSecretKey)
}

func (ac *AuthConfig) DetectCallbackURI() (string, error) {
	url := os.Getenv(McCallbackUriKey)
	if url != "" {
		return url, nil
	}
	return keyring.Get(ac.keyringService, McCallbackUriKey)
}

func (ac *AuthConfig) DetectOrPromptForHostname() string {
	host, err := ac.DetectHost()
	if err != nil {
		println("Mobicontrol server host name:")

		//TODO: sanitize
		host = ac.scanLine()
	}
	return host
}

func (ac *AuthConfig) DetectOrPromptForClientId() string {
	clientId, err := ac.DetectClientId()
	if err != nil {
		println("Mobicontrol client id:")

		//TODO: sanitize
		clientId = ac.scanLine()
	}
	return clientId
}

func (ac *AuthConfig) DetectOrPromptForClientSecret() string {
	secret, err := ac.DetectClientSecret()
	if err != nil {
		println("Mobicontrol client secret:")

		//TODO: sanitize
		secret = ac.scanLine()
	}
	return secret
}

func (ac *AuthConfig) DetectOrPromptForCallbackURL() string {
	url, err := ac.DetectClientSecret()
	if err != nil {
		println("Mobicontrol client callback url: (default mcauth://callback)")

		//TODO: sanitize
		url = ac.scanLine()
		if url == "" {
			url = "mcauth://callback"
		}
	}
	return url
}

func (ac *AuthConfig) isInitialized() bool {
	return ac.host != "" && ac.clientId != "" && ac.clientSecret != "" && ac.callbackURL != ""
}

func (ac *AuthConfig) scanLine() string {
	if ac.stdinScanner == nil {
		ac.stdinScanner = bufio.NewScanner(os.Stdin)
	}

	var line string
	scanner := ac.stdinScanner
	if scanner.Scan() {
		line = scanner.Text()
	}
	utils.Check(scanner.Err())

	return line
}
