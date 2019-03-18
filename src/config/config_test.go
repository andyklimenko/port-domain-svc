package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadConfigFromEnv(t *testing.T) {
	assert.NoError(t, os.Setenv("POSTGRES_HOST", "test-host"))
	assert.NoError(t, os.Setenv("POSTGRES_PORT", "1234"))
	assert.NoError(t, os.Setenv("POSTGRES_USER", "user"))
	assert.NoError(t, os.Setenv("POSTGRES_PASSWORD", "pwd"))
	assert.NoError(t, os.Setenv("POSTGRES_DBNAME", "test-db"))

	cfg := Config{}
	cfg.Load()
	assert.Equal(t, "test-host", cfg.Pg.Host)
	assert.Equal(t, 1234, cfg.Pg.Port)
	assert.Equal(t, "user", cfg.Pg.User)
	assert.Equal(t, "pwd", cfg.Pg.Password)
	assert.Equal(t, "test-db", cfg.Pg.DbName)
	assert.Equal(t, ":50051", cfg.Port)
}
