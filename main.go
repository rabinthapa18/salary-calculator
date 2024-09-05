// TODO:
// 1. remove debug print statements
// 2. implement calculateAmount function

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
	MIDNIGHT_START_TIME = 20
)

func main() {
	// reader to read the full line
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter input: ")

	// reading the input line
	input, _ := reader.ReadString('\n')

	// splitting the input into an array of strings
	parts := strings.Fields(input)

	// parsing to integers
	regularRrate, _ := strconv.Atoi(parts[0])
	nightRrate, _ := strconv.Atoi(parts[1])
	midNightRate, _ := strconv.Atoi(parts[2])
	days, _ := strconv.Atoi(parts[3])

	// for debugging
	fmt.Printf("Regular Rate: %d\nNight Rate: %d\nMidnight Rate: %d\nDays: %d\n", regularRrate, nightRrate, midNightRate, days)

	// parsing the intime and outtime pairs
	for i := 0; i < days; i++ {
		startTime, _ := strconv.Atoi(parts[4+i*2])
		endTime, _ := strconv.Atoi(parts[5+i*2])

		// for debugging
		fmt.Printf("start time: %d\nend time: %d\n", startTime, endTime)
	}
}

func calculateAmount(rate int, timeStamps [][]int) {}
