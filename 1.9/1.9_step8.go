package main

import (
	"fmt"
)

/*
Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
*/

func main() {
	var n int
	fmt.Scan(&n)

	/*
	left := n/100000%10 + n/10000%10 + n/1000%10
	right := n/100%10 + n/10%10 + n%10
	if left == right {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
    */
    
	number1 := n/100000
    number2 := n/10000%10
    number3 := n/1000%10
    number4 := n/100%10
    number5 := n/10%10
    number6 := n%10
	

     if number1 + number2 + number3 == number4 + number5 + number6 {
       fmt.Println("YES")
     }else {
		fmt.Println("NO")
	}

}