package storage

import (
	"database/sql"
	"github.com/fortytw2/dockertest"
	_ "github.com/lib/pq"
)

func NewDockerStorage() (*Storage, func(), error) {
	container, runErr := dockertest.RunContainer("postgres:alpine", "5432", func(addr string) error {
		db, err := sql.Open("postgres", "postgres://postgres:postgres@"+addr+"?sslmode=disable")
		if err != nil {
			return err
		}

		return db.Ping()
	})
	if runErr != nil {
		return nil, func() {}, runErr
	}

	db, openErr := sql.Open("postgres", "postgres://postgres:postgres@"+container.Addr+"?sslmode=disable")
	if openErr != nil {
		return nil, func() {}, openErr
	}

	if migrateErr := MigrateUp(db); migrateErr != nil {
		db.Close()
		return nil, func() {}, migrateErr
	}

	s := NewStorage(db)
	return s, func() {
		s.Close()
		container.Shutdown()
	}, nil
}
