package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func RunServer() {
	server := gin.Default()

	server.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET,POST,PUT,DELETE,OPTIONS",
		RequestHeaders: "Origin, X-Requested-With, Content-Type, Accept, Authorization, Access-Control-Allow-Origin ,Access-Control-Request-Method, Access-Control-Request-Headers",
		MaxAge:         50 * time.Second,
	}))

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		},
		)
	})

	server.Run(":3006")
}
