package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goTest/config/utils"
	"os"
)

var clientOptions *options.ClientOptions
var Client, _ *mongo.Client
var MClient *mongo.Database

func init() {
	utils.LoadEnv()
	clientOptions = options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	Client, _ = mongo.Connect(context.Background(), clientOptions)
	MClient = Client.Database(os.Getenv("MONGODB_DATABASE"))
}
