package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RespLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func locationListAPI() (RespLocations, error) {
	limit := 20
	offset := 0
	apiAddress := fmt.Sprintf("https://pokeapi.co/api/v2/location/?limit=%v&offset=%v", limit, offset)

	req, err := http.NewRequest("GET", apiAddress, nil)
	if err != nil {
		return RespLocations{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocations{}, err
	}

	locationsResp := RespLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespLocations{}, err
	}

	return locationsResp, nil
}
