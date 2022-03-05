// Удалить i-ый элемент из слайса.
package main

import (
	"fmt"
)

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

// slice [i, j)
func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	a := removeIndex(arr, 2)
	fmt.Println(a)
}
