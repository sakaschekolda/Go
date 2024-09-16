package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Enter 1st integer: ")
	fmt.Scan(&num1)
	fmt.Print("Enter 2nd integer: ")
	fmt.Scan(&num2)

	sum := num1 + num2
	difference := num1 - num2
	multiplication := num1 * num2
	division := (num1) / (num2)

	fmt.Printf("Summary: %d + %d = %d\n", num1, num2, sum)
	fmt.Printf("Difference: %d - %d = %d\n", num1, num2, difference)
	fmt.Printf("Multiplication: %d * %d = %d\n", num1, num2, multiplication)
	fmt.Printf("Division: %d / %d = %d.\n", num1, num2, division)
}
