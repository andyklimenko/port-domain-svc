package storage

import (
	"github.com/fortytw2/dockertest"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"port-domain-svc/src/config"
	"port-domain-svc/src/service/storage/postgres"
	"strconv"
	"strings"
)

func NewDockerStorage() (*Storage, func(), error) {
	cfg := config.Postgres{
		User:     "postgres",
		Password: "postgres",
		DbName:   "postgres",
	}
	container, runErr := dockertest.RunContainer("postgres:alpine", "5432", func(addr string) error {
		cfg.Host = addr
		splitted := strings.Split(addr, ":")
		if len(splitted) != 2 {
			return errors.New("wrong addr format")
		}

		port, convErr := strconv.Atoi(splitted[1])
		if convErr != nil {
			return convErr
		}

		cfg.Host = splitted[0]
		cfg.Port = port
		if _, connectErr := postgres.Connect(cfg); connectErr != nil {
			return connectErr
		}

		return nil
	})
	if runErr != nil {
		return nil, func() {}, runErr
	}

	db, connectErr := postgres.Connect(cfg)
	if connectErr != nil {
		return nil, func() {}, connectErr
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
