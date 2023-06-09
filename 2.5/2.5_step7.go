package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
На вход подается строка. Нужно определить, является ли она правильной или нет.
Правильная строка начинается с заглавной буквы и заканчивается точкой.
Если строка правильная - вывести Right иначе - вывести Wrong

подсказка: fmt.Scan() считывает строку до первого пробела, вы можете
считать полностью строку за раз с помощью bufio:

text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

Sample Input:
Быть или не быть.

Sample Output:
Right
*/

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)

	rs := []rune(text)

	//if unicode.IsUpper(rs[0]) && rs[len(rs)-1] == '.' {
	if unicode.IsUpper(rs[0]) && strings.HasSuffix(text, ".") {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}