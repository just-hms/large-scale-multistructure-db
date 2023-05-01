package webapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/env"
	"github.com/tidwall/gjson"
)

type GeocodingWebAPI struct {
	apikey string
}

func NewGeocodingWebAPI() (*GeocodingWebAPI, error) {
	apikey, err := env.GetString("GEOCODE_API_SECRET")
	if err != nil {
		return nil, err
	}
	return &GeocodingWebAPI{
		apikey: apikey,
	}, nil
}

func (a *GeocodingWebAPI) Search(ctx context.Context, area string) ([]entity.GeocodingInfo, error) {

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

	geocods := []entity.GeocodingInfo{}
	err = json.Unmarshal([]byte(res.Raw), &geocods)
	if err != nil {
		return nil, err
	}

	return geocods, nil
}
