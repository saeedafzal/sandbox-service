package store

import (
	"testing"

	"github.com/saeedafzal/sandbox-service/tester/assert"
)

func TestInitCreatesEmptyGlobalStore(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, globalStore)
	assert.NotNil(t, globalStore.data)
}

func TestPut(t *testing.T) {
	Put("hello", "world")
	value, ok := globalStore.data["hello"]

	assert.True(t, ok)
	assert.Equals(t, "world", value)

	clear(globalStore.data)
}

func TestGetString(t *testing.T) {
	cases := []struct {
		name     string
		key      string
		value    interface{}
		expected string
	}{
		{
			name:     "non-existing key",
			key:      "does not exist",
			value:    "does not matter",
			expected: "",
		},
		{
			name:     "key with string value",
			key:      "hello",
			value:    "world",
			expected: "world",
		},
		{
			name:     "key with non-string value",
			key:      "hello",
			value:    123,
			expected: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			globalStore.data["hello"] = c.value
			actual := GetString(c.key)
			assert.Equals(t, c.expected, actual)
		})
	}

	clear(globalStore.data)
}

func TestGetInt(t *testing.T) {
	cases := []struct {
		name     string
		key      string
		value    interface{}
		expected int
	}{
		{
			name:     "non-existing key",
			key:      "does not exist",
			value:    0,
			expected: 0,
		},
		{
			name:     "key with int value",
			key:      "hello",
			value:    123,
			expected: 123,
		},
		{
			name:     "key with negative int value",
			key:      "hello",
			value:    -123,
			expected: -123,
		},
		{
			name:     "key with non-int value",
			key:      "hello",
			value:    "blah",
			expected: 0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			globalStore.data["hello"] = c.value
			actual := GetInt(c.key)
			assert.Equals(t, c.expected, actual)
		})
	}

	clear(globalStore.data)
}

func TestGetBool(t *testing.T) {
	cases := []struct {
		name     string
		key      string
		value    interface{}
		expected bool
	}{
		{
			name:     "non-existing key",
			key:      "does not exist",
			value:    false,
			expected: false,
		},
		{
			name:     "key with true value",
			key:      "hello",
			value:    true,
			expected: true,
		},
		{
			name:     "key with false value",
			key:      "hello",
			value:    false,
			expected: false,
		},
		{
			name:     "key with non-bool value",
			key:      "hello",
			value:    "this is not a bool",
			expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			globalStore.data["hello"] = c.value
			actual := GetBool(c.key)
			assert.Equals(t, c.expected, actual)
		})
	}

	clear(globalStore.data)
}
