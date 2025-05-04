package configurer

import (
	"flag"
	"maps"
	"strconv"

	"github.com/saeedafzal/sandbox-service/store"
)

type Configurer struct {
	// defaults map[string]string
	file   map[string]string
	flags  map[string]string
	merged map[string]string
}

var (
	c       *Configurer
	version = "dev"
)

func init() {
	c = &Configurer{
		// defaults: make(map[string]string),
		file:   make(map[string]string),
		flags:  make(map[string]string),
		merged: make(map[string]string),
	}
}

// If you need custom configuration file path flag, define it first as a flag
// defined as `config` i.e `flag.String("config", "config.conf", "Path to file.")`.
func LoadConfiguration() {
	// Get all flag defaults
	flag.VisitAll(func(f *flag.Flag) {
		c.merged[f.Name] = f.DefValue
	})

	// Store actual set flags
	flag.Visit(func(f *flag.Flag) {
		c.file[f.Name] = f.Value.String()
	})

	// Parse and load configuration file
	if path := GetConfigPath(); path != "" {
		c.parseConfiguration(path)
	}

	// Merge maps with priority: defaults -> file -> flags
	// maps.Copy(c.merged, c.defaults)
	maps.Copy(c.merged, c.file)
	maps.Copy(c.merged, c.flags)
	// c.merged = merged

	// Store values in global store
	for key, raw := range c.merged {
		val := cast(raw)
		store.Put(key, val)
	}
}

func GetConfigPath() string {
	if val, ok := c.flags["config"]; ok {
		return val
	}
	if val, ok := c.merged["config"]; ok {
		return val
	}
	return ""
}

func cast(raw string) interface{} {
	if val, err := strconv.ParseBool(raw); err == nil {
		return val
	}
	if val, err := strconv.Atoi(raw); err == nil {
		return val
	}

	return raw
}
