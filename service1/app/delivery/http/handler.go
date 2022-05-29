package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
