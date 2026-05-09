package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/uudon/API-Service-Hub/internal/config"
	"github.com/uudon/API-Service-Hub/internal/handler"
	_ "github.com/uudon/API-Service-Hub/docs" // Import docs for swagger
	"github.com/uudon/API-Service-Hub/internal/middleware"
	"github.com/uudon/API-Service-Hub/pkg/logger"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/rs/zerolog"
)

// @title           API Service Hub
// @version         1.0
// @description     API Service Hub is a Go-based microservice with health check endpoints.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @tag.name health
// @tag.description Health check operations
func main() {
	// Initialize logger
	log := logger.New()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	// Set log level based on config
	zerolog.SetGlobalLevel(cfg.LogLevel)

	// Create Gin router
	router := setupRouter(cfg, log)

	// Create server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
	}

	// Start server in goroutine
	go func() {
		log.Info().Msgf("Starting server on port %d", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exited")
}

func setupRouter(cfg *config.Config, log *zerolog.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/healthz", middleware.Logger(log, handler.HealthCheck))

	// Swagger documentation
	mux.HandleFunc("/swagger/*", httpSwagger.WrapHandler)

	return mux
}
