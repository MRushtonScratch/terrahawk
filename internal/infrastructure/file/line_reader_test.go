package file

import (
	"testing"
	"path/filepath"
	"errors"
	"fmt"
)

func TestReadFileLines(t *testing.T) {

	path := filepath.Join("testData", "readLines.txt")

	lines, err := ReadLines(path)

	if err != nil {
		t.Error(err)
	}


	if len(lines) != 2 {
		t.Error("Expected two lines in file")
	}

	result, err := assertTrue("Hello World", lines[0], "")

	if !result {
		t.Error(err)
	}

	result, err = assertTrue("World Hello", lines[1], "")

	if !result {
		t.Error(err)
	}

}

func assertTrue(actual string, expected string, msg string) (bool, error) {

	result := actual == expected

	if result {
		return result, nil
	}

	if len(msg) == 0 {
		msg = fmt.Sprintf("Expected %v actual value was %v", expected, actual)
	}

	return result, errors.New(msg)
}