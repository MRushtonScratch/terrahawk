package state

import (
	"testing"
	"path/filepath"
)

func TestPilotDetachSimpleState(t *testing.T) {

	path := filepath.Join("testData", "simple.tfstate")

	state, err := Detach(path)

	if err != nil {
		t.Error(err)
	}

	if len(state.Modules) != 2 {
		t.Error("Expected to map 2 modules")
	}
}

func TestPilotDetachComplexState(t *testing.T) {

	path := filepath.Join("testData", "complex.tfstate")

	state, err := Detach(path)

	if err != nil {
		t.Error(err)
	}

	if len(state.Modules) != 2 {
		t.Error("Expected to map 2 modules")
	}
}