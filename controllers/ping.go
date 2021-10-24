package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping() func(*gin.Context) {
	type Response struct {
		Message string
	}

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, Response{
			Message: "Pong",
		})
	}
}
