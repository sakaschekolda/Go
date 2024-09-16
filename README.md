1) Написать программу, которая выводит текущее время и дату.
2) Создать переменные различных типов (int, float64, string, bool) и вывести их на экран.
3) Использовать краткую форму объявления переменных для создания и вывода переменных.
4) Написать программу для выполнения арифметических операций с двумя целыми числами и выводом результатов.
5) Реализовать функцию для вычисления суммы и разности двух чисел с плавающей запятой.
6) Написать программу, которая вычисляет среднее значение трех чисел.

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
	
		fmt.Println("Date and time now:", formattedTime)
	}

```
![{F6BA2D89-0DF5-45B9-9127-556A68F1CB03}](https://github.com/user-attachments/assets/d37578cd-82f0-4100-8412-e62ec9a55e9a)
2. 
```
   	package main

	import "fmt"
	
	func main() {
		var integerVar int = 11
		var floatVar float64 = 3.14
		var stringVar string = "Hello world"
		var boolVar bool = true
	
		fmt.Println("Integer:", integerVar)
		fmt.Println("Float:", floatVar)
		fmt.Println("String:", stringVar)
		fmt.Println("Bool:", boolVar)
	}
```
![{B1E939F5-B7EC-4882-BC9D-23D06C423F38}](https://github.com/user-attachments/assets/802c28f9-0cba-4118-8538-9a5734463fa2)
3.
```
	package main

	import "fmt"
	
	func main() {
		integerVar := 11
		floatVar := 3.14
		stringVar := "Hello world"
		boolVar := true
	
		fmt.Println("Integer:", integerVar)
		fmt.Println("Float:", floatVar)
		fmt.Println("String:", stringVar)
		fmt.Println("Bool:", boolVar)
	}
```
![{B1E939F5-B7EC-4882-BC9D-23D06C423F38}](https://github.com/user-attachments/assets/a94ebc06-f748-4cd0-8afa-ebbbd897f88e)

4.
```
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
```
![{700732BD-044E-4428-890E-E3D7CB4C9645}](https://github.com/user-attachments/assets/0115c3e7-d4f8-4318-b474-c253f55b5260)

4.txt
```
	package main

	import "fmt"
	
	func main() {
		// Объявление двух целых переменных
		var num1, num2 int
	
		// Ввод значений от пользователя
		fmt.Print("Enter 1st integer: ")
		fmt.Scan(&num1)
		fmt.Print("Enter 2nd integer: ")
		fmt.Scan(&num2)
	
		// Выполнение арифметических операций
		sum := num1 + num2
		difference := num1 - num2
		Multiplication := num1 * num2
		// Проверка на деление на ноль для num1
		if num1 != 0 {
			quotient2 := float64(num2) / float64(num1) // num2 / num1
			fmt.Printf("Divide num2 / num1: %d / %d = %.2f\n", num2, num1, quotient2)
		} else {
			fmt.Println("Can't divide by zero")
		}
	
		// Проверка на деление на ноль для num2
		if num2 != 0 {
			quotient1 := float64(num1) / float64(num2) // num1 / num2
			fmt.Printf("division num1 / num2: %d / %d = %.2f\n", num1, num2, quotient1)
		} else {
			fmt.Println("Can't divide by zero")
		}
	
		// Вывод результатов на экран
		fmt.Printf("Summary: %d + %d = %d\n", num1, num2, sum)
		fmt.Printf("Difference: %d - %d = %d\n", num1, num2, difference)
		fmt.Printf("Multiplication: %d * %d = %d\n", num1, num2, Multiplication)
	}
```
5. 
```
	ppackage main
	
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
```
![{7DAF94E2-E981-4EC1-BB07-AA39ADCE3A5B}](https://github.com/user-attachments/assets/8573932f-27f2-4734-b8c5-ac7f379f7596)

6.
```
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
```
![{890899C1-1AA3-4E4E-A841-5CC40034366D}](https://github.com/user-attachments/assets/6336d340-ddbb-40ec-9a24-6e8f176bec2e)
