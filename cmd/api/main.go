package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"techtrack/internal/config"
	"techtrack/internal/database"
	"techtrack/internal/repository/postgres"
	"techtrack/internal/transport/http/middleware"
	v1 "techtrack/internal/transport/http/v1"
)

func main() {
	// 1. Setup Structured Logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// 2. Load Configuration
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		os.Exit(1)
	}

	// 3. Connect to Database
	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	// 4. Setup Router
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	// Core Dependency Injection
	tenantRepo := postgres.NewTenantRepository(db)
	userRepo := postgres.NewUserRepository(db)
	assetRepo := postgres.NewAssetRepository(db)
	ticketRepo := postgres.NewTicketRepository(db)
	auditRepo := postgres.NewAuditLogRepository(db)

	authHandler := v1.NewAuthHandler(userRepo)
	tenantHandler := v1.NewTenantHandler(tenantRepo)
	userHandler := v1.NewUserHandler(userRepo)
	assetHandler := v1.NewAssetHandler(assetRepo)
	ticketHandler := v1.NewTicketHandler(ticketRepo)
	auditHandler := v1.NewAuditHandler(auditRepo)

	// API Routes
	r.Route("/api/v1", func(r chi.Router) {
		// Public Routes
		authHandler.RegisterRoutes(r)

		// Protected Routes
		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)

			tenantHandler.RegisterRoutes(r)
			userHandler.RegisterRoutes(r)
			assetHandler.RegisterRoutes(r)
			ticketHandler.RegisterRoutes(r)
			auditHandler.RegisterRoutes(r)
		})
	})

	// Health Check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// 5. Start Server
	server := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: r,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		slog.Info("Server starting", "port", cfg.HTTPPort, "env", cfg.Environment)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed", "error", err)
			os.Exit(1)
		}
	}()

	<-stop
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
	}

	slog.Info("Server exited properly")
}
