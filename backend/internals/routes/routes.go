package routes

import (
	"backend/internals/database"
	"backend/internals/handler"
	"backend/internals/services"
	"backend/internals/store"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func RegisterRoutes() http.Handler {
	db, err := database.New(
		"postgresql://postgres:admin@localhost:5432/quikshop?sslmode=disable",
		30,
		30,
		"15m",
	)
	if err != nil {
		log.Fatal("error intializing database connection : ", db)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	userHandler := handler.NewUserHandler(services.NewUserService(store.NewUserStore(db)))
	productHandler := handler.NewProductHandler(*services.NewProductService(store.NewProductStore(db)))
	cartHandler := handler.NewCartHandler(*services.NewCartService(store.NewCartStore(db)))

	r.Post("/users/register", userHandler.RegisterUserHandler)
	r.Post("/users/login", userHandler.LoginUserHandler)
	r.Get("/users/{userID}", userHandler.GetUserProfile)

	r.Get("/products", productHandler.GetProducts)
	r.Post("/products", productHandler.CreateProduct)
	r.Put("/products", productHandler.UpdateProduct)
	r.Delete("/products/{productID}", productHandler.DeleteProduct)

	r.Get("/cart", cartHandler.GetCart)
	r.Post("/cart", cartHandler.AddToCart)
	r.Delete("/cart", cartHandler.ClearCart)
	r.Delete("/cart/{productID}", cartHandler.RemoveFromCart)

	return r
}
