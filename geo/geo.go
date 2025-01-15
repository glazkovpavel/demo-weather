package geo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulationResponse struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := CheckCity(city)
		if !isCity {
			panic("Такого города нет")
		}
		return &GeoData{
			City: city,
		}, nil
	}
	res, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code %d", res)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}

func CheckCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{"city": city})
	res, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return false
	}
	var response CityPopulationResponse
	json.Unmarshal(body, &response)
	return !response.Error

}
