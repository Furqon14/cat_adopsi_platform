package controller

import (
	"cat_adoption_platform/model"
	"cat_adoption_platform/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CatController struct {
	service service.CatService
	rg      *gin.RouterGroup
}

// GetAllCats mendapatkan daftar semua kucing
func (c *CatController) GetAllCats(ctx *gin.Context) {
	cats, err := c.service.GetAllCats()
	if err != nil {
		fmt.Println("Error getting all cats:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cats"})
		return
	}
	if len(cats) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No cats found"})
		return
	}
	ctx.JSON(http.StatusOK, cats)
}

// GetCatByID mendapatkan detail kucing berdasarkan ID
func (c *CatController) GetCatByID(ctx *gin.Context) {
	id := ctx.Param("id")
	cat, err := c.service.GetCatByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get by id cat"})
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
	var newCat model.Cat
	if err := ctx.ShouldBindJSON(&newCat); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCat, err := c.service.CreateCat(&newCat)
	if err != nil {
		// Tambahkan log error
		fmt.Println("Error creating cat:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cat"})
		return
	}

	ctx.JSON(http.StatusCreated, createdCat)
}

// DeleteCat menghapus kucing berdasarkan ID
func (c *CatController) DeleteCat(ctx *gin.Context) {
	catID := ctx.Param("id")

	err := c.service.DeleteCat(catID)
	if err != nil {
		log.Println("Error deleting cat:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cat"})
		return
	}

	// Berikan respons pesan sukses
	ctx.JSON(http.StatusOK, gin.H{"message": "Cat deleted successfully"})
}

func (u *CatController) Route() {
	router := u.rg.Group("/cats")
	router.GET("", u.GetAllCats)
	router.GET("/:id", u.GetCatByID)
	router.POST("", u.CreateCat)
	router.DELETE("/:id", u.DeleteCat)
}

func NewCatController(service *service.CatService, rg *gin.RouterGroup) *CatController {
	return &CatController{service: *service, rg: rg}
}
