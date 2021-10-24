package main

import (
	"github.com/curlymon/gormfoo/controllers"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *gin.Engine {
	router := gin.New()

	router.Use(
		logger.SetLogger(),
		gin.Recovery(),
	)

	router.GET("/ping", controllers.Ping())

	router.GET("/product", controllers.GET_Products(db))
	router.GET("/product/:id", controllers.GET_Product(db))
	router.POST("/product", controllers.POST_Product(db))

	return router
}
