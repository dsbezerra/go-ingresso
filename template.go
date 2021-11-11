package ingresso

import "fmt"

// Template struct
type Template struct {
	Items []Event `json:"items"`
	Count int     `json:"count"`
}

// GetSoonEvents gets all soon events
// https://api-content.ingresso.com/v0/templates/soon/partnership/{partnership}
func (ing *Ingresso) GetSoonEvents() (*Template, error) {
	var template Template
	url := fmt.Sprintf("%stemplates/soon/partnership/%s", ing.BaseURL, ing.Partnership)
	_, err := ing.getIngresso(url, &template)
	return &template, err
}

// GetSoonEventsByCity gets all soon events
// https://api-content.ingresso.com/v0/templates/soon/{cityId}/partnership/{partnership}
func (ing *Ingresso) GetSoonEventsByCity(city string) (*Template, error) {
	var template Template
	url := fmt.Sprintf("%stemplates/soon/%s/partnership/%s", ing.BaseURL, city, ing.Partnership)
	_, err := ing.getIngresso(url, &template)
	return &template, err
}

// GetHighlightEvents gets all highlights events by city and theater
// https://api-content.ingresso.com/v0/templates/highlights/{cityId}/partnership/{partnership}
func (ing *Ingresso) GetHighlightEvents(city string, options ...func(*QueryOptions)) (*Template, error) {
	var template Template
	url := fmt.Sprintf("%stemplates/highlights/%s/partnership/%s", ing.BaseURL, city, ing.Partnership)
	_, err := ing.getIngresso(url, &template, options...)
	return &template, err
}

// GetNowPlayingEvents gets all now playing events
// https://api-content.ingresso.com/v0/templates/nowplaying/partnership/{partnership}
func (ing *Ingresso) GetNowPlayingEvents(options ...func(*QueryOptions)) (*Template, error) {
	var template Template
	url := fmt.Sprintf("%stemplates/nowplaying/partnership/%s", ing.BaseURL, ing.Partnership)
	_, err := ing.getIngresso(url, &template, options...)
	return &template, err
}

// GetNowPlayingEventsByCity gets all now playing events by city
// https://api-content.ingresso.com/v0/templates/nowplaying/{cityId}/partnership/{partnership}
func (ing *Ingresso) GetNowPlayingEventsByCity(city string, options ...func(*QueryOptions)) (*Template, error) {
	var template Template
	url := fmt.Sprintf("%stemplates/nowplaying/%s/partnership/%s", ing.BaseURL, city, ing.Partnership)
	_, err := ing.getIngresso(url, &template, options...)
	return &template, err
}

// GetPremiereEventsByCity gets all premiere events by city
// https://api-content.ingresso.com/v0/templates/premiere/{cityId}/partnership/{partnership}
func (ing *Ingresso) GetPremiereEventsByCity(city string) (*Template, error) {
	var template Template
	url := fmt.Sprintf("%stemplates/premiere/%s/partnership/%s", ing.BaseURL, city, ing.Partnership)
	_, err := ing.getIngresso(url, &template)
	return &template, err
}
