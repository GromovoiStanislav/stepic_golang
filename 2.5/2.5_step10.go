package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

/*
На вход дается строка, из нее нужно сделать другую строку, оставив только
нечетные символы (считая с нуля)

Sample Input:
ihgewlqlkot

Sample Output:
hello
*/

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	var b strings.Builder
	//var str []rune
	for i, r := range text {
		if i%2 != 0 {
			b.WriteRune(r)
			//str = append(str, r)
			//fmt.Print(string(r))
		}
	}

	fmt.Println(b.String())
	//fmt.Print(string(str))
}