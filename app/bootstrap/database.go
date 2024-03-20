package bootstrap

import (
	"fmt"

	"github.com/Zainal21/go-bone/pkg/config"
	"github.com/Zainal21/go-bone/pkg/database/mongodb"
	"github.com/Zainal21/go-bone/pkg/database/mysql"
	"github.com/Zainal21/go-bone/pkg/logger"
)

func RegistryDatabase(cfg *config.Config) *mysql.DB {
	// Remove this code below if no need database
	db, err := mysql.ConnectDatabase(cfg)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	return mysql.New(db, false, cfg.DatabaseConfig.DBName)
}

func RegistryMongoDB(cfg *config.Config) *mongodb.DB {
	db, err := mongodb.Connect(cfg)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	return mongodb.New(db, false, cfg.DatabaseConfig.DBName)
}
