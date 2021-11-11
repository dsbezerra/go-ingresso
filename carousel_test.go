package ingresso

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCarousels(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	result, err := ing.GetCarousels("364")
	assert.Nil(t, err)
	assert.NotEmpty(t, result)
}

func TestEmptyGetCarousels(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	result, err := ing.GetCarousels("XX")
	assert.Nil(t, err)
	assert.Empty(t, result)
}
