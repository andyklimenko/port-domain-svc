package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStorage(t *testing.T) {
	_, closer, storageErr := NewDockerStorage()
	defer closer()
	assert.NoError(t, storageErr)
}
