package main

import (
	"fmt"
)

/*
По данному числу N распечатайте все целые значения степени двойки, не
превосходящие N, в порядке возрастания.

Входные данные
Вводится натуральное число.

Выходные данные
Выведите ответ на задачу.

Sample Input:
50

Sample Output:
1 2 4 8 16 32
*/

func main() {
	var n int
	fmt.Scan(&n)
	
	// for i:=1;i<=n;i*=2 {
    //     fmt.Print(i," ")
    // }

	// for i:=1;i<=n;i+=i {
    //     fmt.Print(i," ")
    // }

	a := 1
	for a <= n {
		fmt.Print(a, " ")
		a *= 2
	}
	
}