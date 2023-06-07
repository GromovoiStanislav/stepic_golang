package main

import "fmt"


/*
Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0.
Напишите программу, которая выведет элементы массива, индексы которых четны
(0, 2, 4...).
Входные данные
Сначала задано число N — количество элементов в массиве (1 ≤ N ≤ 100).
Далее через пробел записаны NNN чисел — элементы массива.
Массив состоит из целых чисел.
Выходные данные
Необходимо вывести все элементы массива с чётными номерами.
Sample Input:
6
4 5 3 4 2 3
Sample Output:
4 3 2
*/

func main() {
	var n uint
	fmt.Scan(&n)

	items := make([]int, n)
	for i := range items {
		fmt.Scan(&items[i])
	}

	for i, x := range items {
		if i%2 == 0 {
			fmt.Print(x, " ")
		}
	}
}