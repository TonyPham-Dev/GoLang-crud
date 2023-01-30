package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(Getenv("DB_URL")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

var DB *mongo.Client = ConnectDatabase()

func GetConlections(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("db-golang").Collection(collectionName)
	return collection
}

// config env
func Getenv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Cannot loading environment ")
	}
	return os.Getenv(key)
}
