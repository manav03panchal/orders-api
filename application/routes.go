package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/manav03panchal/orders-api.git/handler"
)

// load the routes.
// returns pointer to a new Chi Multiplexer.
func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	// invoke logger for verbosity.
	router.Use(middleware.Logger)
	// set StatusOK for "/"
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// run function loadOrderRoutes to handle all possible CRUD operations.
	router.Route("/orders", loadOrderRoutes)

	return router
}

func loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{}

	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetByID)
	router.Put("/{id}", orderHandler.UpdateByID)
	router.Delete("/{id}", orderHandler.DeleteByID)
}
