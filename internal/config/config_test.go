package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("../../config/config.toml")
	if err != nil {
		t.Errorf("Error loading config: %v", err)
	}

	assert.NotNil(t, config)
	assert.NotEmpty(t, config.Settings.LogFile)
	assert.NotEmpty(t, config.Keys.Key1Path)
	assert.NotEmpty(t, config.Keys.Key1Password)
	assert.NotEmpty(t, config.Keys.Key2Path)
	assert.NotEmpty(t, config.Keys.Key2Password)

	// Check specific values from the config.toml file
	assert.Equal(t, "logs/app.log", config.Settings.LogFile)
	assert.Equal(t, "wallets/erd136ag3fktmudhc52ssh4a46f2av0rtnywh89qlf39nsjgh6nqyzlqj56m9d.json", config.Keys.Key1Path)
	assert.Equal(t, "SomeSecretPass@2000", config.Keys.Key1Password)
	assert.Equal(t, "wallets/erd1wr5ffdllrmuq75x7stntkg27t8dk2quncd7dcdjd5ja0k4xjuegquq0tak.json", config.Keys.Key2Path)
	assert.Equal(t, "SomeOtherSecretPass@2000", config.Keys.Key2Password)
}
