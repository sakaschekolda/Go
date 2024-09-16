package main

import "fmt"

func main() {
	var num1, num2 float64

	fmt.Print("Введите первое число с плавающей запятой: ")
	fmt.Scan(&num1)
	fmt.Print("Введите второе число с плавающей запятой: ")
	fmt.Scan(&num2)

	sum := num1 + num2
	difference := num1 - num2

	fmt.Printf("Сумма: %.2f + %.2f = %.2f\n", num1, num2, sum)
	fmt.Printf("Разность: %.2f - %.2f = %.2f\n", num1, num2, difference)
}
