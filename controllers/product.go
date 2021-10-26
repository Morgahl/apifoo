package controllers

import (
	"net/http"

	"github.com/curlymon/gormfoo/models"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Product struct{}

func (Product) List(c *gin.Context) {
	type Response struct {
		Products []models.Product
	}

	products, err := models.ProductList()
	if err != nil {
		log.Debug().Err(err).Msg("error retrieving products from database")
		c.JSON(http.StatusNotFound, Error("products not found"))
		return
	}

	c.JSON(http.StatusOK, Response{
		Products: products,
	})
}

func (Product) Show(c *gin.Context) {
	type Response struct {
		Product models.Product
	}

	id := c.Param("id")
	product, err := models.ProductByID(id)
	if err != nil {
		log.Debug().Err(err).Str("id", id).Msg("error retrieving product from database")
		c.JSON(http.StatusNotFound, Error("product not found"))
		return
	}

	c.JSON(http.StatusOK, Response{
		Product: product,
	})
}

func (Product) Create(c *gin.Context) {
	type Response struct {
		Product models.Product
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		log.Debug().Err(err).Msg("error parsing product from request")
		c.JSON(http.StatusBadRequest, Error("invalid object"))
		return
	}

	if err := models.ProductCreate(product); err != nil {
		log.Debug().Err(err).Msg("error creating product in database")
		c.JSON(http.StatusBadRequest, Error("unknown error occured"))
		return
	}

	c.JSON(http.StatusCreated, Response{
		Product: product,
	})
}
