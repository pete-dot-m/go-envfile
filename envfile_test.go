package envfile

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// Tests LoadEnv by first creating a temporary env file and parses it, setting the values in the environment
// Then checks the environment to ensure values are properly set
func Test_envfile_LoadEnv(t *testing.T) {
	var tests = []struct {
		wantKey   string
		wantValue string
	}{
		{"ONE", "1"},
		{"NAME", "My Name"},
		{"COLOR_BLUE", "0x0000FF"},
	}

	fileNames := []string{"", "my.env"}

	for _, fileName := range fileNames {
		prepare(t, fileName, tests)

		var err error
		if fileName == "" {
			err = LoadEnv()
		} else {
			err = LoadEnv(fileName)
		}
		if err != nil {
			t.Fatalf("LoadEnv() got %q, want nil", err)
		}

		for _, tt := range tests {
			testName := "Load env from default file"
			t.Run(testName, func(t *testing.T) {
				gotValue := os.Getenv(tt.wantKey)
				if gotValue != tt.wantValue {
					t.Errorf("os.Getenv(%s) failed, got %s want %s", tt.wantKey, gotValue, tt.wantValue)
				}
			})
		}

		cleanup(t, fileName)
	}
}

func prepare(t *testing.T, fileName string, tests []struct {
	wantKey   string
	wantValue string
}) {
	var name string
	if fileName == "" {
		name = ".env"
	} else {
		name = fileName
	}
	file, err := os.Create(name)
	if err != nil {
		t.Fatalf("Unable to create test %s file: %s", name, err.Error())
	}

	writer := bufio.NewWriter(file)
	for _, tt := range tests {
		writer.WriteString(fmt.Sprintf("%s=%s\n", tt.wantKey, tt.wantValue))
	}
	writer.Flush()
}

func cleanup(t *testing.T, fileName string) {
	var name string
	if fileName == "" {
		name = ".env"
	} else {
		name = fileName
	}
	if err := os.Remove(name); err != nil {
		t.Errorf("Cleanup: couldn remove %s", name)
	}
}
