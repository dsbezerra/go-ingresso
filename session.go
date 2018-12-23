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

// GetSession gets a single session
// https://api-content.ingresso.com/v0/swagger/ui/index#!/Sessions/Sessions_GetAsync
func (ing *Ingresso) GetSession(id string) (*Session, error) {
	var session Session
	url := fmt.Sprintf("%ssessions/%s/partnership/%s", ing.BaseURL, id, ing.Partnership)
	_, err := ing.GetIngresso(url, &session)
	return &session, err
}

// GetSessionsByCityTheater gets a list of sessions by city and theater
// https://api-content.ingresso.com/v0/swagger/ui/index#!/Sessions/Sessions_GetByCityTheaterAsync
func (ing *Ingresso) GetSessionsByCityTheater(city, theater, date string) ([]Session, error) {
	var sessions []Session
	url := fmt.Sprintf("%sessions/city/%s/theater/%s/partnership/%s", ing.BaseURL, city, theater, ing.Partnership)
	if date != "" {
		url += fmt.Sprintf("?date=%s", date)
	}
	_, err := ing.GetIngresso(url, &sessions)
	return sessions, err
}

// GetSessionsByCityTheaterURLKey gets a list of sessions by city and theater's URLKey
// https://api-content.ingresso.com/v0/swagger/ui/index#!/Sessions/Sessions_GetByCityTheaterUrlKeyAsync
func (ing *Ingresso) GetSessionsByCityTheaterURLKey(city, theaterURLKey, date string) ([]Session, error) {
	var sessions []Session
	url := fmt.Sprintf("%sessions/city/%s/theater/url-key/%s/partnership/%s", ing.BaseURL, city, theaterURLKey, ing.Partnership)
	if date != "" {
		url += fmt.Sprintf("?date=%s", date)
	}
	_, err := ing.GetIngresso(url, &sessions)
	return sessions, err
}

// GetSessionsByCityEvent gets a list of sessions by city and event
// https://api-content.ingresso.com/v0/swagger/ui/index#!/Sessions/Sessions_GetByCityEventAsync
func (ing *Ingresso) GetSessionsByCityEvent(city, event, date string) ([]Session, error) {
	var sessions []Session
	url := fmt.Sprintf("%ssessions/city/%s/event/%s/partnership/%s", ing.BaseURL, city, event, ing.Partnership)
	if date != "" {
		url += fmt.Sprintf("?date=%s", date)
	}
	_, err := ing.GetIngresso(url, &sessions)
	return sessions, err
}

// GetSessionsByCityEventURLKey gets a list of sessions by city and event's URLKey
// https://api-content.ingresso.com/v0/swagger/ui/index#!/Sessions/Sessions_GetByCityEventUrlKeyAsync
func (ing *Ingresso) GetSessionsByCityEventURLKey(city, eventURLKey, date string) ([]Session, error) {
	var sessions []Session
	url := fmt.Sprintf("%ssessions/city/%s/event/url-key/%s/partnership/%s", ing.BaseURL, city, eventURLKey, ing.Partnership)
	if date != "" {
		url += fmt.Sprintf("?date=%s", date)
	}
	_, err := ing.GetIngresso(url, &sessions)
	return sessions, err
}

// GetSessionsDaysByCityTheater gets the dates that have sessions in the informed theater
// https://api-content.ingresso.com/v0/swagger/ui/index#!/Sessions/Sessions_GetSessionsDaysByTheater
func (ing *Ingresso) GetSessionsDaysByCityTheater(city, theater string) ([]Session, error) {
	var sessions []Session
	url := fmt.Sprintf("%ssessions/city/%s/theater/%s/dates/partnership/%s", ing.BaseURL, city, theater, ing.Partnership)
	_, err := ing.GetIngresso(url, &sessions)
	return sessions, err
}

// GetSessionsDaysByCityEvent gets the available dates for the informed event
// https://api-content.ingresso.com/v0/swagger/ui/index#!/Sessions/Sessions_GetSessionsDaysByEvent
func (ing *Ingresso) GetSessionsDaysByCityEvent(city, event string) ([]Session, error) {
	var sessions []Session
	url := fmt.Sprintf("%ssessions/city/%s/event/%s/dates/partnership/%s", ing.BaseURL, city, event, ing.Partnership)
	_, err := ing.GetIngresso(url, &sessions)
	return sessions, err
}