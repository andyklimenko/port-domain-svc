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

	ports := []model.Port{
		{
			PortCode:  "mt1",
			Name:      "1st port of Gondor",
			City:      "Minas-Tirith",
			Latitude:  "42",
			Longitude: "43",
			Province:  "MiddleEarth",
			Country:   "Gondor",
			Timezone:  "Minas-Tirith GMT+3",
			Code:      "1234",
		},
		{
			PortCode:  "mm1",
			Name:      "2nd port of Mordor",
			City:      "Minas-Morgul",
			Latitude:  "66",
			Longitude: "67",
			Province:  "MiddleEarth",
			Country:   "Mordor",
			Timezone:  "Minas-Morgul GMT+4",
			Code:      "9876",
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if addErr := s.AddPorts(ctx, ports); addErr != nil {
		t.Fatal(addErr)
	}

	assert.EqualError(t, s.AddPorts(ctx, ports), "pq: duplicate key value violates unique constraint \"ports_pkey\"")

	savedGondorPort, getErr := s.GetPort(ctx, "mt1")
	if getErr != nil {
		t.Fatal(getErr)
	}
	assert.Equal(t, "mt1", savedGondorPort.PortCode)
	assert.Equal(t, "1st port of Gondor", savedGondorPort.Name)
	assert.Equal(t, "Minas-Tirith", savedGondorPort.City)
	assert.Equal(t, "42", savedGondorPort.Latitude)
	assert.Equal(t, "43", savedGondorPort.Longitude)
	assert.Equal(t, "MiddleEarth", savedGondorPort.Province)
	assert.Equal(t, "Gondor", savedGondorPort.Country)
	assert.Equal(t, "Minas-Tirith GMT+3", savedGondorPort.Timezone)
	assert.Equal(t, "1234", savedGondorPort.Code)

	minasMorgulPort, getErr := s.GetPort(ctx, "mm1")
	if getErr != nil {
		t.Fatal(getErr)
	}
	assert.Equal(t, "mm1", minasMorgulPort.PortCode)
	assert.Equal(t, "2nd port of Mordor", minasMorgulPort.Name)
	assert.Equal(t, "Minas-Morgul", minasMorgulPort.City)
	assert.Equal(t, "66", minasMorgulPort.Latitude)
	assert.Equal(t, "67", minasMorgulPort.Longitude)
	assert.Equal(t, "MiddleEarth", minasMorgulPort.Province)
	assert.Equal(t, "Mordor", minasMorgulPort.Country)
	assert.Equal(t, "Minas-Morgul GMT+4", minasMorgulPort.Timezone)
	assert.Equal(t, "9876", minasMorgulPort.Code)

	_, portNotExistErr := s.GetPort(ctx, "fake code")
	assert.EqualError(t, portNotExistErr, ErrNotFound.Error())
}
