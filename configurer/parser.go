package configurer

import (
	"bufio"
	"log/slog"
	"os"
	"strings"
)

// Parses simple key=value configuration files.
func (c *Configurer) parseConfiguration(path string) {
	f, err := os.Open(path)
	if err != nil {
		slog.Error("Error opening configuration file:", "path", path, "err", err)
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Ignore blanks and comments
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Get key/value pair
		pair := strings.SplitN(line, "=", 2)
		if len(pair) != 2 {
			slog.Warn("Malformed line:", "line", line)
			panic(err)
		}

		key := strings.TrimSpace(pair[0])
		val := strings.TrimSpace(pair[1])
		c.file[key] = val
	}

	if err := scanner.Err(); err != nil {
		slog.Error("Error scanning configuration file:", "path", path, "err", err)
		panic(err)
	}

	slog.Info("Configuration loaded:", "path", path)
}
