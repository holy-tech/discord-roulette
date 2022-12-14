package repo

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var uri = fmt.Sprintf("mongodb://%s:%s@%s:%s",
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASS"),
	os.Getenv("DB_IP"),
	os.Getenv("DB_PORT"),
)

func init() {
	var err error
	Client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Could not return DB connection: %v", err)
	}

}
