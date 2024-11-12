package controllers

import (
	"net/http"

	"github.com/bsorawit1234/e-commerce-api/database"
	"github.com/bsorawit1234/e-commerce-api/models"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var input struct {
		Items []struct {
			ProductID uint `json:"product_id"`
			Quantity  uint `json:"quantity"`
		} `json:"items"`
	}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var orderItems []models.OrderItem
	var totalAmount uint

	for _, item := range input.Items {
		var product models.Product
		err := database.DB.First(&product, item.ProductID).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
			return
		}

		if product.Stock < item.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock for product: " + product.Name})
			return
		}

		product.Stock -= item.Quantity
		database.DB.Save(&product)

		orderItem := models.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Product:   product,
		}

		orderItems = append(orderItems, orderItem)
		totalAmount += product.Price * item.Quantity
	}

	order := models.Order{
		OrderItems:  orderItems,
		TotalAmount: totalAmount,
	}
	database.DB.Create(&order)
	c.JSON(http.StatusCreated, order)
}

func GetOrders(c *gin.Context) {
	var orders []models.Order
	database.DB.Preload("OrderItems.Product").Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	err := database.DB.Preload("OrderItems.Product").First(&order, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}
