package mongo

import (
	"github.com/DarkSoul94/services/models"
	"github.com/google/uuid"
)

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
