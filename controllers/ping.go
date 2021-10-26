package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	type Response struct {
		Message string
	}

	c.JSON(http.StatusOK, Response{
		Message: "Pong",
	})
}
