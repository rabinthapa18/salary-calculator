package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"salary-calculator/pkg"
)

// readTestCases reads test cases from a file
func ReadTestCases(filepath string) ([]struct {
	input    string
	expected int
}, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open test cases file: %w", err)
	}

	// ensuring the file is closed after reading
	defer file.Close()

	var testCases []struct {
		input    string
		expected int
	}

	// reading the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// splitting the line into input and expected value
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid test case format: %v", line)
		}

		var input string = parts[0]
		var expected int = pkg.ParseInt(parts[1])

		testCases = append(testCases, struct {
			input    string
			expected int
		}{input, expected})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading test cases: %w", err)
	}

	return testCases, nil
}

func main() {
	// read test cases from file
	testCases, err := ReadTestCases("testCases/test_cases.txt")
	if err != nil {
		fmt.Printf("error reading test cases: %v\n", err)
		os.Exit(1)
	}

	// testing each test case
	for _, tc := range testCases {
		result, err := pkg.ProcessInput(tc.input)
		if err != nil {
			fmt.Printf("error processing input %s: %v\n", tc.input, err)
			continue
		}
		if result != tc.expected {
			fmt.Printf("test failed: for input %s, expected %d, got %d\n", tc.input, tc.expected, result)
		} else {
			fmt.Printf("test passed: for input %s, got %d\n", tc.input, result)
		}
	}
}
