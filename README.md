Написать программу, которая выводит текущее время и дату. Создать переменные различных типов (int, float64, string, bool) и вывести их на экран.
Использовать краткую форму объявления переменных для создания и вывода переменных. Написать программу для выполнения арифметических операций с двумя целыми числами и выводом результатов.
Реализовать функцию для вычисления суммы и разности двух чисел с плавающей запятой. Написать программу, которая вычисляет среднее значение трех чисел.

1. 
```
   	package main

	import (
	"fmt"
	"time"
	)

	func main() {
		currentTime := time.Now()

		formattedTime := currentTime.Format("02-01-2006 15:04:05")

		fmt.Println("Текущая дата и время:", formattedTime)
	}
```
2. 
```
   	package main

	import "fmt"

	func main() {
		var integerVar int = 11
		var floatVar float64 = 3.14
		var stringVar string = "Hello world"
		var boolVar bool = true
	
		fmt.Println("Целочисленное значение:", integerVar)
		fmt.Println("Число с плавающей запятой:", floatVar)
		fmt.Println("Строка:", stringVar)
		fmt.Println("Логическое значение:", boolVar)
	}

```
3.
```
	package main

	import "fmt"

	func main() {
		integerVar := 11
		floatVar := 3.14
		stringVar := "Hello world"
		boolVar := true
	
		fmt.Println("Целочисленное значение:", integerVar)
		fmt.Println("Число с плавающей запятой:", floatVar)
		fmt.Println("Строка:", stringVar)
		fmt.Println("Логическое значение:", boolVar)
	}
```
4.
```
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

```
1.txt
```
package main

import "fmt"

func main() {
	// Объявление двух целых переменных
	var num1, num2 int

	// Ввод значений от пользователя
	fmt.Print("Введите первое целое число: ")
	fmt.Scan(&num1)
	fmt.Print("Введите второе целое число: ")
	fmt.Scan(&num2)

	// Выполнение арифметических операций
	sum := num1 + num2
	difference := num1 - num2
	Multiplication := num1 * num2
	// Проверка на деление на ноль для num1
	if num1 != 0 {
		quotient2 := float64(num2) / float64(num1) // num2 / num1
		fmt.Printf("Деление num2 / num1: %d / %d = %.2f\n", num2, num1, quotient2)
	} else {
		fmt.Println("Ноль нельзя делить")
	}

	// Проверка на деление на ноль для num2
	if num2 != 0 {
		quotient1 := float64(num1) / float64(num2) // num1 / num2
		fmt.Printf("Частное num1 / num2: %d / %d = %.2f\n", num1, num2, quotient1)
	} else {
		fmt.Println("Нельзя делить на ноль")
	}

	// Вывод результатов на экран
	fmt.Printf("Сумма: %d + %d = %d\n", num1, num2, sum)
	fmt.Printf("Разность: %d - %d = %d\n", num1, num2, difference)
	fmt.Printf("Произведение: %d * %d = %d\n", num1, num2, Multiplication)
}
```
5. tx
```
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

```
6.
```
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

```
