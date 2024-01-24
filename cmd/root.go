package cmd

import (
	"log"

	"github.com/Zainal21/go-bone/cmd/http"
	"github.com/Zainal21/go-bone/cmd/migration"
	"github.com/spf13/cobra"
)

func Start() {

	rootCmd := &cobra.Command{}

	migrateCmd := &cobra.Command{
		Use:   "db:migrate",
		Short: "database migration",
		Run: func(c *cobra.Command, args []string) {
			migration.MigrateDatabase()
		},
	}

	migrateCmd.Flags().BoolP("version", "", false, "print version")
	migrateCmd.Flags().StringP("dir", "", "database/migration/", "directory with migration files")
	migrateCmd.Flags().StringP("table", "", "db", "migrations table name")
	migrateCmd.Flags().BoolP("verbose", "", false, "enable verbose mode")
	migrateCmd.Flags().BoolP("guide", "", false, "print help")

	seederCmd := &cobra.Command{
		Use:   "db:seed",
		Short: "Run database seeder",
		Run: func(c *cobra.Command, args []string) {
			migration.SeedDatabase()
		},
	}
	seederCmd.Flags().BoolP("version", "", false, "print version")

	cmd := []*cobra.Command{
		{
			Use:   "http",
			Short: "Run HTTP Server",
			Run: func(cmd *cobra.Command, args []string) {
				http.Start()
			},
		},
		migrateCmd,
		seederCmd,
	}

	rootCmd.AddCommand(cmd...)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
