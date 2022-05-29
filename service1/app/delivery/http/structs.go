package http

import (
	"github.com/DarkSoul94/services/models"
	"github.com/DarkSoul94/services/service1/app"
	"github.com/google/uuid"
)

type Handler struct {
	uc app.Usecase
}

func NewHandler(uc app.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

type newUser struct {
	Email string `json:"email"`
}

func (h *Handler) toModelUser(user newUser) models.User {
	return models.User{
		Email: user.Email,
	}
}

type hUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (h *Handler) tohUser(user models.User) hUser {
	return hUser{
		ID:    user.ID.String(),
		Email: user.Email,
	}
}

type newTicket struct {
	UserID string `json:"user_id"`
}

func (h *Handler) toModelTicket(ticket newTicket) models.Ticket {
	return models.Ticket{
		UserID: uuid.MustParse(ticket.UserID),
	}
}

type hTicket struct {
	ID     string
	UserID string
	Result bool
	Status string
}

func (h *Handler) tohTicket(ticket models.Ticket) hTicket {
	return hTicket{
		ID:     ticket.ID.String(),
		UserID: ticket.UserID.String(),
		Result: ticket.Result,
		Status: models.TicketStatusString[ticket.Status],
	}
}
