package ingresso

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
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

// GetIngresso performs a GET request to Ingresso's API.
func (ing *Ingresso) GetIngresso(url string, data interface{}) (interface{}, error) {
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

// ToJSON converts data from struct to JSON
func ToJSON(data interface{}) (string, error) {
	r := []byte("{}")
	r, err := json.MarshalIndent(data, "", "  ")
	return string(r), err
}
