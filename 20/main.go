/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	words := strings.Fields(str)

	for i := 0; i < (len(words)-1)/2; i++ {
		j := len(words) - i - 1
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func rev(str string) string {
	words := strings.Fields(str)
	var res strings.Builder

	// пррсто будем записывать в обратном порядке
	for i := len(words) - 1; i >= 0; i-- {
		res.WriteString(words[i])
		res.WriteString(" ")
	}

	return strings.TrimSpace(res.String())
}

func main() {
	fmt.Printf("Reversed: %s\n", reverseWords("one   two      three"))
	fmt.Println(rev("snow     dog  sun"))
}
