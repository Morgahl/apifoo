package main

import (
	"github.com/curlymon/gormfoo/controllers"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()

	router.Use(
		logger.SetLogger(),
		gin.Recovery(),
	)

	router.GET("/ping", controllers.Ping)

	{
		g := router.Group("/product")
		p := controllers.Product{}
		g.GET("/", p.List)
		g.GET("/:id", p.Show)
		g.POST("/", p.Create)
	}

	return router
}
