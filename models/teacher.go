package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection = ConnectDatabase().Database("test").Collection("teachers")

type Teacher struct {
	Id   string `bson:"_id,omitempty"`
	Code string `bson:"code,omitempty"`
	Name string `bson:"name,omitempty"`
}

func StoreTeacherMongo(payload Teacher) *mongo.InsertManyResult {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	docs := []interface{}{
		bson.D{{"code", payload.Code}, {"name", payload.Name}},
	}
	res, err := collection.InsertMany(ctx, docs)
	if err != nil {
		fmt.Println("ERROR WOY %v", err)
	}
	return res
}

func ListTeacherMongo() []Teacher {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	defer cur.Close(ctx)

	var teachers []Teacher
	cur.All(ctx, &teachers)
	return teachers
}
