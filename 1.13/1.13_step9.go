package main

import (
	"fmt"
)

/*
Количество минимумов
Найдите количество минимальных элементов в последовательности.

Входные данные
Вводится натуральное число N, а затем N чисел.

Выходные данные
Выведите количество минимальных элементов.

Sample Input:
4
21
1
11
1

Sample Output:
2
*/

func main() {
	var n, x int
	fmt.Scan(&n)

	fmt.Scan(&x)
	min := x
	count := 1

	for i := 1; i < n; i++ {
		fmt.Scan(&x)

		if x < min {
			min = x
			count = 0
		}
		if x == min {
			count++
		}
	}
	fmt.Println(count)
}