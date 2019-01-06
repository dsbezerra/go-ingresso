package ingresso

import (
	"fmt"
	"time"
)

// Event type represents Ingresso.com's Database Event
type Event struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	OriginalTitle  string `json:"originalTitle"`
	AncineID       string `json:"ancineId"`
	CountryOrigin  string `json:"countryOrigin"`
	Priority       int    `json:"priority"`
	ContentRating  string `json:"contentRating"`
	Duration       string `json:"duration"`
	Rating         int    `json:"rating"`
	Synopsis       string `json:"synopsis"`
	Cast           string `json:"cast"`
	Director       string `json:"director"`
	Distributor    string `json:"distributor"`
	InPreSale      bool   `json:"inPreSale"`
	Type           string `json:"type"`
	URLKey         string `json:"urlKey"`
	IsPlaying      bool   `json:"isPlaying"`
	CountIsPlaying int    `json:"countIsPlaying"`
	PremiereDate   struct {
		LocalDate   time.Time `json:"localDate"`
		IsToday     bool      `json:"isToday"`
		DayOfWeek   string    `json:"dayOfWeek"`
		DayAndMonth string    `json:"dayAndMonth"`
		Hour        string    `json:"hour"`
		Year        string    `json:"year"`
	} `json:"premiereDate"`
	CreationDate time.Time `json:"creationDate"`
	City         string    `json:"city"`
	SiteURL      string    `json:"siteURL"`
	Images       []Image   `json:"images"`
	Genres       []string  `json:"genres"`
	CompleteTags []struct {
		Name       string `json:"name"`
		Background string `json:"background"`
		Color      string `json:"color"`
	} `json:"completeTags"`
	Tags     []string `json:"tags"`
	Trailers []struct {
		Type        string `json:"type"`
		URL         string `json:"url"`
		EmbeddedURL string `json:"embeddedUrl"`
	} `json:"trailers"`
	BoxOfficeID     string `json:"boxOfficeId"`
	PartnershipType string `json:"partnershipType"`
}

// EventsResult struct
type EventsResult struct {
	Items []Event `json:"items"`
	Count int     `json:"count"`
}

// GetEvents gets all events available
// https://api-content.ingresso.com/v0/swagger/ui/index#!/Events/Events_GetAsync
func (ing *Ingresso) GetEvents() (*EventsResult, error) {
	var eventsResult EventsResult
	url := fmt.Sprintf("%sevents/partnership/%s", ing.BaseURL, ing.Partnership)
	_, err := ing.getIngresso(url, &eventsResult)
	return &eventsResult, err
}

// GetEvent gets a single event
// https://api-content.ingresso.com/v0/swagger/ui/index#!/Events/Events_Get
func (ing *Ingresso) GetEvent(id string) (*Event, error) {
	var event Event
	url := fmt.Sprintf("%sevents/%s/partnership/%s", ing.BaseURL, id, ing.Partnership)
	_, err := ing.getIngresso(url, &event)
	return &event, err
}

// GetEventByURLKey gets a single event by his URLKey
// https://api-content.ingresso.com/v0/swagger/ui/index#!/Events/Events_GetEventBySlugAsync
func (ing *Ingresso) GetEventByURLKey(urlKey string) (*Event, error) {
	var event Event
	url := fmt.Sprintf("%sevents/url-key/%s/partnership/%s", ing.BaseURL, urlKey, ing.Partnership)
	_, err := ing.getIngresso(url, &event)
	return &event, err
}

// GetComingSoonEvents gets coming soon events available
// https://api-content.ingresso.com/v0/swagger/ui/index#!/Events/Events_GetComingSoonAsync
func (ing *Ingresso) GetComingSoonEvents(options ...func(*QueryOptions)) (*EventsResult, error) {
	var eventsResult EventsResult
	url := fmt.Sprintf("%sevents/coming-soon/partnership/%s", ing.BaseURL, ing.Partnership)
	_, err := ing.getIngresso(url, &eventsResult, options...)
	return &eventsResult, err
}

// GetEventsByCity gets all events available
// https://api-content.ingresso.com/v0/swagger/ui/index#!/Events/Events_GetByCityAsync
func (ing *Ingresso) GetEventsByCity(city string, options ...func(*QueryOptions)) (*EventsResult, error) {
	var eventsResult EventsResult
	url := fmt.Sprintf("%sevents/city/%s/partnership/%s", ing.BaseURL, city, ing.Partnership)
	_, err := ing.getIngresso(url, &eventsResult, options...)
	return &eventsResult, err
}
