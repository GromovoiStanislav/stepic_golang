package main

import "fmt"

/*
Дано неотрицательное целое число не превосходящее 10000.
Найдите и выведите первую цифру числа.
*/

func main() {
	/*
	var n string
	fmt.Scan(&n)

	fmt.Println(string(n[0]))
	// OR:
	// fmt.Println(n[:1])
	*/

	var n uint32
    fmt.Scan(&n)
    
    switch {
		case n < 10: fmt.Println(n)
		case n >= 10 && n < 100: fmt.Print(n/10)
		case n >= 100 && n < 1000: fmt.Print(n/100)
		case n >= 1000 && n < 10000: fmt.Print(n/1000)
		default: fmt.Print(10000/10000)   
    }
}