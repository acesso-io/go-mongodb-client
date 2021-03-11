package blueprint

import (
	"github.com/acesso-io/mango/pkg/lib/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

// Blueprint represents a MongoDB databases structure
type Blueprint struct {
	Databases []Database
}

// Database is a MongoDB database blueprint
type Database struct {
	Type        string       `json:"type" bson:"type"`
	Name        string       `json:"name" bson:"name"`
	Collections []Collection `json:"collections" bson:"collections"`
}

// Collection is a MongoDB collection blueprint
type Collection struct {
	Type    string             `json:"type" bson:"type"`
	Name    string             `json:"name" bson:"name"`
	Indexes []mongo.IndexModel `json:"indexes" bson:"indexes"`
}

// ReadFromFile creates a new Blueprint from a given file
func ReadFromFile(path string) (*Blueprint, error) {
	var blueprint Blueprint

	if err := utils.ReadDataFromFile(path, &blueprint); err != nil {
		return nil, err
	}

	return &blueprint, nil
}

// ReadFromBytes creates a new Blueprint from a byte array
func ReadFromBytes(b []byte) (*Blueprint, error) {
	var blueprint Blueprint

	if err := utils.ReadDataFromBytes(b, &blueprint); err != nil {
		return nil, err
	}

	return &blueprint, nil
}
