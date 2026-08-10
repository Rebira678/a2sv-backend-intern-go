package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Trainer struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
	City string `bson:"city"`
}

func main() {
	// 1. Set client options and connect

	// "246810Re;:::" encoded becomes "246810Re%3B%3A%3A%3A"
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Ping to verify connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}
	fmt.Println("Connected to MongoDB!")

	// Ensure client disconnects when main exits
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection closed.")
	}()

	// 3. Get collection handle
	collection := client.Database("test").Collection("trainers")

	// Run CRUD steps
	executeCRUD(collection)
}

func executeCRUD(collection *mongo.Collection) {
	ctx := context.TODO()

	// --- CREATE (Insert One & Insert Many) ---
	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}

	insertRes, err := collection.InsertOne(ctx, ash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted Single ID:", insertRes.InsertedID)

	trainers := []interface{}{misty, brock}
	insertManyRes, err := collection.InsertMany(ctx, trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted Multiple IDs:", insertManyRes.InsertedIDs)

	// --- UPDATE (Update One) ---
	filter := bson.D{{"name", "Ash"}}
	update := bson.D{{"$inc", bson.D{{"age", 1}}}}

	updateRes, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v and Modified %v\n", updateRes.MatchedCount, updateRes.ModifiedCount)

	// --- READ (Find One) ---
	var ashResult Trainer
	err = collection.FindOne(ctx, filter).Decode(&ashResult)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found Single Document: %+v\n", ashResult)

	// --- READ (Find Multiple via Cursor) ---
	findOpts := options.Find().SetLimit(2)
	cur, err := collection.Find(ctx, bson.D{{}}, findOpts)
	if err != nil {
		log.Fatal(err)
	}

	var results []*Trainer
	for cur.Next(ctx) {
		var elem Trainer
		if err := cur.Decode(&elem); err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	cur.Close(ctx)
	fmt.Printf("Found Multiple Documents: %d retrieved\n", len(results))

	// --- DELETE (Delete Many) ---
	deleteRes, err := collection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents\n", deleteRes.DeletedCount)
}
