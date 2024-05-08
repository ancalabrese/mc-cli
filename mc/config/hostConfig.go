package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ancalabrese/mc-cli/codec"
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

type Host struct {
	Configuration  `yaml:"-"`
	HostName       string         `yaml:"host,omitempty"`
	ClientId       string         `yaml:"clientId,omitempty"`
	ClientSecret   string         `yaml:"-"`
	CallbackURL    string         `yaml:"callbackURL,omitempty"`
	keyringService string         `yaml:"-"`
	stdinScanner   *bufio.Scanner `yaml:"-"`
	c              *Config        `yaml:"-"`
}

func NewHost(c *Config) *Host {
	return &Host{
		keyringService: "mc-cli",
		c:              c,
	}
}

func (ac *Host) Write() error {
	err := keyring.Set(ac.keyringService, ac.ClientId, ac.ClientSecret)

	os.MkdirAll(filepath.Dir(ac.c.Location), os.FileMode(0755))
	fp, err := os.OpenFile(ac.c.Location, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	utils.Check(err)

	c := codec.Codec{}
	err = c.Encode(fp, ac, codec.YAML)
	if err != nil {
		return fmt.Errorf("Host config write failed: %w", err)
	}
	return nil
}

func (ac *Host) Load() error {
	f, err := os.OpenFile(ac.c.Location, os.O_RDONLY, 0644)
	utils.Check(err)

	c := codec.Codec{}
	err = c.Decode(f, ac, codec.YAML)
	if err != nil {
		return fmt.Errorf("Couldn't retrieve host config: %w", err)
	}

	secret, err := keyring.Get(ac.keyringService, ac.ClientId)
	if err != nil {
		return fmt.Errorf("Couldn't retrieve secret: %w", err)
	}
	ac.ClientSecret = secret
	return nil
}
