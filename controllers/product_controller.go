package controllers

import (
	"fmt"
	"inventory-management/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{DB: db}
}

func (c *ProductController) GetProducts(ctx *gin.Context) {
	var products []models.Product
	c.DB.Find(&products)

	ctx.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func (c *ProductController) GetProductByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var product models.Product

	if err := c.DB.First(&product, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": product})
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Body",
		})
		return
	}

	if err := c.DB.Create(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create product",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": product,
	})
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var product models.Product

	if err := c.DB.First(&product, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := c.DB.Save(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": product})
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var product models.Product

	if err := c.DB.First(&product, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.DB.Delete(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func (c *ProductController) GetProductsByCategory(ctx *gin.Context) {
	category := ctx.Param("category")
	var products []models.Product

	if err := c.DB.Where("category = ?", category).Find(&products).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products by category"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": products})
}

func (c *ProductController) UploadImage(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var product models.Product

	if err := c.DB.First(&product, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload image"})
		return
	}

	path := fmt.Sprintf("uploads/%d_%s", product.ID, file.Filename)
	if err := ctx.SaveUploadedFile(file, path); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save image"})
		return
	}

	product.ImageURL = "/" + path
	if err := c.DB.Save(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product with image URL"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "image_url": product.ImageURL})
}

func (c *ProductController) GetImageByProductID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var product models.Product

	if err := c.DB.First(&product, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if product.ImageURL != "" {
		ctx.JSON(http.StatusOK, gin.H{"image_url": product.ImageURL})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No image found for this product"})
	}
}
