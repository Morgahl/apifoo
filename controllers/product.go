package controllers

import (
	"net/http"
	"strconv"

	"github.com/curlymon/gormfoo/models"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GET_Products(db *gorm.DB) func(*gin.Context) {
	type Response struct {
		Products []models.Product
	}

	type ErrorResponse struct {
		Error string
	}

	return func(c *gin.Context) {
		products, err := models.ProductsAll(db)
		if err != nil {
			log.Debug().
				Err(err).
				Msg("error retrieving products from database")
			c.JSON(http.StatusNotFound, ErrorResponse{
				Error: "products not found",
			})
			return
		}

		c.JSON(http.StatusOK, Response{
			Products: products,
		})
	}
}

func GET_Product(db *gorm.DB) func(*gin.Context) {
	type Response struct {
		Product models.Product
	}

	type ErrorResponse struct {
		Error string
	}

	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Error: "invalid id format",
			})
			return
		}

		product, err := models.ProductByID(db, id)
		if err != nil || product == nil {
			log.Debug().
				Err(err).
				Int("id", id).
				Msg("error retrieving product from database")
			c.JSON(http.StatusNotFound, ErrorResponse{
				Error: "product not found",
			})
			return
		}

		c.JSON(http.StatusOK, Response{
			Product: *product,
		})
	}
}

func POST_Product(db *gorm.DB) func(*gin.Context) {
	type Response struct {
		Product models.Product
	}

	type ErrorResponse struct {
		Error string
	}

	return func(c *gin.Context) {
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			log.Debug().
				Err(err).
				Msg("error parsing product from request")
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Error: "invalid object",
			})
			return
		}

		if err := models.ProductNew(db, &product); err != nil {
			log.Debug().
				Err(err).
				Msg("error creating product in database")
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Error: "unknown error occured",
			})
			return
		}

		c.JSON(http.StatusCreated, Response{
			Product: product,
		})
	}
}
