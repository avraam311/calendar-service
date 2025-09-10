package main

import (
	"context"
	"errors"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	eventHandler "github.com/avraam311/calendar-service/internal/api/handlers/event"
	"github.com/avraam311/calendar-service/internal/api/server"
	"github.com/avraam311/calendar-service/internal/config"
	"github.com/avraam311/calendar-service/internal/pkg/logger"
	"github.com/avraam311/calendar-service/internal/pkg/validator"
	eventRepo "github.com/avraam311/calendar-service/internal/repository/event"
	eventService "github.com/avraam311/calendar-service/internal/service/event"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Logger.Env, cfg.Logger.LogFilePath)
	val := validator.New()

	dbpool, err := pgxpool.New(ctx, cfg.DatabaseURL())
	if err != nil {
		log.Fatal("error creating connection pool", zap.Error(err))
	}

	eventR := eventRepo.New(dbpool)
	eventS := eventService.New(eventR)
	eventPostH := eventHandler.NewPostHandler(log, val, eventS)
	eventGetH := eventHandler.NewGetHandler(log, val, eventS)
	r := server.NewRouter(eventPostH, eventGetH)
	s := server.NewServer(cfg.Server.HTTPPort, r)

	go func() {
		log.Info("starting HTTP server", zap.String("port", cfg.Server.HTTPPort))
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("server failed", zap.Error(err))
		}
	}()

	<-ctx.Done()
	log.Info("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("shutting down HTTP server...")
	if err = s.Shutdown(shutdownCtx); err != nil {
		log.Error("could not shutdown HTTP server", zap.Error(err))
	}

	if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
		log.Fatal("timeout exceeded, forcing shutdown")
	}

	log.Info("closing database pool...")
	dbpool.Close()
}
