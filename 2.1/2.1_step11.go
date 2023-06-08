package main

import (
	"fmt"
)

/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int,
и возвращающую количество полученных функцией аргументов и их сумму.

Пример вызова вашей функции:

a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1
*/

func sumInt(numbers ...int) (count int, sum int) {
	for _, number := range numbers {
		sum += number
		count++
	}
	return count, sum
}

func main() {
	fmt.Println(sumInt(4))
	fmt.Println(sumInt(1, 2, 3, 4))
}