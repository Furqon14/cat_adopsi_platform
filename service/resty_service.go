package service

import (
	"cat_adoption_platform/config"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// CatResty adalah interface untuk layanan geocoding dan reverse geocoding
type CatResty interface {
	Geocode(address string) (lat, lon, locationName string, err error)
	ReverseGeocode(lat, lon string) (address string, err error)
}

// GeocodingResponse menampung response dari OSM API untuk geocoding
type GeocodingResponse []struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
}

// ReverseGeocodingResponse menampung response dari OSM API untuk reverse geocoding
type ReverseGeocodingResponse struct {
	DisplayName string `json:"display_name"`
}

// CatRestyService implementasi dari CatResty
type CatRestyService struct{}

// Geocode menghubungi OSM API untuk mendapatkan koordinat lokasi dari alamat
func (s *CatRestyService) Geocode(address string) (lat, lon, locationName string, err error) {
	client := resty.New()
	cfg := config.AppConfig
	url := fmt.Sprintf("%s/search", cfg.OSMAPIEndpoint)

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"q":      address,
			"format": "json",
			"limit":  "1",
		}).
		SetHeader("Accept", "application/json").
		Get(url)

	if err != nil {
		return "", "", "", fmt.Errorf("error making request to OSM API: %v", err)
	}

	var results GeocodingResponse
	err = json.Unmarshal(resp.Body(), &results)
	if err != nil {
		return "", "", "", fmt.Errorf("error unmarshalling response: %v", err)
	}

	if len(results) == 0 {
		return "", "", "", fmt.Errorf("no results found for query: %s", address)
	}

	return results[0].Lat, results[0].Lon, results[0].DisplayName, nil
}

// ReverseGeocode menghubungi OSM API untuk mendapatkan alamat dari koordinat
func (s *CatRestyService) ReverseGeocode(lat, lon string) (string, error) {
	client := resty.New()
	url := "https://nominatim.openstreetmap.org/reverse"

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"lat":    lat,
			"lon":    lon,
			"format": "json",
		}).
		SetHeader("Accept", "application/json").
		Get(url)

	if err != nil {
		return "", fmt.Errorf("error making request to OSM API: %v", err)
	}

	var reverseResponse ReverseGeocodingResponse
	err = json.Unmarshal(resp.Body(), &reverseResponse)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response: %v", err)
	}

	return reverseResponse.DisplayName, nil
}

// NewRestyService membuat instance baru dari CatRestyService
func NewRestyService() CatResty {
	return &CatRestyService{}
}
