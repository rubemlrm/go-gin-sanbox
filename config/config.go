package config

import (
	"errors"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config contains app configurations
type Config struct {
	Address        string `env:"ADDRESS" env-default:"8181"`
	TrustedProxies string `env:"TRUSTED_PROXIES" env-default:"0.0.0.0"`
}

// LoadConfig makes app available
func LoadConfig() (Config, error) {
	cfg := Config{}
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return cfg, errors.New("can't read env file")
	}
	return cfg, nil
}
