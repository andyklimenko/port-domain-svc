package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"port-domain-svc/src/config"
	"port-domain-svc/src/service/model"
)

type postgres struct {
	db *sql.DB
}

func NewPostgres(cfg config.Postgres) (*postgres, error) {
	db, dbErr := Connect(cfg)
	if dbErr != nil {
		return nil, dbErr
	}
	return &postgres{db: db}, nil
}

func Connect(cfg config.Postgres) (*sql.DB, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.DbName)
	return sql.Open("postgres", url)
}

func (p *postgres) GetPort(ctx context.Context, portCode string) (model.Port, error) {
	query := `
SELECT name, city, latitude, longitude, province, country, timezone, code from ports
WHERE port_id=$1;`
	tx, txErr := p.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if txErr != nil {
		return model.Port{}, txErr
	}

	row := tx.QueryRowContext(ctx, query, portCode)
	port := model.Port{PortCode: portCode}
	scanErr := row.Scan(&port.Name, &port.City, &port.Latitude, &port.Longitude, &port.Province,
		&port.Country, &port.Timezone, &port.Code)

	return port, scanErr
}

func (s *postgres) Close() error {
	return s.db.Close()
}

func (s *postgres) AddPorts(ctx context.Context, ports []model.Port) error {
	tx, txErr := s.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable, ReadOnly: false})
	if txErr != nil {
		return txErr
	}

	stmt, prepareErr := tx.PrepareContext(ctx, pq.CopyIn("ports",
		"port_id", "name", "city", "latitude", "longitude", "province", "country", "timezone", "code"))
	if prepareErr != nil {
		return prepareErr
	}

	for _, p := range ports {
		_, execErr := stmt.ExecContext(ctx, p.PortCode, p.Name, p.City, p.Latitude, p.Longitude, p.Province, p.Country, p.Timezone, p.Code)
		if execErr != nil {
			return execErr
		}
	}

	if closeErr := stmt.Close(); closeErr != nil {
		return closeErr
	}

	return tx.Commit()
}
