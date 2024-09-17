package controller

import (
	"cat_adoption_platform/model"
	"cat_adoption_platform/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type reviewController struct {
	service service.ReviewService
	rg      *gin.RouterGroup
}

func NewReviewController(service service.ReviewService, rg *gin.RouterGroup) *reviewController {
	return &reviewController{
		service: service,
		rg:      rg,
	}
}

func (c *reviewController) Route() {
	router := c.rg.Group("/review")
	router.POST("", c.Create)
	router.GET("/:id", c.GetByID)
	router.PUT("/:id", c.Update)
	router.DELETE("/:id", c.Delete)
	router.GET("", c.GetAll)
}

func (c *reviewController) Create(ctx *gin.Context) {
	var review model.Review
	if err := ctx.ShouldBindJSON(&review); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdReview, err := c.service.Create(review)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review"})
		return
	}

	ctx.JSON(http.StatusCreated, createdReview)
}

func (c *reviewController) GetByID(ctx *gin.Context) {
	reviewId := ctx.Param("review_id")
	review, err := c.service.GetByID(reviewId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	ctx.JSON(http.StatusOK, review)
}

func (c *reviewController) Update(ctx *gin.Context) {
	reviewId := ctx.Param("review_id")
	var updatedReview model.Review
	if err := ctx.ShouldBindJSON(&updatedReview); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedReview.ReviewID = reviewId
	_, err := c.service.Update(updatedReview)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update review"})
		return
	}

	ctx.JSON(http.StatusOK, updatedReview)
}

func (c *reviewController) Delete(ctx *gin.Context) {
	reviewId := ctx.Param("review_id")
	err := c.service.Delete(reviewId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete review"})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *reviewController) GetAll(ctx *gin.Context) {
	reviews, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all reviews"})
		return
	}

	ctx.JSON(http.StatusOK, reviews)
}
