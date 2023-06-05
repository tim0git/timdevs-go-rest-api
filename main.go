package main

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	_ "timdevs.rest.api.com/m/v2/docs"
	"timdevs.rest.api.com/m/v2/handlers"
	"time"
)

var port = os.Getenv("PORT")

// @title Vehicle API
// @description This is the eVe API for vehicle management
// @version 1
// @host localhost:8443
// @BasePath /
// @schemes http
// @produce json
// @consumes json

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	router.GET("/health", handlers.Health)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.POST("/vehicle", handlers.RegisterVehicle)
	router.GET("/vehicle/:vin", handlers.RetrieveVehicle)
	router.PATCH("/vehicle/:vin", handlers.UpdateVehicle)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("Server shutting down gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exited gracefully")
}
