package routes

import (
	"backend/internals/database"
	"backend/internals/handler"
	customMiddlewares "backend/internals/middlewares"
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
	r.Use(middleware.Recoverer)

	// TODO: Add rate limiting middleware

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
	orderHandler := handler.NewOrderHandler(*services.NewOrderService(store.NewOrderStore(db)))

	r.Post("/users/register", userHandler.RegisterUserHandler)
	r.Post("/users/login", userHandler.LoginUserHandler)

	r.Get("/products", productHandler.GetProducts)

	// user only routes
	r.Group(func(r chi.Router) {
		r.Use(customMiddlewares.UserOnlyMiddleware)
		r.Get("/users/{userID}", userHandler.GetUserProfile)

		r.Get("/cart", cartHandler.GetCart)
		r.Post("/cart", cartHandler.AddToCart)
		r.Delete("/cart", cartHandler.ClearCart)
		r.Delete("/cart/{productID}", cartHandler.RemoveFromCart)

		r.Post("/orders", orderHandler.CreateOrder)
		r.Get("/orders", orderHandler.GetOrders)
		r.Get("/orders/{orderID}", orderHandler.GetOrder)
	})

	// admin only routes
	r.Group(func(r chi.Router) {
		r.Use(customMiddlewares.AdminOnlyMiddleware)

		r.Post("/products", productHandler.CreateProduct)
		r.Put("/products", productHandler.UpdateProduct)
		r.Delete("/products/{productID}", productHandler.DeleteProduct)

		r.Put("/orders/{orderID}", orderHandler.UpdateOrderStatus)
	})

	return r
}
