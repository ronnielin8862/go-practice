package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	//test :=Test{"A","B"}
	//_, _ = InsertOneTest(test, TestCollection)

	//insertMany
	//var tests []interface{}
	//tests = append(tests, Test{"C", "D"})
	//tests = append(tests, Test{"E", "F"})
	//InsertMany(tests, TestCollection)

	//GetOneById
	//GetTestById("624acea08370a5e5a98427cb",TestCollection)

	//GetByColumn
	//GetAll
	//d := bson.D{{"name", "AAAAA"}}
	//GetTestByColumn(d, TestCollection)

	//UpdateById
	//id := "624acea08370a5e5a98427cb"
	//d := bson.D{{"name", "EEEEeee"}, {"identify", "FFFFFFfff"}}
	//UpdateTest(id, TestCollection, d)

	//deleteById
	DeleteTest("624ad285bbce277fb8fea11e", TestCollection)
}

func DeleteTest(id string, Collection *mongo.Collection) {
	objectId, err := primitive.ObjectIDFromHex(id)
	_, err = Collection.DeleteOne(context.TODO(), bson.D{{"_id", objectId}})
	if err != nil {
		fmt.Printf("Error")
	}
}

func UpdateTest(id string, Collection *mongo.Collection, d bson.D) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error")
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", d}}
	_, _ = Collection.UpdateOne(
		context.TODO(),
		filter,
		update,
	)
}
func GetTestByColumn(d bson.D, Collection *mongo.Collection) {
	var test Test
	var tests []Test
	Ctx := context.TODO()

	//by column
	//cursor, err := Collection.Find(Ctx, d)
	//all
	cursor, err := Collection.Find(Ctx, bson.D{})
	if err != nil {
		defer cursor.Close(Ctx)
		fmt.Printf("Error")
	}

	for cursor.Next(Ctx) {
		err := cursor.Decode(&test)
		if err != nil {
			fmt.Printf("Error")
		}
		fmt.Printf("Gets : %v \n", test)
		tests = append(tests, test)
	}
}

func GetTestById(id string, Collection *mongo.Collection) {
	var test Test
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error")
	}

	err = Collection.
		FindOne(context.TODO(), bson.D{{"_id", objectId}}).
		Decode(&test)
	if err != nil {
		fmt.Printf("Error")
	}
	fmt.Printf("Get One : %v", test)
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
