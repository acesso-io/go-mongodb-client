package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
)

// Query is a MongoDB query builder. New Queries should be created via NewQuery function.
type Query struct {
	bson.M
}

// NewQuery creates a new Query
func NewQuery() *Query {
	return &Query{
		M: make(bson.M),
	}
}

// Equal adds a key-value pair to the query to match a input value in a specific key
func (q *Query) Equal(key string, value interface{}) {
	q.M[key] = value
}

// In adds a key-value pair to find entries with keys that match values in a set
func (q *Query) In(key string, values ...interface{}) {
	q.M[key] = bson.M{"$in": values}
}

// GreaterThan adds a key-value pair to find entries with keys that match values greater than the input value
func (q *Query) GreaterThan(key string, value interface{}) {
	q.M[key] = bson.M{"$gt": value}
}

// GreaterThanEqual adds a key-value pair to find entries with keys that match values greater than/equal to the input value
func (q *Query) GreaterThanEqual(key string, value interface{}) {
	q.M[key] = bson.M{"$gte": value}
}

// LessThan adds a key-value pair to find entries with keys that match values smaller than the input value
func (q *Query) LessThan(key string, value interface{}) {
	q.M[key] = bson.M{"$lt": value}
}

// LessThanEqual adds a key-value pair to find entries with keys that match values smaller than/equal to the input value
func (q *Query) LessThanEqual(key string, value interface{}) {
	q.M[key] = bson.M{"$lte": value}
}

// BSON returns the Query inner bson.M
func (q *Query) BSON() bson.M {
	return q.M
}
