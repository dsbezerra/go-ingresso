package ingresso

import "testing"

func TestGetHighlightEvents(t *testing.T) {
	ing := New()
	_, err := ing.GetHighlightEvents("21", "372", false)
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}
