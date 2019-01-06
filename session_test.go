package ingresso

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test404GetSession(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetSession("XX") // Do not exist
	assert.EqualError(t, err, "404 Not Found", "It should fail with '404 Not Found'")
}

func TestGetSessionsByCityTheater(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetSessionsByCityTheater("21", "372") // Sessions of Cinemark BH Shopping in Belo Horizonte
	assert.NoError(t, err, "It should not fail, but got an error")
}

// NOTE: This should return 404 Not Found as stated in the docs
// but is returning 204 No Content at the time this was written
// 5 jan 2019
func Test204GetSessionsByCityTheater(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	sessions, _ := ing.GetSessionsByCityTheater("XX", "XX") // Do not exist
	assert.Equal(t, 0, len(sessions), "Expected 0 sessions, but got %d", len(sessions))
}

func TestGetSessionsByCityTheaterURLKey(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetSessionsByCityTheaterURLKey("21", "cinemark-bh-shopping") // Sessions of Cinemark BH Shopping in Belo Horizonte
	assert.NoError(t, err, "It should not fail, but got an error")
}

func Test404GetSessionsByCityTheaterURLKey(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetSessionsByCityTheaterURLKey("XX", "definitely-not-a-theater-url-key") // Do not exist
	assert.EqualError(t, err, "404 Not Found", "It should fail with '404 Not Found'")
}

// NOTE: If this fails check the Test204GetSessionsByCityTheater note
func Test204SessionsByCityEvent(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	sessions, _ := ing.GetSessionsByCityEvent("21", "XX") // Do not exist
	assert.Equal(t, 0, len(sessions), "Expected 0 sessions, but got %d", len(sessions))
}

func Test404SessionsByCityEventURLKey(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetSessionsByCityEventURLKey("21", "definitely-not-an-event-url-key") // Do not exist
	assert.EqualError(t, err, "404 Not Found", "It should fail with '404 Not Found'")
}

func TestSessionsDaysByCityTheater(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetSessionsDaysByCityTheater("21", "372") // Session days of Cinemark BH Shopping in Belo Horizonte
	assert.NoError(t, err, "It should not fail, but got an error")
}

// NOTE: If this fails check the Test204GetSessionsByCityTheater note
func Test204SessionsDaysByCityTheater(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	sessions, _ := ing.GetSessionsDaysByCityTheater("XX", "372") // Do not exists
	assert.Equal(t, 0, len(sessions), "Expected 0 sessions, but got %d", len(sessions))
}

// NOTE: If this fails check the Test204GetSessionsByCityTheater note
func Test204SessionsDaysByCityEvent(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	sessions, _ := ing.GetSessionsDaysByCityEvent("XX", "XX") // Do not exists
	assert.Equal(t, 0, len(sessions), "Expected 0 sessions, but got %d", len(sessions))
}
