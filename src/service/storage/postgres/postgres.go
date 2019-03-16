package postgres

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"ports/port-domain-svc/src/service/model"
)

type postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *postgres {
	return &postgres{db: db}
}

func (p *postgres) GetPort(ctx context.Context, portCode string) (model.Port, error) {
	query := `
SELECT name, city, latitude, longitude, province, timezone, code from ports
WHERE port_id=$1;`
	tx, txErr := p.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if txErr != nil {
		return model.Port{}, txErr
	}

	row := tx.QueryRowContext(ctx, query, portCode)
	port := model.Port{PortCode: portCode}
	scanErr := row.Scan(&port.Name, &port.City, &port.Latitude, &port.Longitude, &port.Province, &port.Timezone, &port.Code)
	return port, scanErr
}

func (s *postgres) Close() error {
	return s.db.Close()
}

func (s *postgres) AddPort(ctx context.Context, p model.Port) error {
	query := `
INSERT INTO ports (port_id, name, city, latitude, longitude, province, timezone, code)
VALUES($1, $2, $3, $4, $5, $6, $7, $8);`
	tx, txErr := s.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable, ReadOnly: false})
	if txErr != nil {
		return txErr
	}

	_, execErr := tx.ExecContext(ctx, query, p.PortCode, p.Name, p.City, p.Latitude, p.Longitude, p.Province, p.Timezone, p.Code)
	if execErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return errors.Wrap(execErr, rollbackErr.Error())
		}
		return execErr
	}
	return tx.Commit()
}
