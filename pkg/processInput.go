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

	// checking if the number of pairs of in time and out time is correct
	if !CheckPairQuantity(days, parts) {
		return 0, fmt.Errorf("expected %d pairs of in time and out time", days)
	}

	// calculating the total amount
	totalAmount := GetTotalAmount(regularRate, nightRate, midNightRate, days, parts[4:])
	return totalAmount, nil
}
