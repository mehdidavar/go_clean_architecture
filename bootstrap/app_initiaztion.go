package bootstrap

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go-clean-code/internal/controller"
	"go-clean-code/internal/db"
	"go-clean-code/internal/repository"
	"go-clean-code/internal/usecase"
	"net/http"
)

func InitApp() {

	db, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	// Initialize dependencies
	productRepository := repository.NewProductRepository(db)
	insertProductUseCase := usecase.NewInsertProductUseCase(productRepository)
	getAllProductUseCase := usecase.NewGetAllProductUseCase(productRepository)

	// Create the ProductController and pass the required use cases
	productController := controller.NewProductController(
		insertProductUseCase,
		getAllProductUseCase,
		// Pass other use cases as needed
	)

	// Initialize Chi router
	r := chi.NewRouter()

	// Middleware setup (optional, but useful for logging, etc.)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Post("/product", productController.InsertProduct)
	r.Get("/products", productController.GetAllProducts)

	// Start the server
	http.ListenAndServe(":8080", r)
}
