package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Введите первое целое число: ")
	fmt.Scan(&num1)
	fmt.Print("Введите второе целое число: ")
	fmt.Scan(&num2)

	sum := num1 + num2
	difference := num1 - num2
	multiplication := num1 * num2
	division := (num1) / (num2)

	fmt.Printf("Сумма: %d + %d = %d\n", num1, num2, sum)
	fmt.Printf("Разность: %d - %d = %d\n", num1, num2, difference)
	fmt.Printf("Произведение: %d * %d = %d\n", num1, num2, multiplication)
	fmt.Printf("Деление: %d / %d = %d.\n", num1, num2, division)
}
