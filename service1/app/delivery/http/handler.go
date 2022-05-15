package http

import (
	"context"
	"net/http"

	"github.com/DarkSoul94/services/models"
	"github.com/DarkSoul94/services/service1/app"
	"github.com/gin-gonic/gin"
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

func (h *Handler) SignUp(c *gin.Context) {
	var newUser newUser

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	id, err := h.uc.Registration(context.Background(), h.toModelUser(newUser))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "success", "id": id})
}

func (h *Handler) GetUserList(c *gin.Context) {
	mUsers, err := h.uc.GetUserList(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	var hUsers []hUser = make([]hUser, 0)
	for _, user := range mUsers {
		hUsers = append(hUsers, h.tohUser(user))
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "success", "users": hUsers})
}

func (h *Handler) NewTicket(c *gin.Context) {
	var newTicket newTicket

	if err := c.BindJSON(&newTicket); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	id, err := h.uc.AcceptNewTicket(context.Background(), h.toModelTicket(newTicket))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "success", "id": id})
}

func (h *Handler) GetTicketList(c *gin.Context) {
	userID := c.Query("id")

	mTickets, err := h.uc.GetTicketList(context.Background(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	var hTickets []hTicket = make([]hTicket, 0)
	for _, ticket := range mTickets {
		hTickets = append(hTickets, h.tohTicket(ticket))
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "success", "tickets": hTickets})
}
