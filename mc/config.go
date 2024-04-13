package mc

import (
	"bufio"
	"os"

	"github.com/ancalabrese/mc-cli/utils"
)

const mcTokenKey = "MC_TOKEN"
const mcHostKey = "MC_HOST"
const mcClientIdKey = "MC_CLIENT_ID"
const mcSecretKey = "MC_SECRET"
const mcCallbackUriKey = "MC_CALLBACK_URI"

type Config struct {
	stdinScanner *bufio.Scanner
}

func (c *Config) DetectApiToken() string {
	return os.Getenv(mcTokenKey)
}

func (c *Config) DetectHost() string {
	return os.Getenv(mcHostKey)
}

func (c *Config) DetectClientId() string {
	return os.Getenv(mcClientIdKey)
}

func (c *Config) DetectClientSecret() string {
	return os.Getenv(mcSecretKey)
}

func (c *Config) DetectCallbackURI() string {
	return os.Getenv(mcCallbackUriKey)
}

func (c *Config) PromptForHostname() string {
	host := c.DetectHost()
	if host != "" {
		return host
	}

	//TODO: sanitize
	host = c.scanLine()

	return host
}

func (c *Config) PromptForClientId() string {
	clientId := c.DetectClientId()
	if clientId != "" {
		return clientId
	}
	//TODO: sanitize
	clientId = c.scanLine()

	return clientId
}

func (c *Config) PromptForSecret() string {
	secret := c.DetectClientSecret()
	if secret != "" {
		return secret
	}

	//TODO: sanitize
	secret = c.scanLine()

	return secret
}

func (c *Config) scanLine() string {
	if c.stdinScanner == nil {
		c.stdinScanner = bufio.NewScanner(os.Stdin)
	}

	var line string
	scanner := c.stdinScanner
	if scanner.Scan() {
		line = scanner.Text()
	}
	utils.Check(scanner.Err())

	return line
}
