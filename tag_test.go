package ingresso

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTags(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetTags()
	assert.NoError(t, err, "It should not fail, but got an error")
}
