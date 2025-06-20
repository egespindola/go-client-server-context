package main

import (
	"log"

	"github.com/egespindola/go-client-server-context/internal/handler"
	"github.com/gin-gonic/gin"
)

var wsPort string = "8080"

func main() {
	h := handler.NewServerHandler()

	r := gin.Default()
	r.GET("/cotacao", h.GetUsdExchangeRate)

	if err := r.Run(":" + wsPort); err != nil {
		panic("Server failed to start: " + err.Error())
	}
	log.Println("Server is running on port " + wsPort)

}
