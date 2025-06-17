package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Mini Calculator in Go")
	fmt.Println("---------------------")

	// Read first number
	fmt.Print("Enter first number: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	num1, err := strconv.ParseFloat(input1, 64)
	if err != nil {
		fmt.Println("Invalid number:", input1)
		return
	}

	// Read operator
	fmt.Print("Enter operator (+, -, *, /): ")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	// Read second number
	fmt.Print("Enter second number: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	num2, err := strconv.ParseFloat(input2, 64)
	if err != nil {
		fmt.Println("Invalid number:", input2)
		return
	}

	// Calculate result
	var result float64
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operator:", op)
		return
	}

	fmt.Printf("Result: %.4f\n", result)
}
