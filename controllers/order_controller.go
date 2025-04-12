package controllers

import (
	"inventory-management/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	DB *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{DB: db}
}

func (o *OrderController) CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	if err := o.DB.First(&product, order.ProductID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	order.OrderDate = time.Now()

	if err := o.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat pesanan"})
		return
	}

	if err := o.DB.Preload("Product").First(&order, order.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil pesanan"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": order})
}

func (o *OrderController) GetOrderByID(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	if err := o.DB.Preload("Product").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pesanan tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}
