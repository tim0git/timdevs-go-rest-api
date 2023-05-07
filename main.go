package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"timdevs.rest.api.com/m/v2/controllers"
	"time"
)

var port = os.Getenv("PORT")

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	router.GET("/health", controllers.Health)
	router.POST("/vehicle", controllers.RegisterVehicle)

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

// POST /vehicle
// GET /vehicle?vin=1234567890
// PATCH /vehicle?vin=1234567890
// DELETE /vehicle?vin=1234567890

// POST vehicle/user
// GET vehicle/user?user_id=1234567890
// PATCH vehicle/user?user_id=1234567890
// DELETE vehicle/user?user_id=1234567890
