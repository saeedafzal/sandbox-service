package config

import (
	"flag"
	"os"
	"os/exec"
	"testing"

	"github.com/saeedafzal/sandbox-service/tester/assert"
)

// https://go.dev/talks/2014/testing.slide#23
func TestMain(m *testing.M) {
	if os.Getenv("CONFIG_TEST") == "1" {
		for i, arg := range os.Args {
			if arg == "--" {
				os.Args = append(os.Args[:1], os.Args[i+1:]...)
				break
			}
		}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		flag.CommandLine.SetOutput(os.Stdout)
		Initialise()
		os.Exit(0)
	}

	os.Exit(m.Run())
}

func TestInitialise_Help(t *testing.T) {
	out, code := runSubprocess(t, "--help")
	assert.Equals(t, 0, code)
	assert.Contains(t, string(out), "Usage of sandbox-service")
}

func TestInitialise_Version(t *testing.T) {
	out, code := runSubprocess(t, "--version")
	assert.Equals(t, 0, code)
	assert.Contains(t, string(out), "Sandbox Service:")
}

func TestInitialise_Debug(t *testing.T) {
	out, code := runSubprocess(t, "--debug")
	assert.Equals(t, 0, code)
	assert.Contains(t, string(out), "Debug mode enabled.")
}

// Helper function to run test sub-process
func runSubprocess(t *testing.T, arg string) ([]byte, int) {
	cmd := exec.Command(
		os.Args[0],
		"-test.run="+t.Name(),
		"--",
		arg,
	)
	cmd.Env = append(os.Environ(), "CONFIG_TEST=1")
	out, err := cmd.CombinedOutput()
	if exitErr, ok := err.(*exec.ExitError); ok {
		return out, exitErr.ExitCode()
	}
	return out, 0
}
