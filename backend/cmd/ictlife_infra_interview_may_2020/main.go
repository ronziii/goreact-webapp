package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/contrib/secure"
	"github.com/ictlife/infra-interview-may-2020/backend/app/db"
	"github.com/ictlife/infra-interview-may-2020/backend/app/logger"
	"github.com/ictlife/infra-interview-may-2020/backend/app/repos"
	"github.com/ictlife/infra-interview-may-2020/backend/app/services"
	"github.com/ictlife/infra-interview-may-2020/backend/app/web/router"
)

func main() {

	// Initialize database
	appDB := db.InitDB()
	defer appDB.Close()

	// Initialize repos
	productRepository := repos.NewProductRepository()

	productService := services.NewProductService(
		productRepository,
	)

	securer := secure.Secure(secure.Options{
		SSLRedirect:          strings.ToLower(os.Getenv("FORCE_SSL")) == "true",
		SSLProxyHeaders:      map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:           315360000,
		STSIncludeSubdomains: true,
		FrameDeny:            true,
		ContentTypeNosniff:   true,
		BrowserXssFilter:     true,
	})

	compressor := gzip.Gzip(gzip.DefaultCompression)

	appRouter := router.BuildRouter(
		securer,
		compressor,
		appDB,
		productService,
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: appRouter,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		logger.Infof("Process terminated...shutting down")
		if err := server.Close(); err != nil {
			log.Fatalf("Server close: [%v]", err)
		}
	}()

	logger.Infof("Starting server on port %v", port)

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			logger.Infof("Server shut down")
		} else {
			log.Fatalf("Server shut down unexpectedly, err=[%v]!", err)
		}
	}

	logger.Infof("Server exiting")
}
