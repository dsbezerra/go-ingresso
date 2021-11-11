package ingresso

import (
	"fmt"
	"time"
)

// Theater type represents Ingresso.com's Database Theater
type Theater struct {
	ID                    string  `json:"id"`
	Images                []Image `json:"images"`
	URLKey                string  `json:"urlKey"`
	Name                  string  `json:"name"`
	SiteURL               string  `json:"siteURL"`
	CNPJ                  string  `json:"cnpj"`
	DistrictAuthorization string  `json:"districtAuthorization"`
	Address               string  `json:"address"`
	AddressComplement     string  `json:"addressComplement"`
	Number                string  `json:"number"`
	CityID                string  `json:"cityId"`
	CityName              string  `json:"cityName"`
	State                 string  `json:"state"`
	UF                    string  `json:"uf"`
	Neighborhood          string  `json:"neighborhood"`
	Properties            struct {
		HasBomboniere bool `json:"hasBomboniere"`
	} `json:"properties"`
	Telephones  []string `json:"telephones"`
	Geolocation struct {
		Latitude  float32 `json:"lat"`
		Longitude float32 `json:"lng"`
	} `json:"geolocation"`
	DeliveryType                []string `json:"deliveryType"`
	Corporation                 string   `json:"corporation"`
	CorporationID               string   `json:"corporationId"`
	CorporationPriority         int      `json:"corporationPriority"`
	CorporationAvatarBackground string   `json:"corporationAvatarBackground"`
	Rooms                       []struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		FullName  string `json:"fullName"`
		Capacity  int    `json:"capacity"`
		Documents []struct {
			Name       string    `json:"name"`
			Number     string    `json:"number"`
			Expiration time.Time `json:"expiration"`
		} `json:"documents"`
	} `json:"rooms"`
	TotalRooms      int    `json:"totalRooms"`
	Enabled         bool   `json:"enabled"`
	BlockMessage    string `json:"blockMessage"`
	PartnershipType string `json:"partnershipType"`
}

// TheatersResult struct
type TheatersResult struct {
	Items []Theater `json:"items"`
	Count int       `json:"count"`
}

// GetTheaters gets all theaters available
// https://api-content.ingresso.com/v0/theaters/partnership/{partnership}
func (ing *Ingresso) GetTheaters() (*TheatersResult, error) {
	var theatersResult TheatersResult
	url := fmt.Sprintf("%stheaters/partnership/%s", ing.BaseURL, ing.Partnership)
	_, err := ing.getIngresso(url, &theatersResult)
	return &theatersResult, err
}

// GetTheater gets a single theater
// https://api-content.ingresso.com/v0/theaters/{id}/partnership/{partnership}
func (ing *Ingresso) GetTheater(id string) (*Theater, error) {
	var theater Theater
	url := fmt.Sprintf("%stheaters/%s/partnership/%s", ing.BaseURL, id, ing.Partnership)
	_, err := ing.getIngresso(url, &theater)
	return &theater, err
}

// GetTheaterByURLKey gets a single theater by his URLKey
// https://api-content.ingresso.com/v0/theaters​/url-key​/{uRLKey}​/partnership​/{partnership}
func (ing *Ingresso) GetTheaterByURLKey(urlKey string) (*Theater, error) {
	var theater Theater
	url := fmt.Sprintf("%stheaters/url-key/%s/partnership/%s", ing.BaseURL, urlKey, ing.Partnership)
	_, err := ing.getIngresso(url, &theater)
	return &theater, err
}

// GetTheatersByCity gets all theaters by city
// https://api-content.ingresso.com/v0/theaters/city/{cityId}/partnership/{partnership}
func (ing *Ingresso) GetTheatersByCity(city string) (*TheatersResult, error) {
	var theatersResult TheatersResult
	url := fmt.Sprintf("%stheaters/city/%s/partnership/%s", ing.BaseURL, city, ing.Partnership)
	_, err := ing.getIngresso(url, &theatersResult)
	return &theatersResult, err
}
