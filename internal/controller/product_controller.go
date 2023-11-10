package controller

import (
	"encoding/json"
	"go-clean-code/internal/entity"
	"go-clean-code/internal/usecase"
	"net/http"
)

type ProductController struct {
	InsertProductUseCase *usecase.InsertProductUseCase
	GetAllProductUseCase *usecase.GetAllProductUseCase
}

func NewProductController(
	insertProductUseCase *usecase.InsertProductUseCase,
	getAllProductUseCase *usecase.GetAllProductUseCase,
) *ProductController {
	return &ProductController{
		InsertProductUseCase: insertProductUseCase,
		GetAllProductUseCase: getAllProductUseCase,
		// Initialize other use cases
	}
}

func (ctrl *ProductController) InsertProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid request body"})
		return
	}

	ctx := r.Context()

	insertedProduct, err := ctrl.InsertProductUseCase.Execute(ctx, product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertedProduct)
}

func (ctrl *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	products, err := ctrl.GetAllProductUseCase.Execute(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
