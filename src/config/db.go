package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (context.Context, *mongo.Client) {
	// extrayendo variables de entorno
	uri, define := os.LookupEnv("URI_MONGO_DB")
	if !define {
		log.Fatal("no defined URI_MONGO_DB env")
	}
	// conectando con mongo
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return ctx, client
}
func ConnectDB() (context.Context, *mongo.Client, *mongo.Database) {
	// extrayendo variables de entorno
	bdName, define := os.LookupEnv("BD_NAME")
	if !define {
		log.Fatal("no defined BD_NAME env")
	}
	// conectando a la base de datos
	ctx, client := Connect()
	return ctx, client, client.Database(bdName)
}
func ConnectColl(collectionName string) (context.Context, *mongo.Client, *mongo.Collection) {
	ctx, client, db := ConnectDB()
	coll := db.Collection(collectionName)
	return ctx, client, coll
}
