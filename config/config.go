package config

import (
	"flag"
	"fmt"
	"log/slog"

	"github.com/saeedafzal/sandbox-service/store"
)

var version string

// Returns if application should terminate after flag.
func Flags() bool {
	h := flag.Bool("help", false, "Usage of sandbox-service.")
	v := flag.Bool("version", false, "Show application version.")
	d := flag.Bool("debug", false, "Enable debug mode.")
	p := flag.Int("port", 8080, "Port to bind server to.")
	flag.Parse()

	if *h {
		flag.Usage()
		return true
	}

	if *v {
		fmt.Println("Sandbox Service:", version)
		return true
	}

	if *d {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Debug mode enabled.")
	}

	store.GlobalStore.Put("version", version)
	store.GlobalStore.Put("port", *p)
	return false
}
