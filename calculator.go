package main

import "fmt"

func main() {
	num1, num2 := 23, 8

	// Addition
	sum := add(num1, num2)
	fmt.Println("Sum:", sum)

	// Subtraction
	difference := subtract(num1, num2)
	fmt.Println("Difference:", difference)

	// Multiplication
	product := multiply(num1, num2)
	fmt.Println("Product:", product)

	// Division
	quotient := divide(num1, num2)
	fmt.Println("Quotient:", quotient)
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) float64 {
	return float64(a) / float64(b)
}
