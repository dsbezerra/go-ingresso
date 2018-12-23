package ingresso

import (
	"testing"
)

func TestGetStates(t *testing.T) {
	ing := New()
	_, err := ing.GetStates()
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}

func TestGetState(t *testing.T) {
	ing := New()
	_, err := ing.GetState("MG")
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}

func Test404GetState(t *testing.T) {
	ing := New()
	_, err := ing.GetState("XX") // Do not exist
	if err != nil && err.Error() != "404 Not Found" {
		t.Errorf("It should fail with '404 Not Found', but got: %v", err)
	}
}

func TestGetCity(t *testing.T) {
	ing := New()
	_, err := ing.GetCity("21") // Belo Horizonte
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}

	_, err = ing.GetCityURLKey("belo-horizonte") // Belo Horizonte
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}
