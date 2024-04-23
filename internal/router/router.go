package router

import (
	"github.com/gin-gonic/gin"
	"github.com/irvankadhafi/go-boilerplate/database"
	"github.com/irvankadhafi/go-boilerplate/pkg/middleware"
)

func Add(ginEngine *gin.Engine, db *database.DbSql) {
	productController := productController(db)
	ginEngine.Use(middleware.Trace())

	ginEngine.POST("/api/products", productController.Create)
	ginEngine.GET("/api/products/:id", productController.GetDetail)
	ginEngine.DELETE("/api/products/:id", productController.Delete)
	ginEngine.PUT("/api/products/:id", productController.Update)
	ginEngine.GET("/api/products", productController.GetList)
}
