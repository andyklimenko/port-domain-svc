package storage

import (
	"database/sql"
	"github.com/rubenv/sql-migrate"
)

var migrations = migrate.MemoryMigrationSource{
	Migrations: []*migrate.Migration{
		{
			Id: "0001_initial",
			Up: []string{
				`CREATE TABLE ports (
					port_id VARCHAR(50) PRIMARY KEY,
					name VARCHAR(50) NOT NULL,
					city VARCHAR(50) NOT NULL,
					latitude VARCHAR(50) NOT NULL,
					longitude VARCHAR(50) NOT NULL,
					province VARCHAR(50) NOT NULL,
					country VARCHAR(50) NOT NULL,
					timezone VARCHAR(50) NOT NULL,
					code VARCHAR(50) NOT NULL
				);`,
			},
			Down: []string{
				"DROP TABLE IF EXISTS ports;",
			},
		},
	},
}

func MigrateUp(db *sql.DB) error {
	_, migrateErr := migrate.Exec(db, "postgres", migrations, migrate.Up)
	return migrateErr
}
