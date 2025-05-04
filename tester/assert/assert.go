package assert

import (
	"maps"
	"strings"
	"testing"
)

// Will not check equality for slices or maps.
func Equals(t *testing.T, expected, actual interface{}) {
	t.Helper()

	if expected != actual {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func NotNil(t *testing.T, value interface{}) {
	t.Helper()

	if value == nil {
		t.Error("Value is nil.")
	}
}

func True(t *testing.T, value bool) {
	t.Helper()

	if !value {
		t.Error("Value is false.")
	}
}

func GreaterOrEqual(t *testing.T, v1, v2 int) {
	t.Helper()

	if v1 < v2 {
		t.Errorf("Expected value %d to be greater than or equal to %d.", v1, v2)
	}
}

func NoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func Contains(t *testing.T, s, sub string) {
	t.Helper()

	if !strings.Contains(s, sub) {
		t.Errorf("Expected %s to be in %s", sub, s)
	}
}

func Panics(t *testing.T, f func()) {
	t.Helper()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected function to panic.")
		}
	}()

	f()
}

func MapsEqual[K comparable, V comparable](t *testing.T, expected, actual map[K]V) {
	t.Helper()

	if !maps.Equal(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
