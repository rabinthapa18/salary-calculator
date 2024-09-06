package main

import (
	"bufio"
	"fmt"
	"os"
	"salary-calculator/pkg"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter input: ")

	input, _ := reader.ReadString('\n')
	result, err := pkg.ProcessInput(input)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result)
}
