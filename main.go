package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"sample/database"
	"sample/model"
	"sample/router"
	"sample/worker"
)

func main() {
	// Load config
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	port := viper.GetInt("server.port")
	dsn := viper.GetString("database.dsn")

	// Init DB
	database.ConnectDB(dsn)

	// Auto Migrate
	if err := database.DB.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Failed to migrate DB: %v", err)
	}

	// Start Worker Pool
	pool := worker.NewPool(3)
	pool.Start()

	// Feed background jobs
	go func() {
		for i := 1; i <= 20; i++ {
			pool.Jobs <- worker.Job{ID: i, Data: fmt.Sprintf("Payload-%d", i)}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Gin router
	r := router.InitRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutdown signal received")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}

		pool.Stop()

		log.Println("Server exiting")
	}()

	log.Printf("Server running at http://localhost:%d\n", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Listen error: %v", err)
	}
}
