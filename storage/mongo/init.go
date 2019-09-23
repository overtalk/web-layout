package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"

	m "web-layout/utils/mongo"
)

// MongoDB mongodb driver
type MongoDB struct {
	config m.Cfg
	conn   *mongo.Database
}

// NewMongoDB constructor of MongoDB
func NewMongoDB(config m.Cfg) (*MongoDB, error) {
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		config: config,
		conn:   db,
	}, nil
}
