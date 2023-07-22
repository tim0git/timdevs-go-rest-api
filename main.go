package main

import (
	"context"
	_ "eve.vehicle.api.com/m/v2/docs"
	"eve.vehicle.api.com/m/v2/handler_health"
	"eve.vehicle.api.com/m/v2/handler_register_vehicle"
	"eve.vehicle.api.com/m/v2/handler_retrieve_vehicle"
	"eve.vehicle.api.com/m/v2/handler_update_vehicle"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	router.Use(gzip.Gzip(gzip.BestSpeed))
	router.GET("/health", handler_health.Health)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.POST("/vehicle", handler_register_vehicle.RegisterVehicle)
	router.GET("/vehicle/:vin", handler_retrieve_vehicle.RetrieveVehicle)
	router.PATCH("/vehicle/:vin", handler_update_vehicle.UpdateVehicle)

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
