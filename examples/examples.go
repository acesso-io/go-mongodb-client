package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-mongodb-client/pkg/lib/mongodb"
	"go-mongodb-client/pkg/lib/mongodb/uuid"

	"github.com/pkg/errors"
)

func main() {
	client := mongodb.NewClient(
		mongodb.NewOptionsFromConfigFile("config.yaml"),
	)

	if err := client.Connect(); err != nil {
		log.Fatal(errors.Wrap(err, "failed to connect to MongoDB"))
	}

	db := client.Database("my_database")

	// Build new aggregation pipeline
	aggregation := mongodb.NewAggregation()

	// Start new query builder
	q1 := mongodb.NewQuery()
	q1.Equal("_id", uuid.Must(uuid.Parse("548c8a3b-2906-4693-a601-98dcf8225a25")))
	q1.In("meta.generation", 1, 2, 3)

	// Add a $match stage to the aggregation pipeline using the query built
	aggregation.Match(q1)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := db.Collection("my_collection").Aggregate(ctx, aggregation.Pipeline())
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get query results"))
	}

	var results []map[string]interface{}

	if err := cursor.All(context.Background(), &results); err != nil {
		log.Fatal(errors.Wrap(err, "failed to decode query results"))
	}

	fmt.Printf("%+v\n", results)
}
