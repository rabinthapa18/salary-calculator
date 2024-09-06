package test

import (
	"salary-calculator/pkg"
	"testing"
)

func TestGetTotalAmount(t *testing.T) {
	// testing the total amount calculation
	regularRate := 1000
	nightRate := 1300
	midNightRate := 1500
	days := 4
	parts := []string{"0", "9", "9", "17", "17", "22", "22", "23"}

	totalAmount, err := pkg.GetTotalAmount(regularRate, nightRate, midNightRate, days, parts)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if totalAmount != 29500 {
		t.Errorf("Expected 29500, got %d", totalAmount)
	}
}
