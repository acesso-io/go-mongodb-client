package blueprint

import (
	"github.com/acesso-io/mango/pkg/lib/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

// Environment represents a MongoDB databases environment
type Environment struct {
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

// NewEnvironmentFromFile creates a new Environment from a given file
func NewEnvironmentFromFile(path string) (*Environment, error) {
	var environment Environment

	if err := utils.ReadDataFromFile(path, &environment); err != nil {
		return nil, err
	}

	return &environment, nil
}

// NewEnvironmentFromBytes creates a new Environment from a byte array
func NewEnvironmentFromBytes(b []byte) (*Environment, error) {
	var environment Environment

	if err := utils.ReadDataFromBytes(b, &environment); err != nil {
		return nil, err
	}

	return &environment, nil
}

// NewDatabaseFromFile creates a new Database blueprint from a given file
func NewDatabaseFromFile(path string) (*Database, error) {
	var database Database

	if err := utils.ReadDataFromFile(path, &database); err != nil {
		return nil, err
	}

	return &database, nil
}

// NewDatabaseFromBytes creates a new Database blueprint from a byte array
func NewDatabaseFromBytes(b []byte) (*Database, error) {
	var database Database

	if err := utils.ReadDataFromBytes(b, &database); err != nil {
		return nil, err
	}

	return &database, nil
}
