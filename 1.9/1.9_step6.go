package main

import "fmt"

/*
По данному трехзначному числу определите, все ли его цифры различны.
*/

func main() {
	/*
	var n string
	fmt.Scan(&n)
	
	if n[0] != n[1] && n[0] != n[2] && n[1] != n[2] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	*/

	var n int
    fmt.Scan(&n)

	a := n/100
    b := n/10%10
    c := n%10
    
    switch  {
        case a!=b && b!=c && a!=c : fmt.Println("YES")
        default : fmt.Println("NO")
    }
	
}