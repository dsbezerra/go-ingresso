package ingresso

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test404GetEvent(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetEvent("XX") // Do not exist
	assert.EqualError(t, err, "404 Not Found", "It should fail with '404 Not Found'")
}

func Test404GetEventByURLKey(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetEventByURLKey("definitely-not-an-event-url-key") // Do not exist
	assert.EqualError(t, err, "404 Not Found", "It should fail with '404 Not Found'")
}

func TestGetEvents(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetEvents()
	assert.NoError(t, err, "It should not fail, but got an error")
}

func TestGetComingSoonEvents(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetComingSoonEvents()
	assert.NoError(t, err, "It should not fail, but got an error")
}

func TestGetEventsByCity(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetEventsByCity("21") // Events in Belo Horizonte
	assert.NoError(t, err, "It should not fail, but got an error")
}

func Test404GetEventsByCity(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetEventsByCity("XX") // Do not exist
	assert.EqualError(t, err, "404 Not Found", "It should fail with '404 Not Found'")
}
