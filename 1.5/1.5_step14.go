package main

import "fmt"

/*
Дано натуральное число, не превосходящее 10000.
Выведите его последнюю цифру
*/

func main() {
	var a int
    fmt.Scanln(&a)
    fmt.Println(a % 10)
}