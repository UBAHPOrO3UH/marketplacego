package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"marketplace/internal/models"
	"marketplace/internal/repo"
)

type ProductHandler struct {
	Repo *repo.ProductRepo
}

func NewProductHandler(repo *repo.ProductRepo) *ProductHandler {
	return &ProductHandler{Repo: repo}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Repo.Create(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	var products []models.Product
	if err := h.Repo.GetAll(&products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	var product models.Product
	if err := h.Repo.GetById(&product, c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}
