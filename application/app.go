package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler

	rdb *redis.Client
}

func New() *App {
	app := &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	pingerr := a.rdb.Ping(ctx).Err()
	if pingerr != nil {
		return fmt.Errorf("Couldn't ping redis. %w", pingerr)
	}

	las := server.ListenAndServe()
	if las != nil {
		return fmt.Errorf("failed to start server: %w", las)
	}

	return nil
}
