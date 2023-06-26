package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

/*
Дается строка. Нужно удалить все символы, которые встречаются более одного раза
и вывести получившуюся строку

Sample Input:
zaabcbd

Sample Output:
zcd
*/

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	var b strings.Builder
	for _, r := range text {
		if strings.Count(text, string(r)) == 1 {
			b.WriteRune(r)
			//fmt.Print(string(r))
		}
	}

	fmt.Println(b.String())
}