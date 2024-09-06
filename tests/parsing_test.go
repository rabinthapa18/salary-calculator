package test

import (
	"salary-calculator/pkg"
	"testing"
)

func TestParseInt(t *testing.T) {
	// testing the integer parsing
	i := pkg.ParseInt("5")
	if i != 5 {
		t.Errorf("Expected 5, got %d", i)
	}
}
