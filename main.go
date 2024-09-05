package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	REGULAR_START_TIME  = 9
	NIGHT_START_TIME    = 17
	MIDNIGHT_START_TIME = 22
)

func main() {
	// reader to read the full line
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter input: ")

	// reading the input line
	input, _ := reader.ReadString('\n')

	// splitting the input into an array of strings
	parts := strings.Fields(input)

	// checking if there are enough required elements
	if len(parts) < 4 {
		fmt.Println("Error: incomplete required elements")
		return
	}

	// parsing to integers
	regularRate, _ := strconv.Atoi(parts[0])
	nightRate, _ := strconv.Atoi(parts[1])
	midNightRate, _ := strconv.Atoi(parts[2])
	days, _ := strconv.Atoi(parts[3])

	// checking if the number of pairs of in time and out time is correct
	if !checkPairQuantity(days, parts) {
		fmt.Printf("Error: expected %d pairs of in time and out time\n", days)
		return
	}

	totalAmount := 0

	// parsing the intime and outtime pairs
	for i := 0; i < days; i++ {
		startTime, _ := strconv.Atoi(parts[4+i*2])
		endTime, _ := strconv.Atoi(parts[5+i*2])

		totalAmount += calculateDayAmount(startTime, endTime, regularRate, nightRate, midNightRate)
	}
}
func checkPairQuantity(n int, parts []string) bool {
	return len(parts) == 4+n*2
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
