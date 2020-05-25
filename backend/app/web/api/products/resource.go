package products

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ictlife/infra-interview-may-2020/backend/app/db"
	"github.com/ictlife/infra-interview-may-2020/backend/app/forms"
	"github.com/ictlife/infra-interview-may-2020/backend/app/services"
)

func createProduct(dB db.DB, productService services.ProductService) func(*gin.Context) {

	return func(c *gin.Context) {

		var form forms.ProductForm
		err := c.BindJSON(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": fmt.Sprintf("You provided an invalid form: %v", err)})
			return
		}

		product, err := productService.CreateProduct(c.Request.Context(), dB, &form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, product)
	}
}

func getProduct(dB db.DB, productService services.ProductService) func(*gin.Context) {

	return func(c *gin.Context) {

		productID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": fmt.Sprintf("invalid id: %v", c.Param("id"))})
			return
		}

		product, err := productService.GetProduct(c.Request.Context(), dB, productID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

func listProducts(dB db.DB, productService services.ProductService) func(*gin.Context) {

	return func(c *gin.Context) {

		productList, err := productService.ListProducts(c.Request.Context(), dB)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, productList)
	}
}

func updateProduct(dB db.DB, productService services.ProductService) func(*gin.Context) {

	return func(c *gin.Context) {

		productID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": fmt.Sprintf("invalid id: %v", c.Param("id"))})
			return
		}

		var form forms.ProductForm
		err = c.BindJSON(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": fmt.Sprintf("You provided an invalid form: %v", err)})
			return
		}

		product, err := productService.UpdateProduct(c.Request.Context(), dB, productID, &form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

func deleteProduct(dB db.DB, productService services.ProductService) func(*gin.Context) {

	return func(c *gin.Context) {

		productID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": fmt.Sprintf("invalid id: %v", c.Param("id"))})
			return
		}

		product, err := productService.DeleteProduct(c.Request.Context(), dB, productID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}
