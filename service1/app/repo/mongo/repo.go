package mongo

import (
	"context"

	"github.com/DarkSoul94/services/models"
	"github.com/DarkSoul94/services/service1/app"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type dbTciket struct {
	ID     primitive.ObjectID
	UserID string
	Result bool
	Status int
}

func (r *mongoRepo) CreateUser(ctx context.Context, user models.User) error {
	_, err := r.userCollection.InsertOne(ctx, r.toDbUser(user))
	if err != nil {
		return err
	}

	return err
}

func (r mongoRepo) CheckEmailExist(ctx context.Context, email string) bool {
	res := r.userCollection.FindOne(ctx, bson.M{"email": email})
	if res.Err() != nil {
		return false
	}

	return true
}
