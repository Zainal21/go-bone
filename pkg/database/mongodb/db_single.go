package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	_ Adapter = (*DB)(nil)
)

type DB struct {
	db      *mongo.Database
	client  *mongo.Client
	reaMode bool
	dbName  string
}

func New(db *mongo.Database, reaMode bool, dbName string) *DB {
	return &DB{db: db, reaMode: reaMode, dbName: dbName}
}

func (db *DB) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return db.db.Client().Ping(ctx, rp)
}

func (db *DB) Disconnect(ctx context.Context) error {
	return db.db.Client().Disconnect(ctx)
}

func (db *DB) Collection(collection string, opts ...*options.CollectionOptions) *mongo.Collection {
	return db.db.Collection(collection, opts...)
}
