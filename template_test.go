package ingresso

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSoonEvents(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetSoonEvents()
	assert.NoError(t, err, "It should not fail, but got an error")
}

func TestGetSoonEventsByCity(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetEventsByCity("21") // Belo Horizonte
	assert.NoError(t, err, "It should not fail, but got an error")
}

func TestGetSoonEventsByCityLimit(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	events, _ := ing.GetEventsByCity("21", Limit(2)) // 2 events in Belo Horizonte
	assert.Equal(t, 2, events.Count, "Expected 2 events, but got %d", events.Count)
}

// NOTE: If this fails check the Test204GetSessionsByCityTheater note in session_test.go
func Test204GetSoonEvents(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	events, _ := ing.GetEventsByCity("XX") // Do not exist
	assert.Equal(t, 0, events.Count, "Expected 0 events, but got %d", events.Count)
}

func TestGetHighlightEvents(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetHighlightEvents("21", TheaterIDs("372"), JustEvents(false))
	assert.NoError(t, err, "It should not fail, but got an error")
}

func TestGetNowPlayingEvents(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetNowPlayingEvents()
	assert.NoError(t, err, "It should not fail, but got an error")
}

func TestGetNowPlayingEventsLimit(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetNowPlayingEvents(Limit(2))
	assert.NoError(t, err, "It should not fail, but got an error")
}

// NOTE: Doesn't seem to be working (5 jan 2019)
//
// func TestGetNowPlayingEventsSkip(t *testing.T) {
// 	ing := New(
// 		Partnership("PARTNERSHIP_CODE"),
// 	)
// 	_, err := ing.GetNowPlayingEvents(Skip(1))
// 	assert.NoError(t, err, "It should not fail, but got an error")
// }

func TestGetNowPlayingEventsByCity(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetNowPlayingEventsByCity("21") // Belo Horizonte
	assert.NoError(t, err, "It should not fail, but got an error")
}

func TestGetNowPlayingEventsByCityLimit(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	events, err := ing.GetNowPlayingEventsByCity("21", Limit(1)) // Belo Horizonte
	assert.NoError(t, err, "It should not fail, but got an error")
	assert.Equal(t, 1, events.Count, "Expected 1, but got %d", events.Count)
}

func Test204GetNowPlayingEventsByCity(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	events, err := ing.GetNowPlayingEventsByCity("XX") // Do not exist
	assert.NoError(t, err, "It should not fail, but got an error")
	assert.Equal(t, 0, events.Count, "Expected 0, but got %d", events.Count)
}

func TestGetPremiereEventsByCity(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetPremiereEventsByCity("21") // Belo Horizonte
	assert.NoError(t, err, "It should not fail, but got an error")
}

func Test204GetPremiereEventsByCity(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	events, err := ing.GetPremiereEventsByCity("XX") // Do not exist
	assert.NoError(t, err, "It should not fail, but got an error")
	assert.Equal(t, 0, events.Count, "Expected 0, but got %d", events.Count)
}
