package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
На вход подается строка. Нужно определить, является ли она палиндромом.
Если строка является паллиндромом - вывести Палиндром иначе - вывести Нет.
Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих
направлениях (например, "топот", "око", "заказ").

Sample Input:
топот

Sample Output:
Палиндром
*/

func ToReversed(text string) string {
	rs := []rune(text)

	reversed_rs := make([]rune, len(rs))
	for i := 0; i < len(reversed_rs); i++ {
		reversed_rs[i] = rs[len(reversed_rs)-i-1]
	}
	return string(reversed_rs)
}

func main() {
	// var text string
	// _, _ = fmt.Scan(&text)
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	reversed := ToReversed(text)

	if text == reversed {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}