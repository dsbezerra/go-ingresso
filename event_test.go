package ingresso

import (
	"testing"
)

func TestGetComingSoonEvents(t *testing.T) {
	ing := New()
	_, err := ing.GetComingSoonEvents(2)
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}

func Test404GetEvent(t *testing.T) {
	ing := New()
	_, err := ing.GetEvent("XX") // Do not exist
	if err != nil && err.Error() != "404 Not Found" {
		t.Errorf("It should fail with '404 Not Found', but got: %v", err)
	}
}
