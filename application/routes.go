package application

import (
	"net/http"

	"github.com/go-api/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", loadOrderRoutes)

	return router
}

func loadOrderRoutes(router chi.Router) {
	orderRouter := &handler.Order{}

	router.Post("/", orderRouter.Create)
	router.Get("/", orderRouter.List)
	router.Get("/{id}", orderRouter.GetByID)
	router.Put("/{id}", orderRouter.UpdateByID)
	router.Delete("/{id}", orderRouter.DeleteByID)

}
