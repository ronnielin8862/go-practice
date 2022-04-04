package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection URI
const uri = "mongodb://dev:123456@localhost:27017/great?maxPoolSize=20&w=majority"

type Test struct {
	Name     string `bson:"name"`
	Identify string `bson:"identify"`
}

func main() {
	db := ConnectDbGreat()
	TestCollection := db.Collection("test")

	//insertOne
	//test :=Test{"AAAAA","BBBBBBB"}
	//_, _ = InsertOneTest(test, TestCollection)

	//insertMany
	var tests []interface{}
	tests = append(tests, Test{"CCCCCC", "DDDDDD"})
	tests = append(tests, Test{"EEEEEEE", "FFFFFF"})
	InsertMany(tests, TestCollection)

}

func InsertMany(t []interface{}, Collection *mongo.Collection) (string, error) {
	result, err := Collection.InsertMany(context.TODO(), t)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "0", err
	}
	id := fmt.Sprintf("%v", result.InsertedIDs)
	fmt.Printf("insert one success, id =%v", id)
	return fmt.Sprintf("%v", result.InsertedIDs), err
}

func InsertOneTest(t Test, Collection *mongo.Collection) (string, error) {
	result, err := Collection.InsertOne(context.TODO(), t)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "0", err
	}
	id := fmt.Sprintf("%v", result.InsertedID)
	fmt.Printf("insert one success, id =%v", id)
	return fmt.Sprintf("%v", result.InsertedID), err
}

/*Setup opens a database connection to mongodb*/
func ConnectDbGreat() *mongo.Database {
	Ctx := context.TODO()
	host := "127.0.0.1"
	port := "27017"
	connectUri := "mongodb://" + "dev:123456@" + host + ":" + port + "/great"
	clientOptions := options.Client().ApplyURI(connectUri)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("great")
	return db
}
