package cmd

import (
	"github.com/spf13/cobra"
	"port-domain-svc/src/service/storage"
	"port-domain-svc/src/service/storage/postgres"
)

var (
	migrateCmd = &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			cfg.Load()

			db, connectErr := postgres.Connect(cfg.Pg)
			if connectErr != nil {
				panic(connectErr)
			}

			if migrateErr := storage.MigrateUp(db); migrateErr != nil {
				panic(migrateErr)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}
