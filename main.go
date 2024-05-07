package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter first number:")
	var number1, number2 int
	_, _ = fmt.Scan(&number1)
	fmt.Println("Enter second number:")
	_, _ = fmt.Scan(&number2)
	fmt.Println("Enter desired operator (+-*/%):")
	operator, _, _ := reader.ReadRune()

	switch op := operator; op {
	case '+':
		fmt.Println(add(number1, number2))
	case '-':
		fmt.Println(subtract(number1, number2))
	case '*':
		fmt.Println(multiply(number1, number2))
	case '/':
		fmt.Println(divide(number1, number2))
	case '%':
		fmt.Println(modulo(number1, number2))
	default:
		fmt.Println("Operator not found, try again.")
	}
}

func add(number1 int, number2 int) int {
	return number1 + number2
}

func subtract(number1 int, number2 int) int {
	return number1 - number2
}

func multiply(number1 int, number2 int) int {
	return number1 * number2
}

func divide(number1 int, number2 int) int {
	return number1 / number2
}

func modulo(number1 int, number2 int) int {
	return number1 % number2
}
