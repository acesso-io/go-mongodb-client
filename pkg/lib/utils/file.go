package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// UnmarshallingFunc is a generic unmarshalling function
type UnmarshallingFunc func(data []byte, v interface{}) error

// Unmarshal unmarshals some data into a pointer value.
func (fn UnmarshallingFunc) Unmarshal(data []byte, v interface{}) error {
	return fn(data, v)
}

// ReadDataFromFile reads data from a file in a given path and decodes it into a given struct
func ReadDataFromFile(path string, data interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("failed to read config file at path %s", path)))
	}

	extension := filepath.Ext(path)

	var fn UnmarshallingFunc

	switch extension {
	case ".json":
		fn = json.Unmarshal
	case ".yaml", ".yml":
		fn = yaml.Unmarshal
	default:
		return errors.Wrap(ErrCannotReadFileExtension, fmt.Sprintf("failed to unmarshal %s file", extension))
	}

	if err := fn.Unmarshal(b, data); err != nil {
		log.Fatal(errors.Wrap(err, "failed to unmarshal data"))
	}

	return nil
}

// ReadDataFromBytes reads data from a byte array and decodes it into a given struct. This function does not support
// YAML and YML files.
// TODO: update function supported extensions as soon as github.com/gabriel-vasile/mimetype starts giving support to
// YAML and YML files.
func ReadDataFromBytes(b []byte, data interface{}) error {
	var fn UnmarshallingFunc

	extension := mimetype.Detect(b).Extension()

	switch extension {
	case ".json":
		fn = json.Unmarshal
	default:
		return errors.Wrap(ErrCannotReadFileExtension, fmt.Sprintf("failed to unmarshal %s file", extension))
	}

	if err := fn.Unmarshal(b, data); err != nil {
		log.Fatal(errors.Wrap(err, "failed to unmarshal data"))
	}

	return nil
}
