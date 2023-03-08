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

func StoreTeacherMongo(payload Teacher) *mongo.InsertOneResult {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	docs := []Teacher{{
		Code: payload.Code,
		Name: payload.Name,
	}}
	res, err := collection.InsertOne(ctx, docs)
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

func FindTeacherByIdMongo(id string) Teacher {
	var teacher Teacher
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.D{{"_id", ConvertToObjId(id)}}).Decode(&teacher)
	if err != nil {
		fmt.Println("ERROR FETC BY ID %v", err)
	}
	return teacher
}

func UpdateTeacherById(id string, payload Teacher) Teacher {
	var teacher Teacher
	fmt.Println("WHAT IS NAME %v", payload.Name)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := collection.UpdateOne(ctx, bson.D{{"_id", ConvertToObjId(id)}}, bson.D{
		{
			"$set", Teacher{
				Code: payload.Code,
				Name: payload.Name,
			},
		},
	})
	if err != nil {
		fmt.Println("ERROR UPDATE")
	}
	teacher = FindTeacherByIdMongo(id)
	return teacher
}

func DeleteTeacherByID(id string) bool {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := collection.DeleteOne(ctx, bson.D{{"_id", ConvertToObjId(id)}})
	if err != nil {
		fmt.Println("ERROR DELETE %v", err)
	}
	return res.DeletedCount > 0
}
