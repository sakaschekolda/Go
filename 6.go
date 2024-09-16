package main

import "fmt"

func main() {
	var num1, num2, num3 float64

	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)
	fmt.Print("Введите третье число: ")
	fmt.Scan(&num3)

	average := (num1 + num2 + num3) / 3

	fmt.Printf("Среднее значение: (%.2f + %.2f + %.2f) / 3 = %.2f\n", num1, num2, num3, average)
}
