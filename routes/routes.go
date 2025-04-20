package routes

import (
	"inventory-management/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	productController := controllers.NewProductController(db)
	inventoryController := controllers.NewInventoryController(db)
	orderController := controllers.NewOrderController(db)

	products := r.Group("/products")
	{
		products.GET("/", productController.GetProducts)
		products.GET("/:id", productController.GetProductByID)
		products.GET("/category/:category", productController.GetProductsByCategory)
		products.POST("/", productController.CreateProduct)
		products.PUT("/:id", productController.UpdateProduct)
		products.DELETE("/:id", productController.DeleteProduct)

		products.POST("/:id/image", productController.UploadImage)
		products.GET("/:id/image", productController.GetImageByProductID)
	}

	inventory := r.Group("/inventory")
	{
		inventory.GET("/:product_id", inventoryController.GetInventoryByProductID)
		inventory.PUT("/:product_id", inventoryController.UpdateStock)
		inventory.DELETE("/:product_id", inventoryController.DeleteInventory)
		inventory.POST("/", inventoryController.CreateInventory)
	}

	order := r.Group("/orders")
	{
		order.POST("/", orderController.CreateOrder)
		order.GET("/:id", orderController.GetOrderByID)
		order.PUT("/:id", orderController.UpdateOrder)
		order.DELETE("/:id", orderController.DeleteOrder)
	}
}
