package mongo

import (
	"context"

	"github.com/DarkSoul94/services/models"
	"github.com/DarkSoul94/services/service2/app"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ticketCollection = "tickets"
)

type mongoRepo struct {
	ticketCollection *mongo.Collection
}

func NewMongoRepo(db *mongo.Database) app.Repository {
	return &mongoRepo{
		ticketCollection: db.Collection(ticketCollection),
	}
}

type dbTicket struct {
	ID     string `bson:"_id"`
	UserID string `bson:"user_id"`
	Result bool   `bson:"result"`
	Status int    `bson:"status"`
}

func (r *mongoRepo) toDbTicket(ticket models.Ticket) dbTicket {
	return dbTicket{
		ID:     ticket.ID.String(),
		UserID: ticket.UserID.String(),
		Result: ticket.Result,
		Status: int(ticket.Status),
	}
}

func (r *mongoRepo) InsertResult(ctx context.Context, ticket models.Ticket) error {
	_, err := r.ticketCollection.UpdateByID(ctx, ticket.ID, r.toDbTicket(ticket))
	
	return err
}
