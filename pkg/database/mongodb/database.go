package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/Zainal21/go-bone/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect(cfg *config.Config) (*mongo.Database, error) {
	mongoDBConfig := cfg.MongoDBConfig
	url := fmt.Sprintf("mongodb://%v:%v", mongoDBConfig.DBHost, mongoDBConfig.DBPort)
	opts := options.Client().
		ApplyURI(url).
		SetAuth(options.Credential{
			AuthSource:  "admin",
			Username:    mongoDBConfig.DBUser,
			Password:    mongoDBConfig.DBPassword,
			PasswordSet: true,
		}).
		SetMaxConnIdleTime(time.Duration(mongoDBConfig.MaxIdleTime) * time.Hour).
		SetConnectTimeout(time.Duration(mongoDBConfig.MaxConnLifetime) * time.Hour).
		SetMaxConnecting(uint64(mongoDBConfig.MaxOpenConn))

	tlsConfig, err := mongoDBConfig.TlsConfig(cfg.AppEnv)
	if err != nil {
		return nil, err
	}

	if tlsConfig != nil {
		opts.SetTLSConfig(tlsConfig)
	}

	ctx := context.TODO()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	cd := client.Database(mongoDBConfig.DBName)

	// Async execute migrations mongodb collection
	go func() {
		fncs := funcMigrateCollection(ctx, cd)
		for _, fn := range fncs {
			fn()
		}
	}()
	// ============================================

	return cd, err
}

func Connect(cfg *config.Config) (*mongo.Database, error) {
	return connect(cfg)
}
