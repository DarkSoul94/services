package mongo

import (
	"github.com/DarkSoul94/services/service1/app"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ticketCollection = "tickets"
)

type mongoRepo struct {
	db *mongo.Database
}

type dbTciket struct {
	ID     string
	UserID string
	Result bool
}

func NewMongoRepo(db *mongo.Database) app.Repository {
	return &mongoRepo{
		db: db,
	}
}
