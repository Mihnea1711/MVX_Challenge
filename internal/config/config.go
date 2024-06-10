package config

import (
	"github.com/BurntSushi/toml"
)

type Settings struct {
	LogFile string `toml:"logFile"`
}

type Keys struct {
	Key1Path     string `toml:"key1Path"`
	Key1Password string `toml:"key1Password"`
	Key2Path     string `toml:"key2Path"`
	Key2Password string `toml:"key2Password"`
}

type Config struct {
	Settings Settings `toml:"settings"`
	Keys     Keys     `toml:"keys"`
}

func LoadConfig(file string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(file, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
