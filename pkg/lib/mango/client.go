package mango

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"strings"

	"github.com/acesso-io/mango/pkg/lib/utils"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client extends MongoDB's client
type Client struct {
	*mongo.Client

	options *Options
}

// NewClient creates a new MongoDB Client
func NewClient(opts *Options) *Client {
	if opts == nil {
		opts = &DefaultOptions
	}

	return &Client{
		options: opts,
	}
}

// Connect establishes a connection between the client and MongoDB servers
func (c *Client) Connect() error {
	if c.isConnected() {
		return nil
	}

	client, err := mongo.Connect(context.Background(), c.options.clientOptions()...)
	if nil != err {
		log.Fatal(errors.Wrap(err, "failed to connect to MongoDB servers"))
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(errors.Wrap(err, "failed to ping MongoDB servers"))
	}

	c.Client = client

	return nil
}

// Disconnect interrupts an existing connection
func (c *Client) Disconnect() error {
	if c.isConnected() {
		return c.Client.Disconnect(context.Background())
	}

	return nil
}

// isConnected returns true if the current client has a connection available with MongoDB servers
func (c *Client) isConnected() bool {
	if c.Client != nil {
		return true
	}

	return false
}

// Options are options for the MongoDB Client
type Options struct {
	URI          string `json:"uri" yaml:"uri"`
	ClientPEM    string `json:"client_pem" yaml:"client_pem"`
	ClientKeyPEM string `json:"client_key_pem" yaml:"client_key_pem"`
	CA           string `json:"ca" yaml:"ca"`
	Auth         struct {
		Kind string `json:"kind" yaml:"kind"`
	} `json:"auth" yaml:"auth"`
}

// DefaultOptions are the default options for a MongoDB client
var DefaultOptions = Options{
	URI: "mongodb://localhost:27017/",
}

// NewOptionsFromFile reads Options from a config file on a given path
func NewOptionsFromFile(path string) (*Options, error) {
	var opts Options

	err := utils.ReadDataFromFile(path, &opts)
	if err != nil {
		return nil, err
	}

	return &opts, nil
}

// clientOptions transforms Options into an array of ClientOptions for MongoDB's official library
func (o *Options) clientOptions() []*options.ClientOptions {
	clientOptions := make([]*options.ClientOptions, 0)

	if len(o.URI) > 0 {
		clientOptions = append(clientOptions, options.Client().ApplyURI(o.URI))
	}

	if "" != o.CA && "" != o.ClientPEM && "" != o.ClientKeyPEM {
		clientOptions = append(clientOptions, o.getTLSClientOption())
	}

	switch strings.ToLower(o.Auth.Kind) {
	case "x509":
		clientOptions = append(clientOptions, options.Client().SetAuth(
			options.Credential{
				AuthMechanism: "MONGODB-X509",
				AuthSource:    "$extra",
			},
		))
	case "basic":
		clientOptions = append(clientOptions, options.Client().SetAuth(
			options.Credential{
				Username: os.Getenv("BASIC_AUTH_USERNAME"),
				Password: os.Getenv("BASIC_AUTH_PASSWORD"),
			},
		))
	}

	return clientOptions
}

// getTLSClientOption returns a options.ClientOptions from
func (o *Options) getTLSClientOption() *options.ClientOptions {
	if o.CA == "" || o.ClientPEM == "" || o.ClientKeyPEM == "" {
		return nil
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM([]byte(o.CA))

	keyPair, err := tls.X509KeyPair([]byte(o.ClientPEM), []byte(o.ClientKeyPEM))
	if nil != err {
		log.Fatal(err)
	}

	return options.Client().SetTLSConfig(
		&tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{keyPair},
		},
	)
}

// Operation is a Client operation
type Operation func(ctx context.Context, c *Client) error

// Run runs a given operation
func (c *Client) Run(ctx context.Context, op Operation) error {
	return op(ctx, c)
}
