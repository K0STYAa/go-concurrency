/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false
*/
package main

import (
	"fmt"
	"strings"
)

func isUniq(str string) bool {
	m := make(map[string]bool)
	str = strings.ToLower(str)

	for _, i := range str {
		if m[string(i)] == true {
			return false
		}
		m[string(i)] = true
	}

	return true
}

func main() {
	s := "abCdefAaf"
	fmt.Println(isUniq(s))
}
