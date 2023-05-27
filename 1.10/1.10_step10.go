package main

import "fmt"


/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются.
Числа в пределах от 0 до 10000. Программа должна вывести цифры, которые имеются в обоих числах, через пробел.
Цифры выводятся в порядке их нахождения в первом числе.
Ввод: 564 8954 Вывод: 5 4
*/


func main() {
	
	/*
	var a, b string
	fmt.Scan(&a, &b)

	for _, a1 := range a {
		for _, b1 := range b {
			if a1 == b1 {
				fmt.Print(string(a1) + " ")
			}
		}
	}
	*/

	var a, b, x, y int
	fmt.Scan(&a, &b)

	/*
	var res string
	for a > 0 {
		x = a % 10
		a = a / 10
		y = b
		for y > 0 {
			if y % 10 == x {
				res = fmt.Sprint(x) + " " + res
			}
			y = y / 10
		}
	}
	fmt.Println(res)
	*/


	j := 10000
	for a > 0 {
		x = a / j
		a = a % j
		y = b
		for y > 0 && x > 0 {
			if y % 10 == x {
				fmt.Print(x, " ")
			}
			y = y / 10
		}
		j = j / 10
	}

}