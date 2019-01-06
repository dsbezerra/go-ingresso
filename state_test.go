package ingresso

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStates(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetStates()
	assert.NoError(t, err, "It should not fail, but got an error")
}

func TestGetState(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetState("MG")
	assert.NoError(t, err, "It should not fail, but got an error")
}

func Test404GetState(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetState("XX") // Do not exist
	assert.EqualError(t, err, "404 Not Found", "It should fail with '404 Not Found'")
}

func TestGetCity(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetCity("21") // Belo Horizonte
	assert.NoError(t, err, "It should not fail, but got an error")
}

func Test404GetCity(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetCity("XX") // Do not exist
	assert.EqualError(t, err, "404 Not Found", "It should fail with '404 Not Found'")
}

func TestGetCityURLKey(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetCityURLKey("belo-horizonte") // Belo Horizonte
	assert.NoError(t, err, "It should not fail, but got an error")
}

func Test404GetCityURLKey(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetCityURLKey("definitely-not-a-city-url-key") // Do not exist
	assert.EqualError(t, err, "404 Not Found", "It should fail with '404 Not Found'")
}
