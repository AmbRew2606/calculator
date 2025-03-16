package main

import (
	calculator "calculator/calculator"
	"fmt"
)

func main() {
	var x, y int
	var operator string
	var result int

	fmt.Scanln(&x, &operator, &y)
	switch operator {
	case "+":
		result = calculator.Addition(x, y)
	case "-":
		result = calculator.Sddition(x, y)
	case "*":
		result = calculator.Multiplication(x, y)
	case "/":
		result = calculator.Division(x, y)

	default:
		fmt.Println("Unknown operator")
		return
	}

	fmt.Println(result)
}
