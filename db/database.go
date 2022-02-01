package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var clientmongo *mongo.Client

func ConnectDB() *mongo.Database {

	opt := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, opt)

	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		panic(err)
	}

	log.Println("Success Connect To DB")

	return client.Database("xendit")

}

func CloseDB() {
	if clientmongo == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := clientmongo.Disconnect(ctx); err != nil {
		panic(err)
	}

	log.Println("Connection Close")

}
