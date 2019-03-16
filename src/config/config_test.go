package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadConfigFromEnv(t *testing.T) {
	assert.NoError(t, os.Setenv("DB_HOST", "test-host"))
	assert.NoError(t, os.Setenv("DB_PORT", "1234"))
	assert.NoError(t, os.Setenv("DB_USER", "user"))
	assert.NoError(t, os.Setenv("DB_PASSWORD", "pwd"))
	assert.NoError(t, os.Setenv("DB_DBNAME", "test-db"))

	cfg := &Postgres{}
	cfg.Load()
	assert.Equal(t, "test-host", cfg.Host)
	assert.Equal(t, 1234, cfg.Port)
	assert.Equal(t, "user", cfg.User)
	assert.Equal(t, "pwd", cfg.Password)
	assert.Equal(t, "test-db", cfg.DbName)
}
