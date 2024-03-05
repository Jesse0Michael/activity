package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jesse0michael/activity/internal/service"
	"github.com/jesse0michael/go-request"
)

type ActivityHandler interface {
	Activities(ctx context.Context, req service.ActivitiesRequest) ([]service.Activity, error)
	Activity(ctx context.Context, id string) (*service.Activity, error)
	CreateActivity(ctx context.Context, activity service.Activity) error
	UpdateActivity(ctx context.Context, id string, activity service.Activity) error
	PatchActivity(ctx context.Context, id string, activity service.Activity) (*service.Activity, error)
	DeleteActivity(ctx context.Context, id string) (bool, error)
}

type Config struct {
	Port    int           `envconfig:"PORT" default:"8080"`
	Timeout time.Duration `envconfig:"TIMEOUT" default:"30s"`
}

type Server struct {
	*http.Server
}

func New(cfg Config, activityService ActivityHandler) *Server {
	server := &Server{
		Server: &http.Server{
			Handler:     routes(activityService),
			Addr:        fmt.Sprintf(":%d", cfg.Port),
			ReadTimeout: cfg.Timeout,
		},
	}

	return server
}

func decode[T any](r *http.Request, v T) error {
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return fmt.Errorf("failed to decode json: %w", err)
	}
	return nil
}

func decodeRequest[T any](r *http.Request, v T) error {
	if err := request.Decode(r, &v); err != nil {
		return fmt.Errorf("failed to decode request: %w", err)
	}
	return nil
}

func encode[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("failed to encode json: %w", err)
	}
	return nil
}

func encodeError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write([]byte(fmt.Sprintf(`{"error":%q}`, err.Error())))
}
