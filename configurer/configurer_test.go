package configurer

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/saeedafzal/sandbox-service/tester/assert"
)

func TestLoadConfiguration_Defaults(t *testing.T) {
	flag.String("foo", "bar", "")

	LoadConfiguration()

	actual := c.merged["foo"]
	assert.Equals(t, "bar", actual)

	reset()
}

func TestLoadConfiguration_FlagsOverrideDefaults(t *testing.T) {
	flag.String("foo", "bar", "")
	assert.NoError(t, flag.CommandLine.Parse([]string{"-foo", "something_else"}))

	LoadConfiguration()

	actual := c.merged["foo"]
	assert.Equals(t, "something_else", actual)

	reset()
}

func TestLoadConfiguration_defaultOverridenByFile(t *testing.T) {
	flag.String("foo", "bar", "")

	dir := t.TempDir()
	path := filepath.Join(dir, "config.conf")
	assert.NoError(t, os.WriteFile(path, []byte("foo = not_bar"), 0644))
	flag.String("config", path, "")

	assert.NoError(t, flag.CommandLine.Parse([]string{"-config", path}))

	LoadConfiguration()

	actual := c.merged["foo"]
	assert.Equals(t, "not_bar", actual)

	reset()
}

func reset() {
	c = &Configurer{
		file:   make(map[string]string),
		flags:  make(map[string]string),
		merged: make(map[string]string),
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
}
