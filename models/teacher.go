package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Teacher struct {
	Code string `bson:"code,omitempty"`
	Name string `bson:"name,omitempty"`
}

func StoreTeacherMongo(payload Teacher) *mongo.InsertManyResult {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	docs := []interface{}{
		bson.D{{"code", payload.Code}, {"name", payload.Name}},
	}
	// databases, err := ConnectDatabase().ListDatabaseNames(ctx, bson.M{})
	collection := ConnectDatabase().Database("test").Collection("teachers")
	res, err := collection.InsertMany(ctx, docs)
	if err != nil {
		fmt.Println("ERROR WOY %v", err)
	}
	// return res, insertErr
	return res
}
