package mango

// Collection is a MongoDB collection
type Collection struct {
	Type string `json:"type" bson:"type"`
	Name string `json:"name" bson:"name"`
}
