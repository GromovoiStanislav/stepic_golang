package main

import "fmt"

/*
Требуется определить, является ли данный год високосным
*/

func main() {
	var n int
	fmt.Scan(&n)

	if n%400 == 0 || (n%4 == 0 && n%100 != 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}