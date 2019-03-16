package storage

import (
	"github.com/fortytw2/dockertest"
	_ "github.com/lib/pq"
	"ports/port-domain-svc/src/config"
	"ports/port-domain-svc/src/service/storage/postgres"
)

func NewDockerStorage() (*Storage, func(), error) {
	cfg := config.Postgres{
		User:     "postgres",
		Password: "postgres",
		DbName:   "postgres",
	}
	container, runErr := dockertest.RunContainer("postgres:alpine", "5432", func(addr string) error {
		cfg.Host = addr
		db, err := postgres.Connect(cfg)
		if err != nil {
			return err
		}

		return db.Ping()
	})
	if runErr != nil {
		return nil, func() {}, runErr
	}

	db, openErr := postgres.Connect(cfg)
	if openErr != nil {
		return nil, func() {}, openErr
	}

	if migrateErr := MigrateUp(db); migrateErr != nil {
		db.Close()
		return nil, func() {}, migrateErr
	}

	s, newStorageErr := NewStorage(cfg)
	if newStorageErr != nil {
		db.Close()
		return nil, func() {}, newStorageErr
	}

	return s, func() {
		s.Close()
		container.Shutdown()
	}, nil
}
