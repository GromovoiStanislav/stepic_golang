package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	/*
	if n > 0 {
		fmt.Println("Число положительное")
	} else if n < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Ноль")
	}
	*/

	switch  {
        case n > 0: fmt.Println("Число положительное")
        case n == 0: fmt.Println("Ноль")
        default: fmt.Println("Число отрицательное")
    }

}