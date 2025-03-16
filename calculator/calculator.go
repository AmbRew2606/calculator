package calculator

import "fmt"

func Addition(x, y int) (int, error) {
	return x + y, nil
}

func Subtraction(x, y int) (int, error) {
	return x - y, nil
}

func Multiplication(x, y int) (int, error) {
	return x * y, nil
}

func Division(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("деление на ноль запрещено")
	}
	return x / y, nil
}
