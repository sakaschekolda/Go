package main

import "fmt"

func main() {
	var num1, num2, num3 float64

	fmt.Print("Enter 1st number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter 2nd number: ")
	fmt.Scan(&num2)
	fmt.Print("Enter 3rd number: ")
	fmt.Scan(&num3)

	average := (num1 + num2 + num3) / 3

	fmt.Printf("Average: (%.2f + %.2f + %.2f) / 3 = %.2f\n", num1, num2, num3, average)
}
