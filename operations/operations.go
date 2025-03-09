package operations

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
}

type Data struct {
	ID    bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string        `json:"name,omitempty" bson:"name,omitempty"`
	Age   int           `json:"age,omitempty" bson:"age,omitempty"`
	Phone int           `json:"phone,omitempty" bson:"phone,omitempty"`
}

func ConnectToMongo(URI string) (*mongo.Client, error) {
	ctx, close := context.WithTimeout(context.TODO(), time.Second*4)
	defer close()

	client, err := mongo.Connect(options.Client().ApplyURI(URI))
	if err != nil {
		fmt.Println("Failed to form client")
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Failed to ping client")
		return nil, err
	}
	fmt.Println("connection established")
	return client, nil

}
func (m *Mongo) Get(w http.ResponseWriter, r *http.Request) {
	ctx, close := context.WithTimeout(context.TODO(), time.Second*5)
	defer close()

	coll := m.Client.Database("test").Collection("employee")

	cursor, err := coll.Find(ctx, bson.M{})

	if err != nil {
		return
	}
	defer cursor.Close(ctx)

	var results bson.M
	for cursor.Next(ctx) {
		if err := cursor.Decode(&results); err != nil {
			fmt.Println("error iterating over data")
			log.Fatal(err)
		}

	}
	jsonData, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON Output:", string(jsonData))

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonData))

}

func (m *Mongo) Delete(w http.ResponseWriter, r *http.Request) {
	var req Data

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid data")
		return
	}

	ctx, close := context.WithTimeout(context.TODO(), time.Second*5)
	defer close()

	coll := m.Client.Database("test").Collection("employee")
	var myObject any
	err := coll.FindOne(ctx, bson.M{"_id": req.ID}).Decode(myObject)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed decoding results")
		return
	}
	res, err := coll.DeleteOne(ctx, bson.M{"_id": req.ID})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error deleting document")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "delete count is : %d", res.DeletedCount)

}

func (m *Mongo) Add(w http.ResponseWriter, r *http.Request) {
	var req bson.M

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid data")
		return
	}

	ctx, close := context.WithTimeout(context.TODO(), time.Second*5)
	defer close()

	coll := m.Client.Database("test").Collection("employee")

	res, err := coll.InsertOne(ctx, req)
	if err != nil {
		fmt.Fprint(w, "Error inserting document")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Inserted Id is : %d", res.InsertedID)

}
