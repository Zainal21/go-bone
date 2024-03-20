package migration

import (
	"fmt"

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
