// @title Employee Service API
// @version 1.0
// @description This is a simple service for managing employees
// @host localhost:5000
// @BasePath /

package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log/slog"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/sama-kun/ai-plus-test/docs"
	"github.com/sama-kun/ai-plus-test/internal/config"
	"github.com/sama-kun/ai-plus-test/internal/handler"
	"github.com/sama-kun/ai-plus-test/internal/lib/logger/sl"
	"github.com/sama-kun/ai-plus-test/internal/repository"
	"github.com/sama-kun/ai-plus-test/internal/service"
	"github.com/sama-kun/ai-plus-test/internal/storage"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	cfg := config.MustLoad()

	db, err := storage.NewPostgresDB(cfg.Database)
	if err != nil {
		slog.Error("Failed to init DB", sl.Err(err))
		os.Exit(1)
	}
	defer db.Close()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// r.Use(middleware.URLFormat)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},                     
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}, 
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},   
		ExposedHeaders:   []string{"Link"},                                      
		AllowCredentials: true,                                                  
		MaxAge:           300,
	}))

	repo := repository.NewPostgresEmployeeRepo(db)
	svc := service.NewEmployeeService(repo)
	h := handler.NewEmployeeHandler(svc)

	r.Post("/employee", h.CreateEmployee)
	r.Get("/employee", h.GetEmployee)
	r.Get("/swagger/*", httpSwagger.WrapHandler)



	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		slog.Info("Starting server...", slog.String("address", cfg.HTTPServer.Address))
		slog.Info("Swagger address", slog.String("address",`http://localhost:5000/swagger/index.html`))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed to start", sl.Err(err))
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	slog.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", sl.Err(err))
	}

	slog.Info("Server exited successfully")
}
