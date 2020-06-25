package calculator

import "errors"

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) (int, error) {
	if b == 0 {
		err := errors.New("You cannot divide by zero, try again")
		return 0, err
	}
	return a / b, nil
}
