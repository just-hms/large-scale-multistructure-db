package geocodinapi

import (
	"fmt"
	"io"
	"net/http"
)

type GeocodingAPI struct {
	apikey string
}

func New(apikey string) *GeocodingAPI {
	return &GeocodingAPI{
		apikey: apikey,
	}
}

// TODO:
//   - return a struct
func (a *GeocodingAPI) Search(area string) (string, error) {

	// Create a new GET request
	url := fmt.Sprintf("https://api.geoapify.com/v1/geocode/search?text=%s&apiKey=%s", area, a.apikey)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Print the response body
	return string(body), nil
}
