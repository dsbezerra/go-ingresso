package ingresso

import "fmt"

// Carousel type represents Ingresso.com's Carousel
type Carousel struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	CarouselSlug string  `json:"carouselSlug"`
	Priority     int     `json:"priority"`
	Type         string  `json:"type"`
	HasLink      bool    `json:"hasLink"`
	Events       []Event `json:"events"`
}

// GetCaroulsels gets all carousels available for a given city
// https://api-content.ingresso.com/v0/carousel/{cityId}/partnership/${partnership}
func (ing *Ingresso) GetCarousels(cityId string) ([]Carousel, error) {
	var result []Carousel
	url := fmt.Sprintf("%scarousel/%s/partnership/%s", ing.BaseURL, cityId, ing.Partnership)
	_, err := ing.getIngresso(url, &result)
	return result, err
}
