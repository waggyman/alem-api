package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var client *mongo.Client

func ConnectDatabase() *mongo.Client {
	// Test1234
	godotenv.Load(".env")
	fmt.Println(os.Getenv("MONGODB_URI"))
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		// fmt.Println("ADA ERROR DONG %v", err)
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		// fmt.Println("ADA ERROR DONG PAS LIST DB %v", err)
		log.Fatal(err)
	}
	fmt.Println(databases)
	return client
}
