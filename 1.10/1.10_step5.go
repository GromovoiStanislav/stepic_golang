package main

import "fmt"

/*
Последовательность состоит из натуральных чисел и завершается числом 0.
Определите количество элементов этой последовательности, которые равны
ее наибольшему элементу.

Формат входных данных
Вводится непустая последовательность натуральных чисел, оканчивающаяся числом
0 (само число 0 в последовательность не входит, а служит как признак ее окончания).

Формат выходных данных
Выведите ответ на задачу.

Sample Input:
1
3
3
1
0

Sample Output:
2
*/

func main() {
	/*
	var n int
	max := 0
	number := 0

	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}

		if n > max {
			max = n
			number = 0
		}
		if n == max {
			number++
		}
	}
	fmt.Println(number)
	*/


	var n int
	max := 0
	number := 0
    
	for fmt.Scan(&n); n != 0; fmt.Scan(&n){// считываем числа пока не будет введен 0
		if n > max {
			max = n
			number = 0
		}
		if n == max {
			number++
		}
	}
    fmt.Println(number)


}