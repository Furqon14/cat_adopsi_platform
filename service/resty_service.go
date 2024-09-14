package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GeocodingResponse []struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
}

type ReverseGeocodingResponse struct {
	DisplayName string `json:"display_name"`
}

// Geocode converts an address to latitude and longitude.
func Geocode(address string) (lat, lon, locationName string, err error) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1", address)
	resp, err := http.Get(url)
	if err != nil {
		return "", "", "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", err
	}

	var geocodingResponse GeocodingResponse
	err = json.Unmarshal(body, &geocodingResponse)
	if err != nil {
		return "", "", "", err
	}

	if len(geocodingResponse) > 0 {
		return geocodingResponse[0].Lat, geocodingResponse[0].Lon, geocodingResponse[0].DisplayName, nil
	}
	return "", "", "", fmt.Errorf("location not found")
}

// ReverseGeocode converts latitude and longitude to an address.
func ReverseGeocode(lat, lon string) (address string, err error) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?lat=%s&lon=%s&format=json", lat, lon)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var reverseResponse ReverseGeocodingResponse
	err = json.Unmarshal(body, &reverseResponse)
	if err != nil {
		return "", err
	}

	return reverseResponse.DisplayName, nil
}
