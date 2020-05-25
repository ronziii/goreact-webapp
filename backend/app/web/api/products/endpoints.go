package products

import (
	"github.com/gin-gonic/gin"
	"github.com/ictlife/infra-interview-may-2020/backend/app/db"
	"github.com/ictlife/infra-interview-may-2020/backend/app/services"
)

func AddEndpoints(
	router *gin.RouterGroup,
	dB db.DB,
	productService services.ProductService,
) {

	router.POST("/products", createProduct(dB, productService))

	router.GET("/products/:id", getProduct(dB, productService))

	router.GET("/products", listProducts(dB, productService))

	router.PUT("/products/:id", updateProduct(dB, productService))

	router.DELETE("/products/:id", deleteProduct(dB, productService))
}
