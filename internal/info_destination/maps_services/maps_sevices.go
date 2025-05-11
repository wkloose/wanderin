package maps_services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type MapsService struct{}

// **Mengambil informasi lokasi berdasarkan alamat**
func (s *MapsService) FetchLocation(address string) (string, error) {
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY") // Ambil API Key dari ENV
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", address, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch location")
	}
	defer resp.Body.Close()

	var result struct {
		Results []struct {
			FormattedAddress string `json:"formatted_address"`
		} `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response")
	}

	if len(result.Results) == 0 {
		return "", fmt.Errorf("location not found")
	}

	return result.Results[0].FormattedAddress, nil
}

// **Mengambil rekomendasi wisata/kuliner berdasarkan lokasi**
func (s *MapsService) FetchNearbyPlaces(lat, lon string) ([]map[string]interface{}, error) {
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY") // Ambil API Key dari ENV
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%s,%s&radius=5000&type=restaurant&key=%s", lat, lon, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch places")
	}
	defer resp.Body.Close()

	var result struct {
		Results []map[string]interface{} `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response")
	}

	return result.Results, nil
}