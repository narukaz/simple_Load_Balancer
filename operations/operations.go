package operations

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Mongo struct {
	Client *mongo.Client
}

func (m *Mongo) Get(w http.ResponseWriter, r *http.Request) (bson.M, error) {
	ctx, close := context.WithTimeout(context.TODO(), time.Second*5)
	defer close()

	coll := m.Client.Database("test").Collection("employee")

	cursor, err := coll.Find(ctx, bson.M{})

	if err != nil {
		return nil, fmt.Errorf("Operation:Finding \n %w", err)
	}
	defer cursor.Close(ctx)

	var results bson.M
	for cursor.Next(ctx) {
		if err := cursor.Decode(&results); err != nil {
			fmt.Println("error iterating over data")

			log.Fatal(err)
		}

	}
	return results, nil
}
