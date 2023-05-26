package main

import "fmt"

/*
Часовая стрелка повернулась с начала суток на d градусов.
Определите, сколько сейчас целых часов h и целых минут m.
На вход программе подается целое число d (0d360).
*/

func main() {
	/*
	var d int
	fmt.Scan(&d)

	hours := d * 2 / 60
	minutes := d * 2 % 60
	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
	*/

	var d int
    fmt.Scanln(&d)
    h := d/30
    m := 2 * (d  %  30)
    
    fmt.Println("It is",h,"hours",m,"minutes.")
}