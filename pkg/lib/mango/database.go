package mango

import (
	"github.com/acesso-io/mango/pkg/lib/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

// Database extends mongo.Database
type Database struct {
	Type        string       `json:"type" bson:"type"`
	Name        string       `json:"name" bson:"name"`
	Collections []Collection `json:"collections" bson:"collections"`

	*mongo.Database
}

// ReadDatabaseFromFile reads a list of Collection from a given file.
func ReadDatabaseFromFile(path string) (*Database, error) {
	var database Database

	err := utils.ReadDataFromFile(path, &database)
	if err != nil {
		return nil, err
	}

	return &database, nil
}
