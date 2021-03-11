package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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

// ReadDataFromFile reads data from a file in a path and decodes it into a given data struct. The struct passed must be
// a pointer.
func ReadDataFromFile(path string, data interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = ReadDataFromBytes(b, data)
	if err != nil {
		return err
	}

	return nil
}

// ReadDataFromBytes reads data from a byte array and decodes it into a given data struct. The struct passed must be
// a pointer.
func ReadDataFromBytes(b []byte, data interface{}) error {
	var fn UnmarshallingFunc

	extension := mimetype.Detect(b).Extension()

	switch extension {
	case ".json":
		fn = json.Unmarshal
	case ".yaml", ".yml":
		fn = yaml.Unmarshal
	default:
		return errors.Wrap(ErrCannotReadFileExtension, fmt.Sprintf("failed to unmarshal %s file", extension))
	}

	if err := fn.Unmarshal(b, &data); err != nil {
		return err
	}

	return nil
}
