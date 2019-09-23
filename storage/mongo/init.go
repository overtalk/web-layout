package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"

	m "web-layout/utils/mongo"
)

// Mongo mongodb driver
type Mongo struct {
	config m.Cfg
	conn   *mongo.Database
}

// NewMongo constructor of Mongo DB
func NewMongo(config m.Cfg) (*Mongo, error) {
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}

	return &Mongo{
		config: config,
		conn:   db,
	}, nil
}
