package test

import (
	"salary-calculator/pkg"
	"testing"
)

func TestProcessInput(t *testing.T) {
	// testing the input processing
	input := "1000 1300 1500 4 0 9 9 17 17 22 22 23"
	result, err := pkg.ProcessInput(input)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if result != 29500 {
		t.Errorf("Expected 29500, got %d", result)
	}
}
