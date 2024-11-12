package main

import (
	"github.com/bsorawit1234/e-commerce-api/controllers"
	"github.com/bsorawit1234/e-commerce-api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	router := gin.Default()

	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:id", controllers.GetProductByID)
	router.POST("/products", controllers.CreateProduct)
	router.PUT("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrders)
	router.GET("/orders/:id", controllers.GetOrderByID)

	// Start the server
	router.Run(":8080")
}
