package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ancalabrese/mc-cli/codec"
	"github.com/ancalabrese/mc-cli/utils"
	"github.com/hashicorp/go-hclog"
	"gopkg.in/yaml.v3"
)

type Configuration interface {
	Write() error
	Load() error
}

const CONFIG_FILE_NAME = "mc.config.yaml"
const CONFIG_PATH = "mc-cli"

// Config reppresents a persistent configuration
type Config struct {
	Location string `yaml:"configFile"`
	Host     *Host  `yaml:"-"`
	l        hclog.Logger
}

func NewConfig(l hclog.Logger) *Config {
	c := &Config{
		Location: getDefaultConfigFilePath(),
		l:        l,
	}
	c.Host = NewHost(c)

	if err := c.Load(); err != nil {
		c.l.Debug("Error loading config", "err", err, "location", c.Location)
		c.l.Debug("Writing config file", "location", c.Location)
		err := c.Write()
		if err != nil {
			c.l.Error("Config init failed", "err", err)
		}
	} else {
		err = c.Load()
		utils.Check(err)
	}

	return c
}

func (c *Config) Write() error {
	os.MkdirAll(filepath.Dir(c.Location), os.FileMode(0755))
	fp, err := os.OpenFile(c.Location, os.O_CREATE|os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		return fmt.Errorf("Couldn't write config file: %w", err)
	}
	defer fp.Close()

	cc := codec.Codec{}
	if err = cc.Encode(fp, c, codec.YAML); err != nil {
		return err
	}
	return c.Host.Write()
}

func (c *Config) Load() error {
	fp, err := os.OpenFile(c.Location, os.O_RDWR, os.FileMode(0644))
	if err != nil {
		return err
	}
	defer fp.Close()

	dec := yaml.NewDecoder(fp)
	if err = dec.Decode(c); err != nil {
		return err
	}

	if err = c.Host.Load(); err != nil {
		return err
	}

	return nil
}

func getDefaultConfigFilePath() string {
	configDir, err := os.UserConfigDir()
	utils.Check(err)

	return filepath.Join(configDir, CONFIG_PATH, CONFIG_FILE_NAME)
}
