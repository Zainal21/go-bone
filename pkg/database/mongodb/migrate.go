package mongodb

import (
	"context"
	"fmt"

	"github.com/Zainal21/go-bone/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func funcMigrateCollection(ctx context.Context, db *mongo.Database) []func() {
	f := []func(){
		// Key and Index for collections general_logs ================================================================
		func() {
			var (
				mIdxModels []mongo.IndexModel
			)
			collection := db.Collection(COLLECTIONS_GENERAL_LOGS)

			mIdxModels = append(mIdxModels, mongo.IndexModel{
				Keys:    map[string]int{"level": 1},
				Options: options.Index().SetName(fmt.Sprintf("idx_%s_name_asc", COLLECTIONS_GENERAL_LOGS)),
			})

			mIdxModels = append(mIdxModels, mongo.IndexModel{
				Keys:    map[string]int{"user_id": 1},
				Options: options.Index().SetName(fmt.Sprintf("idx_%s_user_id_asc", COLLECTIONS_GENERAL_LOGS)),
			})

			mIdxModels = append(mIdxModels, mongo.IndexModel{
				Keys:    map[string]int{"user_id": -1},
				Options: options.Index().SetName(fmt.Sprintf("idx_%s_user_id_desc", COLLECTIONS_GENERAL_LOGS)),
			})

			mIdxModels = append(mIdxModels, mongo.IndexModel{
				Keys:    map[string]int{"time_stamp": 1},
				Options: options.Index().SetName(fmt.Sprintf("idx_%s_time_stamp_asc", COLLECTIONS_GENERAL_LOGS)),
			})

			mIdxModels = append(mIdxModels, mongo.IndexModel{
				Keys:    map[string]int{"time_stamp": -1},
				Options: options.Index().SetName(fmt.Sprintf("idx_%s_time_stamp_desc", COLLECTIONS_GENERAL_LOGS)),
			})

			_, err := collection.Indexes().CreateMany(ctx, mIdxModels)
			if err != nil {
				logger.Warn(err)
			}

		},
		// =================================================================================================================
	}
	return f
}
