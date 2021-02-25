package mongodb

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// Aggregation is a aggregate pipeline builder for MongoDB queries. New Aggregations should be created via NewAggregation function.
type Aggregation struct {
	pipeline []bson.M
}

// NewAggregation creates a new Aggregation
func NewAggregation() *Aggregation {
	return &Aggregation{
		pipeline: make([]bson.M, 0),
	}
}

// Pipeline returns the pipeline built around an Aggregation
func (a *Aggregation) Pipeline() []bson.M {
	return a.pipeline
}

// Match adds a $match stage to the current Aggregation
func (a *Aggregation) Match(query *Query) {
	a.pipeline = append(a.pipeline, bson.M{"$match": query.BSON()})
}

// Sort adds a $sort stage to the current Aggregation
func (a *Aggregation) Sort(query *Query) {
	a.pipeline = append(a.pipeline, bson.M{"$sort": query.BSON()})
}

// Unwind adds a $unwind stage to the current Aggregation
func (a *Aggregation) Unwind(key string) {
	a.pipeline = append(a.pipeline, bson.M{"$unwind": fmt.Sprintf("$%s", key)})
}
