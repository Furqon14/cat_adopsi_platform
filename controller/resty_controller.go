package controller

import (
	"cat_adoption_platform/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CatControllerApi struct {
	service *service.CatResty
	router  *gin.RouterGroup
}

func NewCatControllerApi(service *service.CatResty, router *gin.RouterGroup) *CatControllerApi {
	return &CatControllerApi{
		service: service,
		router:  router,
	}
}

func (c *CatControllerApi) Route() {
	c.router.GET("/test-location", c.GetLocation)
	c.router.GET("/geocode", c.Geocode)
	c.router.GET("/reverse-geocode", c.ReverseGeocode)
}

// GetLocation adalah handler untuk endpoint /test-location
func (cc *CatControllerApi) GetLocation(c *gin.Context) {
	locationQuery := "Jakarta"

	lat, lon, locationName, err := (*cc.service).Geocode(locationQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"latitude":     lat,
		"longitude":    lon,
		"locationName": locationName,
	})
}

// Geocode adalah handler untuk endpoint /geocode
func (cc *CatControllerApi) Geocode(c *gin.Context) {
	address := c.Query("address")

	lat, lon, locationName, err := (*cc.service).Geocode(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"latitude":     lat,
		"longitude":    lon,
		"locationName": locationName,
	})
}

// ReverseGeocode adalah handler untuk endpoint /reverse-geocode
func (cc *CatControllerApi) ReverseGeocode(c *gin.Context) {
	lat := c.Query("lat")
	lon := c.Query("lon")

	address, err := (*cc.service).ReverseGeocode(lat, lon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"address": address,
	})
}
