package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input1, input2 := getInput()
	fmt.Printf("Inputs are %f, %f\n", input1, input2)

	result := calculate(input1, input2)
	fmt.Printf("The results is %f\n", result)
}

func calculate(input1 float64, input2 float64) float64 {
	return input1 + input2
}

func getInput() (float64, float64) {
	reader := bufio.NewReader(os.Stdin)

	input1 := addInput(reader, "Enter first number: ")
	input2 := addInput(reader, "Enter second number: ")

	return parseInputToFloat(input1), parseInputToFloat(input2)
}

func addInput(reader *bufio.Reader, prompt string) string {
	fmt.Println(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic("Couldn't read input.")
	}
	return input
}

func parseInputToFloat(input string) float64 {
	parsedInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		response := fmt.Sprintf("The following input %s cannot be parsed into a float.", input)
		panic(response)
	}
	return parsedInput
}
