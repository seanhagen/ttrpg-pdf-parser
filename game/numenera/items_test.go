package numenera

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func loadStringFromFile(t *testing.T, name string) string {
	t.Helper()
	bits, err := os.ReadFile(fmt.Sprintf("testdata/%v", name))
	if err != nil {
		t.Fatalf("unable to load file 'testdata/%v': %v", name, err)
	}

	return string(bits)
}

func loadDataFromFile(t *testing.T, name string, dest any) {
	t.Helper()
	bits, err := os.ReadFile(fmt.Sprintf("testdata/%v", name))
	if err != nil {
		t.Fatalf("unable to load file 'testdata/%v': %v", name, err)
	}

	err = json.Unmarshal(bits, dest)
	if err != nil {
		t.Fatalf("unable to unmarshal data from file 'testdata/%v': %v", name, err)
	}
}
