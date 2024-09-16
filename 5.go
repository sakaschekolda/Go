package main

import "fmt"

func main() {
	var num1, num2 float64

	fmt.Print("Enter 1st float: ")
	fmt.Scan(&num1)
	fmt.Print("Enter 2nd float: ")
	fmt.Scan(&num2)

	sum := num1 + num2
	difference := num1 - num2

	fmt.Printf("Summary: %.2f + %.2f = %.2f\n", num1, num2, sum)
	fmt.Printf("Difference: %.2f - %.2f = %.2f\n", num1, num2, difference)
}
