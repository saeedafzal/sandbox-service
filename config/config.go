package config

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/saeedafzal/sandbox-service/configurer"
	"github.com/saeedafzal/sandbox-service/store"
)

var version string

func Initialise() {
	flag.Bool("help", false, "Usage of sandbox-service.")
	flag.Bool("version", false, "Show application version.")
	flag.Bool("debug", false, "Enable debug mode.")
	flag.Int("port", 8080, "Port to bind server to.")
	flag.String("config", "", "Path to the configuration file.")
	flag.Parse()

	configurer.LoadConfiguration()

	// Handle flags that can terminate
	if store.GetBool("help") {
		flag.Usage()
		os.Exit(0)
	}

	if store.GetBool("version") {
		fmt.Println("Sandbox Service:", version)
		os.Exit(0)
	}
	store.Put("version", version)

	// Application flags
	if store.GetBool("debug") {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Debug mode enabled.")
	}
}
