package envfile

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func Test_envfile_LoadEnv(t *testing.T) {
	var tests = []struct {
		wantKey   string
		wantValue string
	}{
		{"ONE", "1"},
		{"NAME", "My Name"},
		{"COLOR_BLUE", "0x0000FF"},
	}

	prepare(t, tests)

	err := LoadEnv()
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

	cleanup(t)
}

func prepare(t *testing.T, tests []struct {
	wantKey   string
	wantValue string
}) {
	fileName := ".env"
	file, err := os.Create(fileName)
	if err != nil {
		t.Fatalf("Unable to create test .env file: %s", err.Error())
	}

	writer := bufio.NewWriter(file)
	for _, tt := range tests {
		writer.WriteString(fmt.Sprintf("%s=%s\n", tt.wantKey, tt.wantValue))
	}
	writer.Flush()
}

func cleanup(t *testing.T) {
	fileName := ".env"
	if err := os.Remove(fileName); err != nil {
		t.Errorf("Cleanup: couldn remove %s", fileName)
	}
}
