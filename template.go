package ingresso

import "fmt"

// Template struct
type Template struct {
	Items []Event `json:"items"`
	Count int     `json:"count"`
}

// GetSoonEvents gets all soon events
// https://api-content.ingresso.com/v0/swagger/ui/index#!/Templates/Templates_GetSoonAsync
func (ing *Ingresso) GetSoonEvents() (*Template, error) {
	var template Template
	url := fmt.Sprintf("%stemplates/soon/partnership/%s", ing.BaseURL, ing.Partnership)
	_, err := ing.GetIngresso(url, &template)
	return &template, err
}

// GetSoonEventsByCity gets all soon events
// https://api-content.ingresso.com/v0/swagger/ui/index#!/Templates/Templates_GetSoonByCityAsync
func (ing *Ingresso) GetSoonEventsByCity(city string) (*Template, error) {
	var template Template
	url := fmt.Sprintf("%stemplates/soon/%s/partnership/%s", ing.BaseURL, city, ing.Partnership)
	_, err := ing.GetIngresso(url, &template)
	return &template, err
}

// GetHighlightEvents gets all highlights events by city and theater
// https://api-content.ingresso.com/v0/swagger/ui/index#!/Templates/Templates_GetHighlightsByCityAsync
func (ing *Ingresso) GetHighlightEvents(city, theaterIds string, justEvents bool) (*Template, error) {
	var template Template
	url := fmt.Sprintf("%stemplates/highlights/%s/partnership/%s?", ing.BaseURL, city, ing.Partnership)
	if theaterIds != "" {
		url += fmt.Sprintf("theaterIds=%s&", theaterIds)
	}
	url += fmt.Sprintf("justEvents=%t", justEvents)
	_, err := ing.GetIngresso(url, &template)
	return &template, err
}
