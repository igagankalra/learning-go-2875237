package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("A simple calculator")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the first value:")
	input1, _ := reader.ReadString('\n')
	value1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		panic("Value 1 must be a number")

	}

	fmt.Print("Enter the second value:")
	input2, _ := reader.ReadString('\n')

	value2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		panic("Value 2 must be a number")
	}
	fmt.Println("select operation by entering one of these acronym \n add for Addition \n sub for subtraction \n mul for Multiplication \n div for division")
	input3, _ := reader.ReadString('\n')
	operator := strings.ToLower(strings.TrimSpace(input3))
	switch operator {
	case "add":
		sum := value1 + value2
		fmt.Printf("The sum of %.2f and %.2f is %.2f", value1, value2, sum)
	case "sub":
		difference := value1 - value2
		fmt.Printf("The Difference of %.1f and %.1f is %.1f", value1, value2, difference)
	case "mul":
		multiply := value1 * value2
		fmt.Printf("The Difference of %.1f and %.1f is %.1f", value1, value2, multiply)
	case "div":
		if value2 <= 0 {
			panic("Can Not divide by zero")
		}
		divide := value1 / value2
		fmt.Printf("The Difference of %.1f and %.1f is %.1f", value1, value2, divide)
	default:
		panic("Value should be one of these only -- >'add', 'sub' 'mul', 'div'  ")

	}

}
