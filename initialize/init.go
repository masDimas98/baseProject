package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"goTest/config/database"
	"time"
)

var Db *mongo.Database
var Ctx context.Context
var RClient0 *redis.Client
var RClient1 *redis.Client

func init() {
	RClient0 = database.RedisClient0.WithTimeout(10 * time.Second)
	RClient1 = database.RedisClient1.WithTimeout(10 * time.Second)
	Db = database.MClient
	Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

}
