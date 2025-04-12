package controllers

import (
	"inventory-management/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InventoryController struct {
	DB *gorm.DB
}

func NewInventoryController(db *gorm.DB) *InventoryController {
	return &InventoryController{DB: db}
}

func (c *InventoryController) GetInventoryByProductID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("product_id"))
	var inventory models.Inventory

	if err := c.DB.Where("product_id = ?", id).First(&inventory).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": inventory})
}

func (c *InventoryController) UpdateStock(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("product_id"))
	var inventory models.Inventory

	if err := c.DB.Where("product_id = ?", id).First(&inventory).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}

	var input struct {
		Quantity int `json:"quantity"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	inventory.Quantity = input.Quantity

	if err := c.DB.Save(&inventory).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stock"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": inventory})
}

func (c *InventoryController) CreateInventory(ctx *gin.Context) {
	var input struct {
		ProductID uint   `json:"product_id"`
		Quantity  int    `json:"quantity"`
		Location  string `json:"location"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	inventory := models.Inventory{
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Location:  input.Location,
	}

	if err := c.DB.Create(&inventory).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create inventory"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": inventory})
}
