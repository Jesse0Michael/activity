package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/jesse0michael/activity/internal/server"
	"github.com/jesse0michael/pkg/logger"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	// Setup application
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	ctx, cancel := context.WithCancelCause(ctx)

	logger.NewLogger()
	var cfg server.Config
	if err := envconfig.Process("", &cfg); err != nil {
		cancel(fmt.Errorf("failed to process config"))
	}

	// Initialize dependencies

	// Start application
	srvr := server.New(cfg)
	go func() {
		if err := srvr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			cancel(err)
		}
	}()
	slog.With("port", cfg.Port).Info("server started")

	// Wait for signal
	<-ctx.Done()

	// Exit gracefully
	if err := srvr.Shutdown(context.Background()); err != nil {
		slog.With("error", err).Error("failed to shut down server")
	}

	if ctx.Err() != context.Canceled || ctx.Err() != context.DeadlineExceeded {
		slog.With("error", ctx.Err()).Error("application terminated unexpectedly")
	}

	slog.Info("exiting")
}
