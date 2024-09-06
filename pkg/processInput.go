package pkg

import (
	"fmt"
	"strings"
)

// processInput handles the main logic and can be called from tests or other entry points.
func ProcessInput(input string) (int, error) {
	// splitting the input into an array of strings
	parts := strings.Fields(input)

	// checking if there are enough required elements
	if len(parts) < 4 {
		return 0, fmt.Errorf("incomplete required elements")
	}

	// parsing to integers
	regularRate := ParseInt(parts[0])
	nightRate := ParseInt(parts[1])
	midNightRate := ParseInt(parts[2])
	days := ParseInt(parts[3])

	// calculating the total amount
	totalAmount, err := GetTotalAmount(regularRate, nightRate, midNightRate, days, parts[4:])
	if err != nil {
		return 0, err
	}
	return totalAmount, nil
}
