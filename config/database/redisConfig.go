package database

import (
	"github.com/redis/go-redis/v9"
	"os"
)

var RedisClient0 = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_HOST"),
	Password: "", // no password set
	DB:       0,  // use default DB
})

var RedisClient1 = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_HOST"),
	Password: "", // no password set
	DB:       1,  // use default DB
})
