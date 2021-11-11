package ingresso

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURL = "https://api-content.ingresso.com/v0/"
)

type (
	// Ingresso struct
	Ingresso struct {
		// HTTP client used to communicate with the API
		client *http.Client
		// BaseURL of Ingresso's website
		BaseURL *url.URL
		// Partnership name used to authenticate requests
		Partnership string
	}
	// Image struct
	Image struct {
		URL  string `json:"url"`
		Type string `json:"type"`
	}

	// QueryOptions map used to build query strings
	QueryOptions map[string]interface{}
)

// New allocates and initializes a new Ingresso instance.
func New(options ...func(ing *Ingresso)) *Ingresso {
	ing := &Ingresso{
		client: http.DefaultClient,
	}
	u, _ := url.Parse(baseURL)
	ing.BaseURL = u
	for _, f := range options {
		f(ing)
	}
	return ing
}

// Partnership sets the partnership code
func Partnership(p string) func(*Ingresso) {
	return func(ing *Ingresso) {
		ing.Partnership = p
	}
}

// Date sets the date (in format YYYY-MM-DD) query option
func Date(date string) func(*QueryOptions) {
	return func(opts *QueryOptions) {
		(*opts)["date"] = date
	}
}

// TheaterIDs sets the theaterIds query option
func TheaterIDs(IDs ...string) func(*QueryOptions) {
	return func(opts *QueryOptions) {
		(*opts)["theaterIds"] = strings.Join(IDs, ",")
	}
}

// JustEvents sets the justEvents query option
func JustEvents(justEvents bool) func(*QueryOptions) {
	return func(opts *QueryOptions) {
		(*opts)["justEvents"] = justEvents
	}
}

// Limit sets the limit query option
func Limit(limit int) func(*QueryOptions) {
	return func(opts *QueryOptions) {
		(*opts)["limit"] = limit
	}
}

// Skip sets the skip query option
func Skip(skip int) func(*QueryOptions) {
	return func(opts *QueryOptions) {
		(*opts)["skip"] = skip
	}
}

// ToJSON converts data from struct to JSON
func ToJSON(data interface{}) (string, error) {
	r := []byte("{}")
	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return string(r), err
	}
	return string(json), err
}

// builds a query string from the given options
func buildQuery(options ...func(*QueryOptions)) string {
	if len(options) == 0 {
		return ""
	}

	opts := QueryOptions{}

	for _, f := range options {
		f(&opts)
	}

	var b strings.Builder
	for param, value := range opts {
		b.WriteString(fmt.Sprintf("&%s=%v", param, value))
	}

	return b.String()
}

// calls Ingresso's API
func (ing *Ingresso) getIngresso(url string, data interface{}, options ...func(*QueryOptions)) (interface{}, error) {
	query := buildQuery(options...)
	if query != "" {
		if !strings.Contains(url, "?") {
			url = fmt.Sprintf("%s?%s", url, query)
		} else {
			url += query
		}
	}

	resp, err := ing.client.Get(url)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		json.Unmarshal(body, &data)
		return data, nil
	}

	return data, errors.New(resp.Status)
}
