package main

import (
	"fmt"
)

/*
На ввод подаются пять чисел, которые записываются в массив.
Вам нужно  найти и вывести максимальное число в этом массиве.
*/

func main() {
	array := [5]int{}
	// var a int
	// for i:=0; i < 5; i++{
	// 	fmt.Scan(&a)
	// 	array[i] = a
	// }

	for i := range array {
		fmt.Scan(&array[i])
	}
    
   maxN :=array[0]
    for _, i:= range array{
        if i > maxN{
            maxN = i
        }
    }
    fmt.Println(maxN)
}