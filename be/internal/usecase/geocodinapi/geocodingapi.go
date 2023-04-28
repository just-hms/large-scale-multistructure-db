package geocodinapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/env"
	"github.com/tidwall/gjson"
)

type GeocodingAPI struct {
	apikey string
}

type GeocodingInfo struct {
	Country   string
	Region    string
	City      string
	Latitude  float64
	Longitude float64
	Address   string
}

func New() (*GeocodingAPI, error) {
	apikey, err := env.GetString("GEOCODE_API_SECRET")
	if err != nil {
		return nil, err
	}
	return &GeocodingAPI{
		apikey: apikey,
	}, nil
}

// TODO:
//   - return a struct
func (a *GeocodingAPI) Search(area string) ([]GeocodingInfo, error) {

	// Create a new GET request
	url := fmt.Sprintf("https://api.geoapify.com/v1/geocode/search?text=%s&apiKey=%s", area, a.apikey)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// extract only the necessary value from the request
	// https: //github.com/tidwall/gjson/blob/master/SYNTAX.md#multipaths
	res := gjson.GetBytes(
		body,
		"features.#.properties."+
			"{"+
			"city:city,"+
			"region:state,"+
			"country:country,"+
			"latitude:lat,"+
			"longitude:lon,"+
			"address:formatted,"+
			"}",
	)

	geocods := []GeocodingInfo{}
	err = json.Unmarshal([]byte(res.Raw), &geocods)
	if err != nil {
		return nil, err
	}

	return geocods, nil
}
