package assert

import "testing"

// Will not check equality for slices or maps.
func Equals(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func NotNil(t *testing.T, value interface{}) {
	if value == nil {
		t.Error("Value is nil.")
	}
}

func True(t *testing.T, value bool) {
	if !value {
		t.Error("Value is false.")
	}
}

func GreaterOrEqual(t *testing.T, v1, v2 int) {
	if v1 < v2 {
		t.Errorf("Expected value %d to be greater than or equal to %d.", v1, v2)
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
