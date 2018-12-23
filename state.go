package ingresso

import (
	"fmt"
)

// State type represents Ingresso.com's Database State
type State struct {
	Name   string `json:"name"`
	UF     string `json:"uf"`
	Cities []City `json:"cities"`
}

// City type represents Ingresso.com's Database City
type City struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UF       string `json:"uf"`
	State    string `json:"state"`
	URLKey   string `json:"urlKey"`
	TimeZone string `json:"timeZone"`
}

// GetStates Get all states and cities available
// https://api-content.ingresso.com/v0/swagger/ui/index#!/States/States_GetAsync
func (ing *Ingresso) GetStates() ([]State, error) {
	var states []State
	url := fmt.Sprintf("%sstates", ing.BaseURL)
	_, err := ing.GetIngresso(url, &states)
	return states, err
}

// GetState Get a single state and its cities
// https://api-content.ingresso.com/v0/swagger/ui/index#!/States/States_GetAsync_0
func (ing *Ingresso) GetState(uf string) (*State, error) {
	var state State
	url := fmt.Sprintf("%sstates/%s", ing.BaseURL, uf)
	_, err := ing.GetIngresso(url, &state)
	return &state, err
}

// GetCity Get a single city by Id
// https://api-content.ingresso.com/v0/swagger/ui/index#!/States/States_GetCityAsync
func (ing *Ingresso) GetCity(id string) (*City, error) {
	var city City
	url := fmt.Sprintf("%sstates/city/%s", ing.BaseURL, id)
	_, err := ing.GetIngresso(url, &city)
	return &city, err
}

// GetCityURLKey Get a single city by urlKey
// https://api-content.ingresso.com/v0/swagger/ui/index#!/States/States_GetCityByUrlKeyAsync
func (ing *Ingresso) GetCityURLKey(urlKey string) (*City, error) {
	var city City
	url := fmt.Sprintf("%sstates/city/name/%s", ing.BaseURL, urlKey)
	_, err := ing.GetIngresso(url, &city)
	return &city, err
}
