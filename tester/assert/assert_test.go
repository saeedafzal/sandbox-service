package assert

import (
	"errors"
	"testing"
)

func TestEquals(t *testing.T) {
	cases := []struct {
		name     string
		expected interface{}
		actual   interface{}
		fails    bool
	}{
		{"equal strings", "hello", "hello", false},
		{"non-equal strings", "hello", "world", true},
		{"equal integers", 8, 8, false},
		{"non-equal integers", 8, 1, true},
		{"equal booleans", true, true, false},
		{"non-equal booleans", true, false, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			tst := &testing.T{}
			Equals(tst, c.expected, c.actual)

			if tst.Failed() != c.fails {
				t.Errorf("expected: %t, actual: %t", c.fails, tst.Failed())
			}
		})
	}
}

func TestNotNil(t *testing.T) {
	cases := []struct {
		name  string
		value interface{}
		fails bool
	}{
		{"non-nil value", "hello", false},
		{"nil value", nil, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			tst := &testing.T{}
			NotNil(tst, c.value)

			if tst.Failed() != c.fails {
				t.Errorf("expected: %t, actual: %t", c.fails, tst.Failed())
			}
		})
	}
}

func TestTrue(t *testing.T) {
	cases := []struct {
		name  string
		value bool
		fails bool
	}{
		{"true value", true, false},
		{"false value", false, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			tst := &testing.T{}
			True(tst, c.value)

			if tst.Failed() != c.fails {
				t.Errorf("expected: %t, actual: %t", c.fails, tst.Failed())
			}
		})
	}
}

func TestGreaterOrEqual(t *testing.T) {
	cases := []struct {
		name  string
		v1    int
		v2    int
		fails bool
	}{
		{"greater", 8, 2, false},
		{"equal", 8, 8, false},
		{"less", 1, 8, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			tst := &testing.T{}
			GreaterOrEqual(tst, c.v1, c.v2)

			if tst.Failed() != c.fails {
				t.Errorf("expected: %t, actual: %t", c.fails, tst.Failed())
			}
		})
	}
}

func TestNoError(t *testing.T) {
	cases := []struct {
		name  string
		value error
		fails bool
	}{
		{"with error", errors.New(""), true},
		{"with no error", nil, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			tst := &testing.T{}
			NoError(tst, c.value)

			if tst.Failed() != c.fails {
				t.Errorf("expected: %t, actual: %t", c.fails, tst.Failed())
			}
		})
	}
}

func TestContains(t *testing.T) {
	cases := []struct {
		name  string
		s     string
		sub   string
		fails bool
	}{
		{"has substring", "hello", "ello", false},
		{"does not have substring", "hello", "bye", true},
		{"empty substring", "hello", "", false},
		{"empty string", "", "why", true},
		{"string and substring empty", "", "", false},
		{"repeated substring", "banana", "ana", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			tst := &testing.T{}
			Contains(tst, c.s, c.sub)

			if tst.Failed() != c.fails {
				t.Errorf("expected: %t, actual: %t", c.fails, tst.Failed())
			}
		})
	}
}
