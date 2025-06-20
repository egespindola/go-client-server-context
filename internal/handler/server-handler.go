package handler

import (
	"github.com/egespindola/go-client-server-context/internal/service"
	"github.com/gin-gonic/gin"
)

type ServerHandler struct {
}

func NewServerHandler() *ServerHandler {
	return &ServerHandler{}
}

func (h *ServerHandler) GetUsdExchangeRate(c *gin.Context) {
	s := service.NewServerSvc()
	exchangeRate, err := s.GetUsdExchangeRate()
	if err != nil {
		c.JSON(500, gin.H{
			"error":   "Failed to fetch exchange rate",
			"details": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "This endpoint will return the current USD to BRL exchange rate.",
		"status":  "success",
		"data":    exchangeRate,
	})
}
