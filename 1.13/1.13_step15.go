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
	var num int
	fmt.Scan(&num)

	//fmt.Printf("%b\n", num)



	// for num > 0 {
    //     defer fmt.Print(num % 2)
    //     num /= 2
    // }



	sum := 0
	digit := 1
	for num != 0 {
		sum += num % 2 * digit
		digit *= 10
		num /= 2
	}
	fmt.Println(sum)
}