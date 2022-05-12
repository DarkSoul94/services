package http

import (
	"github.com/DarkSoul94/services/service1/app"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.RouterGroup, uc app.Usecase) {
	h := NewHandler(uc)

	router.POST("/ticket", h.NewTicket)
	router.GET("/ticket", h.TicketList)
}
