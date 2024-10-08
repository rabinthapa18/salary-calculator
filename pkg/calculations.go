package pkg

import "fmt"

func GetTotalAmount(regularRate, nightRate, midNightRate, days int, parts []string) (int, error) {
	// checking if the number of pairs of in time and out time is correct
	if !CheckPairQuantity(days, parts) {
		return 0, fmt.Errorf("expected %d pairs of in time and out time", days)
	}

	totalAmount := 0

	// parsing the intime and outtime pairs
	for i := 0; i < days; i++ {
		startTime := ParseInt(parts[i*2])
		endTime := ParseInt(parts[1+i*2])

		totalAmount += calculateDayAmount(startTime, endTime, regularRate, nightRate, midNightRate)
	}

	return totalAmount, nil
}

func calculateDayAmount(startTime, endTime, regularRate, nightRate, midNightRate int) int {
	dailyAmount := 0
	for hour := startTime; hour < endTime; hour++ {
		if hour >= 0 && hour < REGULAR_START_TIME || hour >= MIDNIGHT_START_TIME {
			// midnight rate
			dailyAmount += midNightRate
		} else if hour >= REGULAR_START_TIME && hour < NIGHT_START_TIME {
			// regular rate
			dailyAmount += regularRate
		} else if hour >= NIGHT_START_TIME && hour < MIDNIGHT_START_TIME {
			// night rate
			dailyAmount += nightRate
		}
	}
	return dailyAmount
}
