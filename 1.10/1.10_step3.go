package main

import "fmt"

/*
Требуется написать программу, при выполнении которой с клавиатуры считываются
два натуральных числа A и B (каждое не более 100, A < B).
Вывести сумму всех чисел от A до B включительно.
*/

func main() {

    var a, b , sum int
	fmt.Scanln(&a, &b)
    
    for i := a; i <= b; i++ {
		sum += i
	}
    fmt.Println(sum)

}
