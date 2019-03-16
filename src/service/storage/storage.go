package storage

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"ports/port-domain-svc/src/service/model"
	"ports/port-domain-svc/src/service/storage/postgres"
)

var ErrNotFound = errors.New("not found")

type db interface {
	GetPort(ctx context.Context, portCode string) (model.Port, error)
	AddPorts(ctx context.Context, p []model.Port) error
	Close() error
}

type Storage struct {
	db db
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: postgres.NewPostgres(db)}
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) GetPort(ctx context.Context, portCode string) (model.Port, error) {
	p, queryErr := s.db.GetPort(ctx, portCode)
	if queryErr == sql.ErrNoRows {
		return model.Port{}, ErrNotFound
	}
	return p, queryErr
}

func (s *Storage) AddPorts(ctx context.Context, p []model.Port) error {
	return s.db.AddPorts(ctx, p)
}
