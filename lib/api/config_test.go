package api_test

import (
	"testing"

	. "github.com/arv28/wisdom-service/lib/api"
	"github.com/stretchr/testify/assert"
)

//-------------------------------------------------------------------------------------------------

func TestNewConfig(t *testing.T) {
	cfg := NewConfig("env", "0.0.0.0", 1234, "api-path", "web-dir")
	assert.Equal(t, "env", cfg.Env)
	assert.Equal(t, "0.0.0.0", cfg.Host)
	assert.Equal(t, 1234, cfg.Port)
	assert.Equal(t, "api-path", cfg.APIPath)
	assert.Equal(t, "web-dir", cfg.WebDir)
}

//-------------------------------------------------------------------------------------------------

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	assert.Equal(t, "dev", cfg.Env)
	assert.Equal(t, "localhost", cfg.Host)
	assert.Equal(t, 4000, cfg.Port)
	assert.Equal(t, "/", cfg.APIPath)
	assert.Equal(t, "web", cfg.WebDir)
}
