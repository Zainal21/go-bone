package migration

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Zainal21/go-bone/database/seeders"
	"github.com/Zainal21/go-bone/pkg/database/mysql"

	"github.com/Zainal21/go-bone/pkg/config"
	"github.com/Zainal21/go-bone/pkg/logger"
)

func MigrateDatabase() {
	cfg, err := config.LoadAllConfigs()

	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to load configuration file: %v", err))
	}

	mysql.DatabaseMigration(cfg)
}

func SeedDatabase() {
	cfg, err := config.LoadAllConfigs()

	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to load configuration file: %v", err))
	}
	flag.Parse()
	args := flag.Args()
	if len(args) >= 1 {
		log.Println(args[0])
		switch args[0] {
		case "db:seed":
			db, err := mysql.ConnectDatabase(cfg)
			if err != nil {
				logger.Fatal(fmt.Sprintf("Failed to load configuration file: %v", err))
			}

			seeders.Execute(db, args[1:]...)
			os.Exit(0)
		}
	}
}
