package mango

import (
	"github.com/acesso-io/mango/pkg/lib/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

// Collection extends mongo.Collection
type Collection struct {
	Type    string             `json:"type" bson:"type"`
	Name    string             `json:"name" bson:"name"`
	Indexes []mongo.IndexModel `json:"indexes" bson:"indexes"`

	*mongo.Collection
}

// ReadCollectionsFromFile reads a list of Collection from a given file.
func ReadCollectionsFromFile(path string) ([]Collection, error) {
	var collections []Collection

	err := utils.ReadDataFromFile(path, &collections)
	if err != nil {
		return nil, err
	}

	return collections, nil
}
