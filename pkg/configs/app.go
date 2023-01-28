package configs

import (
	"fmt"
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongo.org/mongo-driver/bson"
	"go.mongo.org/mongo-driver/mongo/options"
)

func connect() {
	client, err := mongo.NewClient(options.Client().ApplyURL(""))

}