package main

import (
	"fmt"
	// "strconv"
	// "strings"
)

/*
Из натурального числа удалить заданную цифру.

Входные данные
Вводятся натуральное число и цифра, которую нужно удалить.

Выходные данные
Вывести число без заданных цифр.

Sample Input:
38012732
3

Sample Output:
801272
*/

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	// fmt.Println(strings.Replace(strconv.Itoa(a), strconv.Itoa(b), "", -1))

	result := 0
	for i := 1; a > 0; {
		if a % 10 != b {
			result += a % 10 * i
			i *= 10
		}
		a /= 10
	}
	fmt.Println(result)
}
	
