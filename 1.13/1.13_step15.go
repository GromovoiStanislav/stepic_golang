package main

import (
	"fmt"
)

/*
Двоичная запись
Дано натуральное число N. Выведите его представление в двоичном виде.

Входные данные
Задано единственное число N

Выходные данные
Необходимо вывести требуемое представление числа N.

Sample Input:
12

Sample Output:
1100
*/

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Printf("%b\n", n)

	// for n > 0 {
    //     defer fmt.Print(n % 2)
    //     n /= 2
    // }
}