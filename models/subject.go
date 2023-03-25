package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Subject struct {
	Id   string `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
}

var subjectModel = ConnectDatabase().Database("test").Collection("subjects")

func StoreSubject(payload Subject) *mongo.InsertOneResult {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := subjectModel.InsertOne(ctx, Subject{
		Name: payload.Name,
	})
	if err != nil {
		fmt.Println("Error when inserting object")
	}
	return res
}

func GetSubjectByID(id string) Subject {
	var subject Subject
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := subjectModel.FindOne(ctx, bson.D{{"_id", ConvertToObjId(id)}}).Decode(&subject)
	if err != nil {
		return subject
		fmt.Println("ERROR FETC BY ID %v", err)
	}
	return subject
}
