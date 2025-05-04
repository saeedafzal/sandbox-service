package configurer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/saeedafzal/sandbox-service/tester/assert"
)

var valid = `
# Some comment
key = value
another_key    =    123
`

func TestParseConfiguration(t *testing.T) {
	dir := t.TempDir()

	validPath := filepath.Join(dir, "config.conf")
	assert.NoError(t, os.WriteFile(validPath, []byte(valid), 0666))

	invalidPath := filepath.Join(dir, "invalid.conf")
	assert.NoError(t, os.WriteFile(invalidPath, []byte("bad line"), 0666))

	cases := []struct {
		name      string
		path      string
		willPanic bool
		expected  map[string]string
	}{
		{
			name:      "empty path",
			path:      "",
			willPanic: false,
			expected:  map[string]string{},
		},
		{
			name:      "whitespace path",
			path:      "          ",
			willPanic: false,
			expected:  map[string]string{},
		},
		{
			name:      "non-existent file",
			path:      "does_not_exist.conf",
			willPanic: true,
		},
		{
			name:      "valid configuration",
			path:      validPath,
			willPanic: false,
			expected: map[string]string{
				"key":         "value",
				"another_key": "123",
			},
		},
		{
			name:      "malformed line",
			path:      invalidPath,
			willPanic: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cfg := &Configurer{file: make(map[string]string)}

			if c.willPanic {
				assert.Panics(t, func() {
					cfg.parseConfiguration(c.path)
				})
				return
			}

			cfg.parseConfiguration(c.path)
			assert.MapsEqual(t, c.expected, cfg.file)
		})
	}
}
