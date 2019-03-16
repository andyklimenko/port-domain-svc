package service

import (
	"context"
	"ports/port-domain-svc/src/proto"
	"ports/port-domain-svc/src/service/model"
	"ports/port-domain-svc/src/service/storage"
)

type Service struct {
	st *storage.Storage
}

func NewService(s *storage.Storage) *Service {
	return &Service{st: s}
}

func (s *Service) AddPorts(ctx context.Context, req *portdomainsvc.AddPortsReq) (*portdomainsvc.EmptyReply, error) {
	ports := make([]model.Port, 0, len(req.Ports))
	for _, p := range req.Ports {
		modelPort := model.Port{
			PortCode:  p.PortCode,
			Name:      p.Name,
			City:      p.City,
			Latitude:  p.Latitude,
			Longitude: p.Longitude,
			Country:   p.Country,
			Timezone:  p.Timezone,
			Code:      p.Code,
			Province:  p.Province,
		}
		ports = append(ports, modelPort)
	}

	return &portdomainsvc.EmptyReply{}, s.st.AddPorts(ctx, ports)
}
func (s *Service) GetPortByCode(ctx context.Context, req *portdomainsvc.GetPortByCodeReq) (*portdomainsvc.GetPortByCodeReply, error) {
	p, queryErr := s.st.GetPort(ctx, req.Code)
	if queryErr != nil {
		return nil, queryErr
	}

	grpcPort := &portdomainsvc.Port{
		PortCode:  p.PortCode,
		Name:      p.Name,
		City:      p.City,
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
		Country:   p.Country,
		Timezone:  p.Timezone,
		Code:      p.Code,
		Province:  p.Province,
	}
	return &portdomainsvc.GetPortByCodeReply{Port: grpcPort}, nil
}
