package main

import (
	calculator "calculator/calculator"
	"fmt"
)

func main() {
	var x, y int
	var operator string
	var result int
	var err error

	fmt.Print("Введите выражение (например, 3 + 2): ")
	fmt.Scanln(&x, &operator, &y)

	switch operator {
	case "+":
		result, err = calculator.Addition(x, y)
	case "-":
		result, err = calculator.Subtraction(x, y)
	case "*":
		result, err = calculator.Multiplication(x, y)
	case "/":
		result, err = calculator.Division(x, y)
	default:
		fmt.Println("Неизвестный оператор")
		return
	}

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Результат:", result)
}
