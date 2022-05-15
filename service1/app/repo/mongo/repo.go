package mongo

import (
	"context"

	"github.com/DarkSoul94/services/models"
	"github.com/DarkSoul94/services/service1/app"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userCollection   = "users"
	ticketCollection = "tickets"
)

type mongoRepo struct {
	userCollection   *mongo.Collection
	ticketCollection *mongo.Collection
}

func NewMongoRepo(db *mongo.Database) app.Repository {
	return &mongoRepo{
		userCollection:   db.Collection(userCollection),
		ticketCollection: db.Collection(ticketCollection),
	}
}

type dbUser struct {
	ID    string `bson:"_id"`
	Email string `bson:"email"`
}

func (r *mongoRepo) toDbUser(user models.User) dbUser {
	return dbUser{
		ID:    user.ID.String(),
		Email: user.Email,
	}
}

func (r *mongoRepo) toModelUser(user dbUser) models.User {
	return models.User{
		ID:    uuid.MustParse(user.ID),
		Email: user.Email,
	}
}

func (r *mongoRepo) CreateUser(ctx context.Context, user models.User) error {
	_, err := r.userCollection.InsertOne(ctx, r.toDbUser(user))

	return err
}

func (r *mongoRepo) CheckEmailExist(ctx context.Context, email string) bool {
	res := r.userCollection.FindOne(ctx, bson.M{"email": email})
	if res.Err() != nil {
		return false
	}

	return true
}

func (r *mongoRepo) GetUserList(ctx context.Context) ([]models.User, error) {
	var (
		dbUsers []dbUser      = make([]dbUser, 0)
		mUsers  []models.User = make([]models.User, 0)
	)

	cur, err := r.userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var user dbUser

		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}

		dbUsers = append(dbUsers, user)
	}

	for _, user := range dbUsers {
		mUsers = append(mUsers, r.toModelUser(user))
	}

	return mUsers, nil
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

func (r *mongoRepo) toModelTicket(ticket dbTicket) models.Ticket {
	return models.Ticket{
		ID:     uuid.MustParse(ticket.ID),
		UserID: uuid.MustParse(ticket.UserID),
		Result: ticket.Result,
		Status: models.TicketStatus(ticket.Status),
	}
}

func (r *mongoRepo) CreateTicket(ctx context.Context, ticket models.Ticket) error {
	_, err := r.ticketCollection.InsertOne(ctx, r.toDbTicket(ticket))
	return err
}

func (r *mongoRepo) GetTicketList(ctx context.Context, userID string) ([]models.Ticket, error) {
	var (
		dbTickets []dbTicket      = make([]dbTicket, 0)
		mTicekts  []models.Ticket = make([]models.Ticket, 0)
	)

	cur, err := r.ticketCollection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var ticket dbTicket

		err := cur.Decode(&ticket)
		if err != nil {
			return nil, err
		}

		dbTickets = append(dbTickets, ticket)
	}

	for _, ticket := range dbTickets {
		mTicekts = append(mTicekts, r.toModelTicket(ticket))
	}

	return mTicekts, nil
}
