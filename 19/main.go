/*
	Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
	Символы могут быть unicode.
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

// https://github.com/golang/example/blob/master/stringutil/reverse.go

func reverse(s string) string {
	r := []rune(s)
	for i := 0; i < len(r)/2; i = i + 1 {
		j := len(r) - i - 1
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func rev(str string) string {
	runes := make([]rune, utf8.RuneCountInString(str))
	num := len(runes)

	for _, i := range str {
		num--
		runes[num] = i
	}
	return string(runes)
}

func main() {
	//string := "Hello, World"
	str := "!dlroW ,olleH"
	s := "\"😩😣😢😭😱\""
	fmt.Println(reverse(str))
	fmt.Println(rev(s))
}
