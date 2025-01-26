package config

import (
	"flag"
	"os"
	"testing"

	"github.com/saeedafzal/sandbox-service/tester/assert"
)

func TestFlags(t *testing.T) {
	cases := []struct {
		name     string
		arg      string
		expected bool
	}{
		{
			name:     "no flag",
			arg:      "",
			expected: false,
		},
		{
			name:     "unknown flag",
			arg:      "-unknown-test-flag",
			expected: false,
		},
		{
			name:     "help --help",
			arg:      "--help",
			expected: true,
		},
		{
			name:     "version --version",
			arg:      "--version",
			expected: true,
		},
		{
			name:     "debug --debug",
			arg:      "--debug",
			expected: false,
		},
		{
			name:     "port --port",
			arg:      "--port 10101",
			expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
			os.Args = []string{os.Args[0], c.arg}
			actual := Flags()
			assert.Equals(t, c.expected, actual)
		})
	}
}
