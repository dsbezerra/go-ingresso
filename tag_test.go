package ingresso

import (
	"testing"
)

func TestGetTags(t *testing.T) {
	ing := New()
	_, err := ing.GetTags()
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}
