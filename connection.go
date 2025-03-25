package mgm

import (
	"context"
	"errors"
	"time"

	"github.com/kamva/mgm/v3/internal/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config *Config
var client *mongo.Client
var db *mongo.Database

// Config struct contains extra configuration properties for the mgm package.
type Config struct {
	// Set to 10 second (10*time.Second) for example.
	CtxTimeout time.Duration
}

// NewCtx returns a new Context that will be canceled automatically after the
// specified timeout duration. The caller must call the returned cancel function
// to release associated resources once the context is no longer needed.
func NewCtx(timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return ctx, cancel
}

// Ctx function creates and returns a new context with a default timeout value.
func Ctx() (context.Context, context.CancelFunc) {
	return ctx()
}

func ctx() (context.Context, context.CancelFunc) {
	return NewCtx(config.CtxTimeout)
}

// NewClient returns a new mongodb client.
func NewClient(opts ...*options.ClientOptions) (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewCollection returns a new collection with the supplied database.
func NewCollection(db *mongo.Database, name string, opts ...*options.CollectionOptions) *Collection {
	coll := db.Collection(name, opts...)

	return &Collection{Collection: coll}
}

// ResetDefaultConfig resets the configuration values, client and database.
func ResetDefaultConfig() {
	config = nil
	client = nil
	db = nil
}

// SetDefaultConfig initializes the client and database using the specified configuration values, or default.
func SetDefaultConfig(conf *Config, dbName string, opts ...*options.ClientOptions) (err error) {

	// Use the predefined configuration values as default if the user
	// does not provide any.
	if conf == nil {
		conf = defaultConf()
	}

	config = conf

	if client, err = NewClient(opts...); err != nil {
		return err
	}

	db = client.Database(dbName)

	return nil
}

// CollectionByName returns a new collection using the current configuration values.
func CollectionByName(name string, opts ...*options.CollectionOptions) *Collection {
	return NewCollection(db, name, opts...)
}

// DefaultConfigs returns the current configuration values, client and database.
func DefaultConfigs() (*Config, *mongo.Client, *mongo.Database, error) {
	if util.AnyNil(config, client, db) {
		return nil, nil, nil, errors.New("please setup default config before acquiring it")
	}

	return config, client, db, nil
}

// defaultConf are the default configuration values when none are provided
// to the `SetDefaultConfig` method.
func defaultConf() *Config {
	return &Config{CtxTimeout: 10 * time.Second}
}
