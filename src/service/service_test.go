package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"port-domain-svc/src/proto"
	"port-domain-svc/src/service/storage"
	"testing"
	"time"
)

func TestService(t *testing.T) {
	s, closer, storageErr := storage.NewDockerStorage()
	if storageErr != nil {
		t.Fatal(storageErr)
	}
	defer closer()

	service := NewService(s)
	ports := []*portdomainsvc.Port{
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if _, addPortsErr := service.AddPorts(ctx, &portdomainsvc.AddPortsReq{Ports: ports}); addPortsErr != nil {
		t.Fatal(addPortsErr)
	}

	_, addPortsErr := service.AddPorts(ctx, &portdomainsvc.AddPortsReq{Ports: ports})
	assert.EqualError(t, addPortsErr, "pq: duplicate key value violates unique constraint \"ports_pkey\"")

	minasTirith, getErr := service.GetPortByCode(ctx, &portdomainsvc.GetPortByCodeReq{Code: "mt1"})
	if getErr != nil {
		t.Fatal(getErr)
	}
	assert.Equal(t, "mt1", minasTirith.Port.PortCode)
	assert.Equal(t, "1st port of Gondor", minasTirith.Port.Name)
	assert.Equal(t, "Minas-Tirith", minasTirith.Port.City)
	assert.Equal(t, "42", minasTirith.Port.Latitude)
	assert.Equal(t, "43", minasTirith.Port.Longitude)
	assert.Equal(t, "MiddleEarth", minasTirith.Port.Province)
	assert.Equal(t, "Gondor", minasTirith.Port.Country)
	assert.Equal(t, "Minas-Tirith GMT+3", minasTirith.Port.Timezone)
	assert.Equal(t, "1234", minasTirith.Port.Code)

	minasMorgul, getErr := service.GetPortByCode(ctx, &portdomainsvc.GetPortByCodeReq{Code: "mm1"})
	if getErr != nil {
		t.Fatal(getErr)
	}
	assert.Equal(t, "mm1", minasMorgul.Port.PortCode)
	assert.Equal(t, "2nd port of Mordor", minasMorgul.Port.Name)
	assert.Equal(t, "Minas-Morgul", minasMorgul.Port.City)
	assert.Equal(t, "66", minasMorgul.Port.Latitude)
	assert.Equal(t, "67", minasMorgul.Port.Longitude)
	assert.Equal(t, "MiddleEarth", minasMorgul.Port.Province)
	assert.Equal(t, "Mordor", minasMorgul.Port.Country)
	assert.Equal(t, "Minas-Morgul GMT+4", minasMorgul.Port.Timezone)
	assert.Equal(t, "9876", minasMorgul.Port.Code)

	_, portNotExistErr := s.GetPort(ctx, "fake code")
	assert.EqualError(t, portNotExistErr, storage.ErrNotFound.Error())
}
