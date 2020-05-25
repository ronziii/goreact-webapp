package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ictlife/infra-interview-may-2020/backend/app/db"
	"github.com/ictlife/infra-interview-may-2020/backend/app/services"
	"github.com/ictlife/infra-interview-may-2020/backend/app/web/api/products"
)

type AppRouter struct {
	*gin.Engine
}

func BuildRouter(
	securer gin.HandlerFunc,
	compressor gin.HandlerFunc,
	dB db.DB,
	productService services.ProductService,
) *AppRouter {

	if os.Getenv("ENVIRONMENT") == "development" {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	router.Use(securer)
	router.Use(compressor)
	router.Use(corsMiddleware())

	apiV1Router := router.Group("/api/v1")

	products.AddEndpoints(
		apiV1Router,
		dB,
		productService,
	)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error_message": "Endpoint not found"})
	})

	return &AppRouter{router}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-ICTLIFE-APPLICATION, X-ICTLIFE-TOKEN")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "X-CSRF-Token, Authorization, X-Requested-With, X-ICTLIFE-APPLICATION, X-ICTLIFE-TOKEN")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
