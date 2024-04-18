package config

type ConfigKey string
type ConfigValue string
type DefaultValue string

type Configuration interface {
	Write() error
	Get(ConfigKey) (string, error)
	Load() error
}

// Config reppresents a persistent configuration
type Config struct {
	Authentication Configuration
}

func NewConfig() *Config {
	c := &Config{
		Authentication: NewAuthConfig(),
	}

	return c
}
