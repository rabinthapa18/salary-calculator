// this is the main entry point of the application for CLI

package main

import (
	"fmt"
	"os"
	"salary-calculator/pkg"
	"strings"
)

func main() {
	// getting the input from the command line arguments
	argInput := os.Args[1:]
	input := strings.Join(argInput, " ")

	result, err := pkg.ProcessInput(input)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result)
}
