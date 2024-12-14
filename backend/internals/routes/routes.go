package routes

import (
	"backend/internals/handler"
	"backend/internals/services"
	"backend/internals/store"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	handler := handler.NewHandler(*services.NewProductService(store.NewMockProductStore()), *services.NewCartService(store.NewMockCartStore()))

	r.Get("/products", handler.SearchProductHandler)
	r.Get("/products/{id}", handler.GetProductDetailHandler)

	return r
}
