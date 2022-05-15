package http

import (
	"context"
	"net/http"

	"github.com/DarkSoul94/services/models"
	"github.com/DarkSoul94/services/service1/app"
	"github.com/gin-gonic/gin"
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

type hTicket struct {
	Email  string `json:"email"`
	UserID string `json:"user_id"`
}

func (h *Handler) toModelTicket(t hTicket) models.Ticket {
	return models.Ticket{
		Email:  t.Email,
		UserID: t.UserID,
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

func (h *Handler) NewTicket(c *gin.Context) {
	var newTicket hTicket

	if err := c.BindJSON(&newTicket); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	id, err := h.uc.AcceptNewTicket(h.toModelTicket(newTicket))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "success", "id": id})
}

func (h *Handler) TicketList(c *gin.Context) {

}
