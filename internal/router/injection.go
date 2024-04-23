package router

import (
	"github.com/irvankadhafi/go-boilerplate/database"
	"github.com/irvankadhafi/go-boilerplate/internal/controller/http"
	"github.com/irvankadhafi/go-boilerplate/internal/repository"
	"github.com/irvankadhafi/go-boilerplate/internal/service"
)

func productController(db *database.DbSql) *http.ProductController {
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := http.NewProductController(productService)
	return productController
}
