package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	"github.com/vikas/sharechat-revamp/backend/internal/config"
	"github.com/vikas/sharechat-revamp/backend/internal/handlers"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	var logLevel slog.Level
	switch strings.ToLower(cfg.LogLevel) {
	case "debug":
		logLevel = slog.LevelDebug
	case "warn", "warning":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	_ = logger

	// PostgreSQL — optional; only connect if DB_URL is set.
	var dbPool *pgxpool.Pool
	if cfg.DBURL != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dbPool, err = pgxpool.New(ctx, cfg.DBURL)
		if err != nil {
			log.Fatalf("postgres connect: %v", err)
		}
		defer dbPool.Close()

		if err := dbPool.Ping(ctx); err != nil {
			log.Fatalf("postgres ping: %v", err)
		}
		log.Println("postgres: connected")
	}

	// Redis — optional; only connect if REDIS_URL is set.
	var rdb *redis.Client
	if cfg.RedisURL != "" {
		opts, err := redis.ParseURL(cfg.RedisURL)
		if err != nil {
			log.Fatalf("redis parse url: %v", err)
		}
		rdb = redis.NewClient(opts)
		defer rdb.Close()

		if _, err := rdb.Ping(context.Background()).Result(); err != nil {
			log.Fatalf("redis ping: %v", err)
		}
		log.Println("redis: connected")
	}

	healthHandler := handlers.NewHealthHandler(dbPool, rdb)

	r := chi.NewRouter()
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", healthHandler.Health)
	})

	addr := fmt.Sprintf(":%s", cfg.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("server: listening on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server: %v", err)
		}
	}()

	<-stop
	log.Println("server: shutting down")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("server: forced shutdown: %v", err)
	}
	log.Println("server: stopped")
}
