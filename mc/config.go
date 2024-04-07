package mc

import (
	"bufio"
	"os"

	"github.com/ancalabrese/mc-cli/utils"
)

const mcTokenKey = "MC_TOKEN"
const mcHostKey = "MC_HOST"
const mcClientId = "MC_CLIENT_ID"
const mcSecret = "MC_SECRET"

type Config struct {
	stdinScanner *bufio.Scanner
}

func (c *Config) DetectApiToken() string {
	return os.Getenv(mcTokenKey)
}

func (c *Config) DetectHost() string {
	return os.Getenv(mcHostKey)
}

func (c *Config) PromptForHostname() string {
	host := os.Getenv(mcHostKey)
	if host != "" {
		return host
	}

	host = c.scanLine()

	return ""
}

func (c *Config) PromptForClientId() string {
	host := os.Getenv(mcClientId)
	if host != "" {
		return host
	}

	host = c.scanLine()

	return ""
}

func (c *Config) PromptForSecret() string {
	host := os.Getenv(mcSecret)
	if host != "" {
		return host
	}

	host = c.scanLine()

	return ""
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
