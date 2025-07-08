package pokeapi

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
var limit int = 20
var offset int = 0

func LocationList(o int) (RespLocations, error) {
	if o < 0 {
		offset += o
		if offset < 0 {
			offset = 0
		}
	}

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

	if o > 0 {
		offset += o
	}

	return locationsResp, nil
}
