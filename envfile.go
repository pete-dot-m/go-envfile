package envfile

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

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
		key, value := parts[0], parts[1]
		log.Println(fmt.Sprintf("envfile.LoadEnv: setting environment variable: %s", key))
		if err := os.Setenv(key, value); err != nil {
			return errors.New(fmt.Sprintf("Unable to set env var %s=%s: %s", key, value, err.Error()))
		}
	}

	return nil
}
