package controller

import (
	"cat_adoption_platform/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadImageController struct {
	service service.UploadImageService
	rg      *gin.RouterGroup
}

func NewUploadImageController(service service.UploadImageService, rg *gin.RouterGroup) *UploadImageController {
	return &UploadImageController{service: service, rg: rg}
}

func (c *UploadImageController) UploadImage(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}
	defer file.Close()

	image, err := c.service.Upload(file, header.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"image": image})
}

func (c *UploadImageController) Route() {
	router := c.rg.Group("/upload")
	router.POST("", c.UploadImage)
}
