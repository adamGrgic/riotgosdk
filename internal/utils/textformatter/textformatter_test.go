package utils

import (
	"fmt"
	"testing"
)

const (
	GreenOpenTagTest  = "\033[32"
	GreenCloseTagTest = "\033[0m"
)

func TestGreenTextFormatter(t *testing.T) {
	testPhrase := "test"

	result := FmtTextColor(testPhrase, "green")
	expected := fmt.Sprintf("\033[32%s\033[0m", testPhrase)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
