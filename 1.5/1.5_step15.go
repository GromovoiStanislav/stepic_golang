package main

import "fmt"

/*
Дано неотрицательное целое число, не превосходящее 10000. 
Найдите число десятков (то есть вторую цифру справа). 
*/

func main() {
	var a int
	fmt.Scan(&a)

	fmt.Println(a / 10 % 10)
	// OR
	// fmt.Println(a % 100 / 10)
}