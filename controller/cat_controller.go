package controller

import (
	"cat_adoption_platform/model"
	"cat_adoption_platform/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CatController struct {
	service service.CatService
}

// GetAllCats mendapatkan daftar semua kucing
func (c *CatController) GetAllCats(ctx *gin.Context) {
	cats, err := c.service.GetAllCats()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cats)
}

// GetCatByID mendapatkan detail kucing berdasarkan ID
func (c *CatController) GetCatByID(ctx *gin.Context) {
	id := ctx.Param("id")
	cat, err := c.service.GetCatByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if cat == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Cat not found"})
		return
	}
	ctx.JSON(http.StatusOK, cat)
}

// CreateCat menambahkan kucing baru
func (c *CatController) CreateCat(ctx *gin.Context) {
	var cat model.Cat
	if err := ctx.ShouldBindJSON(&cat); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreateCat(&cat); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, cat)
}

// UpdateCat memperbarui informasi kucing
func (c *CatController) UpdateCat(ctx *gin.Context) {
	id := ctx.Param("id")
	var cat model.Cat
	if err := ctx.ShouldBindJSON(&cat); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cat.CatID = id
	if err := c.service.UpdateCat(&cat); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cat)
}

// DeleteCat menghapus kucing berdasarkan ID
func (c *CatController) DeleteCat(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteCat(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func NewCatController(service *service.CatService) *CatController {
	return &CatController{service: service}
}
