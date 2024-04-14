package mc

import (
	"bufio"
	"errors"
	"os"

	"github.com/ancalabrese/mc-cli/utils"
	"github.com/zalando/go-keyring"
)

const (
	mcTokenKey       = "MC_TOKEN"
	mcHostKey        = "MC_HOST"
	mcClientIdKey    = "MC_CLIENT_ID"
	mcSecretKey      = "MC_SECRET"
	mcCallbackUriKey = "MC_CALLBACK_URI"
)

type ConfigKey string
type ConfigValue string
type DefaultValue string

type Configuration interface {
	Write() error
	Get(ConfigKey, DefaultValue) error
	Load() error
}

// Config reppresents a persistent configuration
type Config struct {
	Authentication Configuration
}

func NewConfig() *Config {
	c := &Config{
		Authentication: &AuthConfig{},
	}

	return c
}

var AuthNotInitializedError = errors.New("AuthConfig not initialized")

type AuthConfig struct {
	c *Config

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

	err := os.Setenv(mcHostKey, ac.host)
	err = os.Setenv(mcClientIdKey, ac.clientId)
	err = os.Setenv(mcSecretKey, ac.clientSecret)
	err = os.Setenv(mcCallbackUriKey, ac.callbackURL)

	err = keyring.Set(ac.keyringService, mcHostKey, ac.host)
	err = keyring.Set(ac.keyringService, mcClientIdKey, ac.clientId)
	err = keyring.Set(ac.keyringService, mcSecretKey, ac.clientSecret)
	err = keyring.Set(ac.keyringService, mcCallbackUriKey, ac.callbackURL)

	return err
}

func (ac *AuthConfig) Get(k ConfigKey, v DefaultValue) error {
	return nil
}

func (ac *AuthConfig) Load() error {
	ac.host = ac.DetectOrPromptForHostname()
	ac.clientId = ac.DetectOrPromptForClientId()
	ac.clientSecret = ac.DetectOrPromptForClientSecret()
	ac.callbackURL = ac.DetectOrPromptForCallbackURL()

	return nil
}

func (ac *AuthConfig) DetectHost() (string, error) {
	host := os.Getenv(mcHostKey)
	if host != "" {
		return host, nil
	}
	return keyring.Get(ac.keyringService, mcHostKey)
}

func (ac *AuthConfig) DetectClientId() (string, error) {
	clientId := os.Getenv(mcClientIdKey)
	if clientId != "" {
		return clientId, nil
	}

	return keyring.Get(ac.keyringService, mcClientIdKey)
}

func (ac *AuthConfig) DetectClientSecret() (string, error) {
	secret := os.Getenv(mcSecretKey)
	if secret != "" {
		return secret, nil
	}
	return keyring.Get(ac.keyringService, mcSecretKey)
}

func (ac *AuthConfig) DetectCallbackURI() (string, error) {
	url := os.Getenv(mcCallbackUriKey)
	if url != "" {
		return url, nil
	}
	return keyring.Get(ac.keyringService, mcCallbackUriKey)
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
