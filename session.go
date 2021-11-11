package ingresso

import (
	"fmt"
	"time"
)

// Session type represents Ingresso.com's Database Session
type Session struct {
	SiteURL string   `json:"siteURL"`
	Room    string   `json:"room"`
	Theater Theater  `json:"theater"`
	Event   Event    `json:"event"`
	ID      string   `json:"id"`
	Price   int      `json:"price"`
	Type    []string `json:"type"`
	Types   []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Alias   string `json:"alias"`
		Display bool   `json:"display"`
	} `json:"types"`
	Date struct {
		LocalDate   time.Time `json:"localDate"`
		IsToday     bool      `json:"isToday"`
		DayOfWeek   string    `json:"dayOfWeek"`
		DayAndMonth string    `json:"dayAndMonth"`
		Hour        string    `json:"hour"`
		Year        string    `json:"year"`
	} `json:"date"`
	RealDate struct {
		LocalDate   time.Time `json:"localDate"`
		IsToday     bool      `json:"isToday"`
		DayOfWeek   string    `json:"dayOfWeek"`
		DayAndMonth string    `json:"dayAndMonth"`
		Hour        string    `json:"hour"`
		Year        string    `json:"year"`
	} `json:"realDate"`
	Time             string `json:"time"`
	DefaultSector    string `json:"defaultSector"`
	MidnightMessage  string `json:"midnightMessage"`
	HasSeatSelection bool   `json:"hasSeatSelection"`
	Enabled          bool   `json:"enabled"`
	BlockMessage     string `json:"blockMessage"`
}

// Day type represents Ingresso.com's Database Day
type Day struct {
	SessionTypes  []string `json:"sessionTypes"`
	Date          string   `json:"date"`
	DateFormatted string   `json:"dateFormatted"`
	DayOfWeek     string   `json:"dayOfWeek"`
	IsToday       bool     `json:"isToday"`
}

// GetSession gets a single session
// https://api-content.ingresso.com/v0/sessions/{id}/partnership/{partnership}
func (ing *Ingresso) GetSession(id string) (*Session, error) {
	var session Session
	url := fmt.Sprintf("%ssessions/%s/partnership/%s", ing.BaseURL, id, ing.Partnership)
	_, err := ing.getIngresso(url, &session)
	return &session, err
}

// GetSessionsByCityTheater gets a list of sessions by city and theater
// https://api-content.ingresso.com/v0/sessions/city/{cityId}/theater/{theaterId}/partnership/{partnership}
func (ing *Ingresso) GetSessionsByCityTheater(city, theater string, options ...func(*QueryOptions)) ([]Session, error) {
	var sessions []Session
	url := fmt.Sprintf("%ssessions/city/%s/theater/%s/partnership/%s", ing.BaseURL, city, theater, ing.Partnership)
	_, err := ing.getIngresso(url, &sessions, options...)
	return sessions, err
}

// GetSessionsByCityTheaterURLKey gets a list of sessions by city and theater's URLKey
// https://api-content.ingresso.com/v0/sessions/city/{cityId}/theater/url-key/{theaterUrlKey}/partnership/{partnership}
func (ing *Ingresso) GetSessionsByCityTheaterURLKey(city, theaterURLKey string, options ...func(*QueryOptions)) ([]Session, error) {
	var sessions []Session
	url := fmt.Sprintf("%ssessions/city/%s/theater/url-key/%s/partnership/%s", ing.BaseURL, city, theaterURLKey, ing.Partnership)
	_, err := ing.getIngresso(url, &sessions, options...)
	return sessions, err
}

// GetSessionsByCityEvent gets a list of sessions by city and event
// https://api-content.ingresso.com/v0/sessions/city/{cityId}/event/{eventId}/partnership/{partnership}
func (ing *Ingresso) GetSessionsByCityEvent(city, event string, options ...func(*QueryOptions)) ([]Session, error) {
	var sessions []Session
	url := fmt.Sprintf("%ssessions/city/%s/event/%s/partnership/%s", ing.BaseURL, city, event, ing.Partnership)
	_, err := ing.getIngresso(url, &sessions, options...)
	return sessions, err
}

// TODO(diego): https://api-content.ingresso.com/v0/sessions/event/{eventId}/partnership/{partnership}/streaming

// GetSessionsByCityEventURLKey gets a list of sessions by city and event's URLKey
// https://api-content.ingresso.com/v0/sessions/city/{cityId}/event/url-key/{eventUrlKey}/partnership/{partnership}
func (ing *Ingresso) GetSessionsByCityEventURLKey(city, eventURLKey string, options ...func(*QueryOptions)) ([]Session, error) {
	var sessions []Session
	url := fmt.Sprintf("%ssessions/city/%s/event/url-key/%s/partnership/%s", ing.BaseURL, city, eventURLKey, ing.Partnership)
	_, err := ing.getIngresso(url, &sessions, options...)
	return sessions, err
}

// GetSessionsDaysByCityTheater gets the dates that have sessions in the informed theater
// https://api-content.ingresso.com/v0/sessions/city/{cityId}/theater/{theaterId}/dates/partnership/{partnership}
func (ing *Ingresso) GetSessionsDaysByCityTheater(city, theater string) ([]Day, error) {
	var days []Day
	url := fmt.Sprintf("%ssessions/city/%s/theater/%s/dates/partnership/%s", ing.BaseURL, city, theater, ing.Partnership)
	_, err := ing.getIngresso(url, &days)
	return days, err
}

// GetSessionsDaysByCityEvent gets the available dates for the informed event
// https://api-content.ingresso.com/v0/sessions/city/{cityId}/event/{eventId}/dates/partnership/{partnership}
func (ing *Ingresso) GetSessionsDaysByCityEvent(city, event string) ([]Day, error) {
	var days []Day
	url := fmt.Sprintf("%ssessions/city/%s/event/%s/dates/partnership/%s", ing.BaseURL, city, event, ing.Partnership)
	_, err := ing.getIngresso(url, &days)
	return days, err
}

// TODO(diego): https://api-content.ingresso.com/v0/sessions/url-key/{urlKey}/exists
