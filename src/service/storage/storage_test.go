package storage

import (
	"context"
	"github.com/stretchr/testify/assert"
	"ports/port-domain-svc/src/service/model"
	"testing"
)

func TestStorage(t *testing.T) {
	s, closer, storageErr := NewDockerStorage()
	defer closer()
	assert.NoError(t, storageErr)

	p := model.Port{
		PortCode:  "Gondor1",
		Name:      "1st port of Gondor",
		City:      "Minas-Tirith",
		Latitude:  "42",
		Longitude: "43",
		Province:  "MiddleEarth",
		Timezone:  "Gondor GMT+3",
		Code:      "1234",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if addErr := s.AddPort(ctx, p); addErr != nil {
		t.Fatal(addErr)
	}

	assert.EqualError(t, s.AddPort(ctx, p), "pq: duplicate key value violates unique constraint \"ports_pkey\"")

	savedPort, getErr := s.GetPort(ctx, "Gondor1")
	if getErr != nil {
		t.Fatal(getErr)
	}

	assert.Equal(t, "Gondor1", savedPort.PortCode)
	assert.Equal(t, "1st port of Gondor", savedPort.Name)
	assert.Equal(t, "Minas-Tirith", savedPort.City)
	assert.Equal(t, "42", savedPort.Latitude)
	assert.Equal(t, "43", savedPort.Longitude)
	assert.Equal(t, "MiddleEarth", savedPort.Province)
	assert.Equal(t, "Gondor GMT+3", savedPort.Timezone)
	assert.Equal(t, "1234", savedPort.Code)

	_, portNotExistErr := s.GetPort(ctx, "fake code")
	assert.EqualError(t, portNotExistErr, ErrNotFound.Error())
}
