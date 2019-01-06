package ingresso

import "testing"

func TestGetTheater(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	theater, err := ing.GetTheater("372") // Cinemark BH Shopping
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
	name := "Cinemark BH Shopping"
	if name != theater.Name {
		t.Errorf("Expected name %s, but got: %s", name, theater.Name)
	}
}

func TestGetTheaters(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetTheaters()
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}

func TestGetTheatersByCity(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetTheatersByCity("21") // Belo Horizonte
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}

func Test404GetTheatersByCity(t *testing.T) {
	ing := New(
		Partnership("PARTNERSHIP_CODE"),
	)
	_, err := ing.GetTheatersByCity("XX") // Do not exist
	if err != nil && err.Error() != "404 Not Found" {
		t.Errorf("It should fail with '404 Not Found', but got: %v", err)
	}
}
