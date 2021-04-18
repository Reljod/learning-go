package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input1 := addInputNum("Enter input 1: ")
	input2 := addInputNum("Enter input 2: ")
	operation := addInputOperation("Choose operation from the following (+ - * /): ")

	result, err := calculate(input1, input2, operation)
	if err != nil {
		response := fmt.Sprintf("Calculation failed for inputs: (%v, %v) and operation: %v", input1, input2, operation)
		panic(response)
	}

	fmt.Printf("%v %v %v = %f", input1, operation, input2, result)
}

func calculate(input1 float64, input2 float64, operation string) (float64, error) {

	var result float64
	var err error

	switch operation {
	case "+":
		result = addValues(input1, input2)
	case "-":
		result = subtractValues(input1, input2)
	case "*":
		result = multiplyValues(input1, input2)
	case "/":
		result = divideValues(input1, input2)
	default:
		result = 0
		err = fmt.Errorf("operator %v invalid", operation)
	}

	return result, err
}

func addValues(input1 float64, input2 float64) float64 {
	return input1 + input2
}

func subtractValues(input1 float64, input2 float64) float64 {
	return input1 - input2
}

func multiplyValues(input1 float64, input2 float64) float64 {
	return input1 * input2
}

func divideValues(input1 float64, input2 float64) float64 {
	if input2 == 0 {
		panic("Cannot divide a number to zero.")
	}
	return input1 / input2
}

func addInputNum(prompt string) float64 {
	inputStr := addInput(prompt)
	parsedInput, err := strconv.ParseFloat(inputStr, 64)
	if err != nil {
		response := fmt.Sprintf("Cannot convert input %v to float", inputStr)
		panic(response)
	}
	return parsedInput
}

func addInputOperation(prompt string) string {
	var operation = make(map[string]string)
	operation["+"] = "+"
	operation["-"] = "-"
	operation["*"] = "*"
	operation["/"] = "/"

	input := addInput(prompt)

	value, ok := operation[input]
	if !(ok) {
		respose := fmt.Sprintf("Opeation %v cannot be used.", input)
		panic(respose)
	}

	return value
}

func addInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		panic("Error reading input.")
	}

	return input
}
