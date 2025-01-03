package envfile

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// LoadEnv reads in the .envfile in the same directory (default), or the file specified
// and sets each name=value pair in the current environment.
// Returns an error or nil on sucess.
func LoadEnv(envFile ...string) error {
	var envFilePath string
	len := len(envFile)
	switch len {
	case 0:
		envFilePath = ".env"
	case 1:
		envFilePath = envFile[0]
	default:
		return errors.New(fmt.Sprintf("Too many arguments passed to LoadEnv. Want 0-1, got %d", len))
	}
	return loadEnvFromFile(envFilePath)
}

// loadEnvFromFile handles scanning the envfile and sets each name=value in the
// current process environment.
// Returns an error or nil on success.
func loadEnvFromFile(envFile string) error {
	// load the file
	file, err := os.Open(envFile)
	if err != nil {
		return err
	}

	// read a line at a time
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		// parse the line: NAME=Value
		line := fileScanner.Text()

		// we can just skip blank lines
		if len(line) == 0 {
			continue
		}
		// but we need the line to have an '='
		if !strings.Contains(line, "=") {
			return errors.New(fmt.Sprintf("Unable to parse line %q", line))
		}

		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			return errors.New(fmt.Sprintf("Unable to parse line %q, not enough parts", line))
		}

		// set the values
		name, value := parts[0], parts[1]
		if err := os.Setenv(name, value); err != nil {
			return errors.New(fmt.Sprintf("Unable to set env var %s=%s: %s", name, value, err.Error()))
		}
	}

	return nil
}
