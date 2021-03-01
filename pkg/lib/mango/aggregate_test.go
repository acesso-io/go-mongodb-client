package mango

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/go-playground/assert.v1"
)

func TestAggregation_Match(t *testing.T) {
	a := NewAggregation()

	q := NewQuery()
	q.Equal("key", "value")

	a.Match(q)

	assert.Equal(t, a.Pipeline(), []bson.M{
		{"$match": bson.M{"key": "value"}},
	})
}

func TestAggregation_Sort(t *testing.T) {
	a := NewAggregation()

	q := NewQuery()
	q.Equal("key", 1)

	a.Sort(q)

	assert.Equal(t, a.Pipeline(), []bson.M{
		{"$sort": bson.M{"key": 1}},
	})
}

func TestAggregation_Unwind(t *testing.T) {
	a := NewAggregation()
	a.Unwind("key")

	assert.Equal(t, a.Pipeline(), []bson.M{
		{"$unwind": "$key"},
	})
}

func TestAggregation_MultipleStages(t *testing.T) {
	a := NewAggregation()

	// Unwind stage
	a.Unwind("key")

	// Sort stage
	q1 := NewQuery()
	q1.Equal("key", 1)

	a.Sort(q1)

	// Match stage
	q2 := NewQuery()
	q2.Equal("key", "value")

	a.Match(q2)

	assert.Equal(t, a.Pipeline(), []bson.M{
		{"$unwind": "$key"},
		{"$sort": bson.M{"key": 1}},
		{"$match": bson.M{"key": "value"}},
	})
}
